package main

// https://space.bilibili.com/206214
func totalNumbers(digits []int) int {
	set := map[int]struct{}{}
	for i, a := range digits {
		if a%2 > 0 {
			continue
		}
		for j, b := range digits {
			if j == i {
				continue
			}
			for k, c := range digits {
				if c == 0 || k == i || k == j {
					continue
				}
				set[c*100+b*10+a] = struct{}{}
			}
		}
	}
	return len(set)
}
