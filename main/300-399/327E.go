package main

import (
	. "fmt"
	"io"
)

// 另一种效率更高的做法是折半枚举，见 https://codeforces.com/blog/entry/8274

// github.com/EndlessCheng/codeforces-go
func CF327E(in io.Reader, out io.Writer) {
	const mod int = 1e9 + 7
	var n, k int
	Fscan(in, &n)
	m := 1 << n
	sum := make([]int, m)
	for i := 0; i < n; i++ {
		Fscan(in, &sum[1<<i])
	}
	b := [2]int{}
	Fscan(in, &k, &b[0], &b[1])
	f := make([]int, m)
	f[0] = 1
	for s := 1; s < m; s++ {
		lb := s & -s
		v := sum[s^lb] + sum[lb]
		if v > 1e9+1 {
			v = 1e9 + 1
		}
		sum[s] = v
		if v != b[0] && v != b[1] {
			dv := 0
			for sub, lb := s, 0; sub > 0; sub ^= lb {
				lb = sub & -sub
				dv += f[s^lb]
				if dv >= mod {
					dv -= mod
				}
			}
			f[s] = dv
		}
	}
	Fprint(out, f[1<<n-1])
}

//func main() { CF327E(os.Stdin, os.Stdout) }
