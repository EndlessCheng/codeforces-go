package main

// github.com/EndlessCheng/codeforces-go
func findSmallestSetOfVertices(n int, edges [][]int) (ans []int) {
	in := make([]bool, n)
	for _, e := range edges {
		in[e[1]] = true
	}
	for i, b := range in {
		if !b {
			ans = append(ans, i)
		}
	}
	return
}
