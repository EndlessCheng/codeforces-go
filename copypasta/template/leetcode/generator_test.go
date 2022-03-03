package leetcode

import (
	"fmt"
	"os"
	"testing"
)

// 由于力扣的限制，登录后会让网页端退出
// 建议额外用个号，这样可免去重登的麻烦
func TestGenLeetCodeTests(t *testing.T) {
	const weekly = true

	username := os.Getenv("LEETCODE_USERNAME_ZH")
	password := os.Getenv("LEETCODE_PASSWORD_ZH")

	var tag, dir string
	if weekly {
		contestID := GetWeeklyContestID(0) // 自动生成下一场周赛 ID
		tag = GetWeeklyContestTag(contestID)
		dir = fmt.Sprintf("../../../leetcode/weekly/%d/", contestID) // 自定义生成目录
	} else {
		contestID := GetBiweeklyContestID(0) // 自动生成下一场双周赛 ID
		tag = GetBiweeklyContestTag(contestID)
		dir = fmt.Sprintf("../../../leetcode/biweekly/%d/", contestID) // 自定义生成目录
	}

	if err := GenLeetCodeTests(username, password, tag, true, dir, ""); err != nil {
		t.Fatal(err)
	}
}
