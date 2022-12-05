package main

// https://space.bilibili.com/206214
func dividePlayers(skill []int) (ans int64) {
	total := 0
	cnt := map[int]int{}
	for _, x := range skill {
		total += x
		cnt[x]++
	}
	m := len(skill) / 2
	if total%m > 0 {
		return -1
	}
	s := total / m
	for x, c := range cnt {
		if c != cnt[s-x] {
			return -1
		}
		ans += int64(c * x * (s - x))
	}
	return ans / 2
}
