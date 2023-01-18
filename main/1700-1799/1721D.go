package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1721D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
		}
		ans := 0
		q := [][2][]int{{a, b}}
	o:
		for k := 29; k >= 0; k-- {
			tmp := q
			q = nil
			for _, p := range tmp {
				var f, g [2][]int
				for _, v := range p[0] {
					f[v>>k&1] = append(f[v>>k&1], v)
				}
				for _, v := range p[1] {
					g[v>>k&1] = append(g[v>>k&1], v)
				}
				if len(f[0]) != len(g[1]) {
					q = tmp
					continue o
				}
				if len(f[0]) > 0 {
					q = append(q, [2][]int{f[0], g[1]})
				}
				if len(f[1]) > 0 {
					q = append(q, [2][]int{f[1], g[0]})
				}
			}
			ans |= 1 << k
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1721D(os.Stdin, os.Stdout) }
