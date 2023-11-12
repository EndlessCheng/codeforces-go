package main

import "slices"

// https://space.bilibili.com/206214
func findHighAccessEmployees(accessTimes [][]string) (ans []string) {
	groups := map[string][]int{}
	for _, p := range accessTimes {
		name, s := p[0], p[1]
		t := int(s[0]&15*10+s[1]&15)*60 + int(s[2]&15*10+s[3]&15)
		groups[name] = append(groups[name], t)
	}
	for name, a := range groups {
		slices.Sort(a)
		for i := 2; i < len(a); i++ {
			if a[i]-a[i-2] < 60 {
				ans = append(ans, name)
				break
			}
		}
	}
	return
}
