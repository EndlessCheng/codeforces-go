package main

// 若 $(a_i,b_i,c_i)$ 各部分均不超过 $(x,y,z)$，则可以执行更新操作
// 因此对所有满足要求的 $(a_i,b_i,c_i)$，只要各部分都出现了 $x,y,z$，最终各部分就能更新成 $x,y,z$

// github.com/EndlessCheng/codeforces-go
func mergeTriplets(a [][]int, t []int) bool {
	found := [3]bool{}
	for _, p := range a {
		if p[0] <= t[0] && p[1] <= t[1] && p[2] <= t[2] {
			found[0] = found[0] || p[0] == t[0]
			found[1] = found[1] || p[1] == t[1]
			found[2] = found[2] || p[2] == t[2]
		}
	}
	return found[0] && found[1] && found[2]
}