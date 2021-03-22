package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1483A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		ans := make([]int, m)
		one := make([]int, n+1)
		pos := make([][]int, n+1)
		mx := 0
		for i := 0; i < m; i++ {
			Fscan(in, &k)
			for j := 0; j < k; j++ {
				Fscan(in, &v)
				if k == 1 {
					one[v]++
					ans[i] = v
				}
				pos[v] = append(pos[v], i)
				if len(pos[v]) > len(pos[mx]) {
					mx = v
				}
			}
		}
		c := one[mx]
		if c > (m+1)/2 {
			Fprintln(out, "NO")
			continue
		}
		for _, p := range pos[mx] {
			if c == (m+1)/2 {
				break
			}
			if ans[p] == 0 {
				ans[p] = mx
				c++
			}
		}
		for i, ps := range pos {
			if i != mx {
				for _, p := range ps {
					if ans[p] == 0 {
						ans[p] = i
					}
				}
			}
		}
		Fprintln(out, "YES")
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1483A(os.Stdin, os.Stdout) }
