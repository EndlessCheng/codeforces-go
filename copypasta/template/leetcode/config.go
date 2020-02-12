package leetcode

import "fmt"

const (
	contestID = 176

	hostZH = "leetcode-cn.com"
	hostEN = "leetcode.com"
	host   = hostZH

	contestPrefixWeekly   = "weekly-contest-"
	contestPrefixBiweekly = "biweekly-contest-"
	contestPrefix         = contestPrefixWeekly

	openWebPageZH = true
	openWebPageEN = true
)

var contestDir string

func init() {
	switch contestPrefix {
	case contestPrefixWeekly:
		contestDir = fmt.Sprintf("../../../leetcode/%d/", contestID)
	case contestPrefixBiweekly:
		contestDir = fmt.Sprintf("../../../leetcode/biweekly/%d/", contestID)
	}
}
