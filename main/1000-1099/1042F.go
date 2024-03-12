package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
var k42, ans42 int
var g42 [][]int

func dfs42(v, fa int) (mx int) {
	if len(g42[v]) == 1 {
		return 1
	}
	for _, w := range g42[v] {
		if w == fa {
			continue
		}
		d := dfs42(w, v)
		if mx+d <= k42 {
			mx = max(mx, d)
		} else {
			ans42++
			mx = min(mx, d)
		}
	}
	if mx > 0 {
		mx++
	}
	return
}

func cf1042F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w, rt int
	Fscan(in, &n, &k42)
	g42 = make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g42[v] = append(g42[v], w)
		g42[w] = append(g42[w], v)
	}
	for i, vs := range g42 {
		if len(vs) > 1 {
			rt = i
			break
		}
	}
	d := dfs42(rt, -1)
	if d > 0 {
		ans42++
	}
	Fprint(out, ans42)
}

//func main() { cf1042F(os.Stdin, os.Stdout) }
