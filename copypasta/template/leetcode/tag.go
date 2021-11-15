package leetcode

import (
	"strconv"
	"time"
)

// id > 0，指定具体的一场周赛
// id = 0，指定下一场或当前正在进行的周赛
// id < 0，指定上 |id| 场周赛（例如 id = -1 表示最近的一场结束的周赛）

func GetWeeklyContestID(contestID int) int {
	if contestID <= 0 {
		utc8, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			panic(err)
		}

		// 以 2020 年第一场周赛的结束时间为基准
		endTime170 := time.Date(2020, 1, 5, 12, 0, 0, 0, utc8)
		weeksSince170 := 1 + int(time.Since(endTime170)/(7*24*time.Hour))
		contestID += 170 + weeksSince170
	}
	return contestID
}

func GetBiweeklyContestID(contestID int) int {
	if contestID <= 0 {
		utc8, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			panic(err)
		}

		// 以 2020 年第一场双周赛的结束时间为基准
		endTime17 := time.Date(2020, 1, 12, 0, 0, 0, 0, utc8)
		twoWeeksSince17 := 1 + int(time.Since(endTime17)/(14*24*time.Hour))
		contestID += 17 + twoWeeksSince17
	}
	return contestID
}

func GetWeeklyContestTag(contestID int) string {
	return "weekly-contest-" + strconv.Itoa(GetWeeklyContestID(contestID))
}

func GetBiweeklyContestTag(contestID int) string {
	return "biweekly-contest-" + strconv.Itoa(GetBiweeklyContestID(contestID))
}
