package main

// github.com/EndlessCheng/codeforces-go
func minSubarray(a []int, p int) int {
	n := len(a)
	ans := n
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	pos := map[int][]int{}
	for i, v := range sum {
		v %= p
		pos[v] = append(pos[v], i)
	}
	for i, v := range sum {
		v = (p - (sum[n]-v)%p) % p
		for len(pos[v]) > 0 && pos[v][0] <= i {
			ans = min(ans, i-pos[v][0])
			pos[v] = pos[v][1:]
		}
	}
	if ans < n {
		return ans
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
