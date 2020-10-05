package main

import (
	"sort"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func alertNames(keyName []string, keyTime []string) (ans []string) {
	tsMap := map[string][]int{}
	for i, name := range keyName {
		s := keyTime[i]
		h, _ := strconv.Atoi(s[:2])
		m, _ := strconv.Atoi(s[3:])
		tsMap[name] = append(tsMap[name], h*60+m)
	}
	for name, ts := range tsMap {
		sort.Ints(ts)
		for i, t := range ts {
			if sort.SearchInts(ts, t+61)-i >= 3 {
				ans = append(ans, name)
				break
			}
		}
	}
	sort.Strings(ans)
	return
}
