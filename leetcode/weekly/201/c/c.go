package main

// github.com/EndlessCheng/codeforces-go
func maxNonOverlapping(a []int, target int) (ans int) {
	sum := make([]int, len(a)+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	prev := 0
	pos := map[int]int{}
	for i, s := range sum {
		if p, ok := pos[s-target]; ok && p >= prev {
			ans++
			prev = i
		}
		pos[s] = i
	}
	return
}
