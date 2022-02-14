package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1093D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	const mx int = 3e5
	pow2 := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		pow2[i] = pow2[i-1] << 1 % mod
	}

	var T, n, m, v, w int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		g := make([][]int, n)
		for ; m > 0; m-- {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		col := make([]int8, n)
		cnt := [3]int{}
		var f func(int, int8) bool
		f = func(v int, c int8) bool {
			cnt[c]++
			col[v] = c
			for _, w := range g[v] {
				if col[w] == c || col[w] == 0 && !f(w, 3^c) {
					return false
				}
			}
			return true
		}
		ans := int64(1)
		for i, c := range col {
			if c == 0 {
				cnt = [3]int{}
				if !f(i, 1) {
					Fprintln(out, 0)
					continue o
				}
				ans = ans * (pow2[cnt[1]] + pow2[cnt[2]]) % mod
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1093D(os.Stdin, os.Stdout) }
