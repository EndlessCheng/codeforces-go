package leetcode

import (
	"os"
	"testing"
)

// TODO: 确认是否登录以及默认语言是否正确
func TestGenLeetCodeTests(t *testing.T) {
	var username, password string
	if host == hostZH {
		username = os.Getenv("LEETCODE_USERNAME_ZH")
		password = os.Getenv("LEETCODE_PASSWORD_ZH")
	} else {
		username = os.Getenv("LEETCODE_USERNAME_EN")
		password = os.Getenv("LEETCODE_PASSWORD_EN")
	}
	if err := GenLeetCodeTests(username, password); err != nil {
		t.Fatal(err)
	}
}

//func TestGenLeetCodeSpecialTests(t *testing.T) {
//	username := os.Getenv("LEETCODE_USERNAME_ZH")
//	password := os.Getenv("LEETCODE_PASSWORD_ZH")
//	urlZHs := []string{
//		"",
//		"",
//		"",
//		"",
//		"",
//	}
//	if err := GenLeetCodeSpecialTests(username, password, urlZHs); err != nil {
//		t.Fatal(err)
//	}
//}
