package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(_r io.Reader, _w io.Writer) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n, m := read(), read()
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		v, w := read()-1, read()-1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	isCut := make([]bool, n)
	dfsClock := 0
	pre := make([]int, n) // 值从 1 开始
	//low := make([]int, n)
	var f func(v, fa int) int
	f = func(v, fa int) int {
		dfsClock++
		pre[v] = dfsClock
		lowV := dfsClock
		childCnt := 0
		for _, w := range g[v] {
			if pre[w] == 0 {
				childCnt++
				lowW := f(w, v)
				if lowW >= pre[v] { // 该子树没有反向边能连回 v 的祖先
					isCut[v] = true
				}
				lowV = min(lowV, lowW)
			} else if w != fa && pre[w] < pre[v] { // 找到反向边，用来更新 lowV
				lowV = min(lowV, pre[w])
			}
		}
		if fa == -1 && childCnt == 1 { // 特判：只有一个儿子的树根，删除后并没有增加连通分量的个数
			isCut[v] = false
		}
		//low[v] = lowV
		return lowV
	}
	for v, timestamp := range pre {
		if timestamp == 0 {
			f(v, -1)
		}
	}
	vs := make([]interface{}, 0, n)
	for v, is := range isCut {
		if is {
			vs = append(vs, v+1)
		}
	}
	Fprintln(out, len(vs))
	Fprintln(out, vs...)
}

func main() {
	solve(os.Stdin, os.Stdout)
}
