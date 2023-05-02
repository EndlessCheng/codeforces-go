package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1720D2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	f := make([]map[int][2][2]int, 30)
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		for i := range f {
			f[i] = map[int][2][2]int{}
		}
		ans := 0
		Fscan(in, &n)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			mx := 0
			for k, m := range f {
				mx = max(mx, m[(v^i)>>(k+1)][v>>k&1^1][i>>k&1])
			}
			mx++
			ans = max(ans, mx)
			for k, m := range f {
				p := m[(v^i)>>(k+1)]
				p[i>>k&1][v>>k&1] = max(p[i>>k&1][v>>k&1], mx)
				m[(v^i)>>(k+1)] = p
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1720D2(os.Stdin, os.Stdout) }
