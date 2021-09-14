package main

import "strconv"

func getFolderNames(names []string) (ans []string) {
	t := ""
	nextID := map[string]int{}
	for _, s := range names {
		i := nextID[s]
		if i == 0 {
			nextID[s] = 1
			ans = append(ans, s)
			continue
		}
		for {
			t = s + "(" + strconv.Itoa(i) + ")"
			if nextID[t] == 0 {
				break
			}
			i++
		}
		nextID[s] = i + 1
		nextID[t] = 1
		ans = append(ans, t)
	}
	return
}
