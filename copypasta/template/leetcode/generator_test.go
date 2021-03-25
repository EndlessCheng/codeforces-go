package leetcode

import (
	"os"
	"testing"
)

// 由于力扣的限制，登录后会让网页端退出
// 建议额外用个号，这样可免去重登的麻烦
func TestGenLeetCodeTests(t *testing.T) {
	var username, password string
	if host == hostZH {
		username = os.Getenv("LEETCODE_USERNAME_ZH")
		password = os.Getenv("LEETCODE_PASSWORD_ZH")
		if vipZH {
			username = os.Getenv("LEETCODE_USERNAME_VIP_ZH")
			password = os.Getenv("LEETCODE_PASSWORD_VIP_ZH")
		}
	} else {
		username = os.Getenv("LEETCODE_USERNAME_EN")
		password = os.Getenv("LEETCODE_PASSWORD_EN")
	}
	if err := GenLeetCodeTests(username, password, "// github.com/EndlessCheng/codeforces-go"); err != nil {
		t.Fatal(err)
	}
}
