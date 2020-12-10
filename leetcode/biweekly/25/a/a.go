package main

// github.com/EndlessCheng/codeforces-go
func kidsWithCandies(a []int, extraCandies int) (ans []bool) {
	mx := a[0]
	for _, v := range a {
		if v > mx {
			mx = v
		}
	}
	for _, v := range a {
		ans = append(ans, v+extraCandies >= mx)
	}
	return
}
