package main

// github.com/EndlessCheng/codeforces-go
func minimumRounds(tasks []int) (ans int) {
	cnt := map[int]int{}
	for _, task := range tasks {
		cnt[task]++
	}
	for _, c := range cnt {
		if c == 1 { return -1 }
		ans += (c + 2) / 3
	}
	return
}
