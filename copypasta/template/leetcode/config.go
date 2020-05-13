package leetcode

import (
	"fmt"
	"time"
)

const (
	hostZH = "leetcode-cn.com"
	hostEN = "leetcode.com"
	host   = hostZH

	contestPrefixWeekly   = "weekly-contest-"
	contestPrefixBiweekly = "biweekly-contest-"
	contestPrefix         = contestPrefixWeekly

	openWebPageZH = true
	openWebPageEN = false
)

var (
	contestID  int
	contestDir string
)

func init() {
	if contestID == 0 {
		contestID = calcNextContestID()
	}

	switch contestPrefix {
	case contestPrefixWeekly:
		contestDir = fmt.Sprintf("../../../leetcode/%d/", contestID)
	case contestPrefixBiweekly:
		contestDir = fmt.Sprintf("../../../leetcode/biweekly/%d/", contestID)
	default:
		contestDir = fmt.Sprintf("../../../leetcode/%s/", contestPrefix)
	}
}

var utc8, _ = time.LoadLocation("Asia/Shanghai")

func calcNextContestID() int {
	switch contestPrefix {
	case contestPrefixWeekly:
		// 以 2020 年第一场周赛的结束时间为基准
		endTime170 := time.Date(2020, 1, 5, 12, 0, 0, 0, utc8)
		weeksSince170 := 1 + int(time.Since(endTime170)/(7*24*time.Hour))
		return 170 + weeksSince170
	case contestPrefixBiweekly:
		// 以 2020 年第一场双周赛的结束时间为基准
		endTime17 := time.Date(2020, 1, 12, 0, 0, 0, 0, utc8)
		twoWeeksSince17 := 1 + int(time.Since(endTime17)/(14*24*time.Hour))
		return 17 + twoWeeksSince17
	default:
		return -1
	}
}
