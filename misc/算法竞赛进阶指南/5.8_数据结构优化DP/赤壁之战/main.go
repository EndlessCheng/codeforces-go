package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n, m int
	Fscan(in, &T)
	for t := 1; t <= T; t++ {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		b := append([]int(nil), a...)
		sort.Ints(b)
		for i, v := range a {
			a[i] = sort.SearchInts(b, v) + 2
		}

		const mod int = 1e9 + 7
		tree := make([]int, n+2)
		add := func(i, val int) {
			for ; i < n+2; i += i & -i {
				tree[i] = (tree[i] + val) % mod
			}
		}
		sum := func(i int) (res int) {
			for ; i > 0; i &= i - 1 {
				res = (res + tree[i]) % mod
			}
			return
		}

		dp := make([][]int, m+1)
		for i := range dp {
			dp[i] = make([]int, n)
		}
		for i := 1; i <= m; i++ {
			tree = make([]int, n+2)
			if i == 1 {
				add(1, 1)
			}
			for j, v := range a {
				dp[i][j] = sum(v - 1)
				add(v, dp[i-1][j])
			}
		}

		ans := 0
		for _, v := range dp[m] {
			ans = (ans + v) % mod
		}
		Fprintf(out, "Case #%d: %d\n", t, ans)
	}
}

func main() { run(os.Stdin, os.Stdout) }
