package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF721C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	var maxT int32
	Fscan(in, &n, &m, &maxT)
	es := make([][3]int32, m)
	for i := range es {
		Fscan(in, &es[i][0], &es[i][1], &es[i][2])
	}

	const mx = 5001
	f := make([][mx]int32, n)
	for i := range f {
		for j := 1; j <= n; j++ {
			f[i][j] = maxT + 1
		}
	}
	f[0][1] = 0
	from := make([][mx]int16, n+1)

	// 不需要拓扑，直接跑 n-1 次就行了
	ans := 0
	for i := 1; i < n; i++ {
		for _, e := range es {
			v, w, t := e[0], e[1], e[2]
			sumT := f[i-1][v] + t
			if sumT < f[i][w] {
				f[i][w] = sumT
				from[i][w] = int16(v)
			}
		}
		if f[i][n] <= maxT {
			ans = i
		}
	}

	Fprintln(out, ans+1)
	path := make([]any, ans+1)
	v := int16(n)
	for i := ans; i >= 0; i-- {
		path[i] = v
		v = from[i][v]
	}
	Fprint(out, path...)
}

//func main() { CF721C(os.Stdin, os.Stdout) }
