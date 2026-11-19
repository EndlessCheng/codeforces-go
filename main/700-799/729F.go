package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf729F(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	s := make([]int32, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] += s[i-1]
	}

	var memo [4000][177][89]int32
	var vis [4000][177][89]bool
	var dfs func(int, int, int) int32
	dfs = func(i, d, k int) (res int32) {
		if i*2-d+k > n {
			return
		}
		if vis[i][d+88][k] {
			return memo[i][d+88][k]
		}
		vis[i][d+88][k] = true

		r := n - i + d
		if d <= 0 {
			res = s[i+k] - s[i] - dfs(i+k, d+k, k)
			if i+k+1 <= r {
				res = max(res, s[i+k+1]-s[i]-dfs(i+k+1, d+k+1, k+1))
			}
		} else {
			res = s[r] - s[r-k] - dfs(i, d-k, k)
			if r-k-1 >= i {
				res = max(res, s[r]-s[r-k-1]-dfs(i, d-k-1, k+1))
			}
		}

		memo[i][d+88][k] = res
		return
	}

	Fprint(out, dfs(0, 0, 1))
}

//func main() { cf729F(bufio.NewReader(os.Stdin), os.Stdout) }
