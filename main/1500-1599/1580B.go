package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1580B(in io.Reader, out io.Writer) {
	var n, m, k, mod int
	Fscan(in, &n, &m, &k, &mod)
	const mx = 100
	F := [mx]int{1}
	for i := 1; i < mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	C := [mx][mx]int{}
	for i := 0; i < mx; i++ {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % mod
		}
	}

	memo := make([][][]int, m)
	for i := range memo {
		memo[i] = make([][]int, n+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, k+1)
			for p := range memo[i][j] {
				memo[i][j][p] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(dep, size, need int) int {
		if dep < 0 {
			if need > 0 {
				return 0
			}
			return F[size] // 随便排
		}
		if size == 0 {
			return 1
		}

		p := &memo[dep][size][need]
		if *p >= 0 {
			return *p
		}

		if dep == 0 { // 这是我们要找的
			need--
		}

		res := 0
		for leftSz := range size {
			for leftNeed := max(need-(size-1-leftSz), 0); leftNeed <= min(leftSz, need); leftNeed++ {
				leftRes := dfs(dep-1, leftSz, leftNeed)
				if leftRes == 0 { // 剪枝，右子树不递归
					continue
				}
				rightRes := dfs(dep-1, size-1-leftSz, need-leftNeed)
				res = (res + C[size-1][leftSz]*leftRes%mod*rightRes) % mod
			}
		}
		*p = res
		return res
	}
	Fprint(out, dfs(m-1, n, k))
}

//func main() { cf1580B(os.Stdin, os.Stdout) }
