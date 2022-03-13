package main

// github.com/EndlessCheng/codeforces-go
func digArtifacts(_ int, artifacts [][]int, dig [][]int) (ans int) {
	type pair struct{ x, y int }
	d := map[pair]bool{}
	for _, p := range dig {
		d[pair{p[0], p[1]}] = true
	}
next:
	for _, art := range artifacts {
		for i := art[0]; i <= art[2]; i++ {
			for j := art[1]; j <= art[3]; j++ {
				if !d[pair{i, j}] {
					continue next
				}
			}
		}
		ans++
	}
	return
}
