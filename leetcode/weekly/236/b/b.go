package main

// github.com/EndlessCheng/codeforces-go
func findTheWinner(n, k int) int {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	cur := 0
	for n--; n > 0; n-- {
		cur = (cur + k - 1) % len(a)
		a = append(a[:cur], a[cur+1:]...)
	}
	return a[0] + 1
}
