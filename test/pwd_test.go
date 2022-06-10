package test

import (
	"fmt"
	"regexp"
	"testing"
)

func TestPwd(t *testing.T) {
	reg := `^[a-zA-Z0-9]{6,12}$`
	reg1 := regexp.MustCompile(reg)
	if !reg1.MatchString("sfa12") {
		fmt.Println("正确，密码不符合条件！")
	}
	if reg1.MatchString("abc1234") {
		fmt.Println("正确，密码符合条件！")
	}
}
