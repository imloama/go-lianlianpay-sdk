package lianlianpay

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/imloama/go-lianlianpay-sdk/schema"
	"github.com/imloama/go-lianlianpay-sdk/sign"
	"github.com/imroc/req/v3"
)

type LianlianPayClient struct {
	Gateway   string
	sig       *sign.Sign
	reqclient *req.Client
}

func NewLianlianPayClient(gateway string, secret string, llpubkey string, debug bool) (*LianlianPayClient, error) {
	sig, err := sign.NewSign(secret, llpubkey)
	if err != nil {
		return nil, err
	}
	reqclient := req.C()
	if debug {
		reqclient = reqclient.DevMode()
	}
	reqclient = reqclient.SetCommonHeader("Content-type", "application/json;charset=utf-8").
		SetCommonHeader("Signature-Type", "RSA")
	return &LianlianPayClient{
		Gateway:   gateway,
		sig:       sig,
		reqclient: reqclient,
	}, nil
}

func (client *LianlianPayClient) SignData(data interface{}) (string, error) {
	dd, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return client.sig.RsaSignWithMd5(dd)
}

func (client *LianlianPayClient) Verify(data []byte, signature string) error {
	return client.sig.RsaVerifySignWithMd5(data, signature)
}

func (client *LianlianPayClient) Exec(reqdata schema.IRequest, result schema.IResult) error {
	signature, err := client.SignData(reqdata)
	if err != nil {
		return err
	}
	url := client.Gateway + reqdata.GetUrl()
	rsp, err := client.reqclient.R().SetHeader("Signature-Data", signature).SetBody(reqdata).
		SetResult(result).EnableDump().
		Post(url)
	if err != nil {
		return err
	}
	if rsp.IsSuccess() {
		if result.GetRetCode() != RET_CODE_OK {
			msg := result.GetRetMsg()
			if msg == "" {
				msg = fmt.Sprintf("[%s]%s", result.GetRetCode(), RET_CODE_ERRORS[result.GetRetCode()])
			}
			return errors.New(msg)
		}
		return nil
	}
	return errors.New(rsp.Dump())
}

//func(client *LianlianPayClient)Execute(reqdata schema.IRequest, result schema.IResult)(res schema.IResult, err error){
//	err = client.Exec(reqdata, result)
//	if err!=nil{
//		return
//	}
//	return result, nil
//}
