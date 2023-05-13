package leetcode

import (
	"fmt"
	"os"
	"testing"
)

// 由于力扣的限制，登录后会让网页端退出
// 建议额外用个号，这样可免去重登的麻烦
func TestWeekly(t *testing.T) {
	if err := genLeetCodeTests(true); err != nil {
		t.Fatal(err)
	}
}

func TestBiweekly(t *testing.T) {
	if err := genLeetCodeTests(false); err != nil {
		t.Fatal(err)
	}
}

var (
	username = os.Getenv("LEETCODE_USERNAME_ZH")
	password = os.Getenv("LEETCODE_PASSWORD_ZH")
	comment  = os.Getenv("LEETCODE_COMMENT")
)

func genLeetCodeTests(weekly bool) error {
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
	return GenLeetCodeTests(username, password, tag, true, dir, comment)
}

func TestGenLeetCodeSpecialTests(t *testing.T) {
	const tag = "" // tianchi2022
	const dir = "../../../leetcode/other/" + tag + "/"
	if err := GenLeetCodeTests(username, password, tag, true, dir, comment); err != nil {
		t.Fatal(err)
	}
}

func TestGenLeetCodeSeasonTests(t *testing.T) {
	const year = "2023"
	const season = SeasonSpring
	const solo = true

	dir := "../../../leetcode/season/" + year + season
	if !solo {
		dir += "2"
	}
	dir += "/"
	if err := GenLeetCodeSeasonTests(username, password, year+"-"+season, solo, true, dir, comment); err != nil {
		t.Fatal(err)
	}
}
