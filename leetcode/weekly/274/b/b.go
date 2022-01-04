package main

// github.com/EndlessCheng/codeforces-go
func numberOfBeams(bank []string) (ans int) {
	pre := 0
	for _, row := range bank {
		cnt := 0
		for _, ch := range row {
			cnt += int(ch & 1)
		}
		if cnt > 0 {
			ans += pre * cnt
			pre = cnt
		}
	}
	return
}
