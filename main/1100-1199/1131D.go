package main

import (
	"bufio"
	. "fmt"
	"io"
)

func Sol1131D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	fa := make([]int, n+m)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) { fa[find(from)] = find(to) }

	cp := make([]string, n)
	for i := range cp {
		Fscan(in, &cp[i])
		for j, c := range cp[i] {
			if c == '=' {
				merge(n+j, i)
			}
		}
	}
	for i := range fa {
		find(i)
	}

	g := make([][]int, n+m)
	inDeg := make([]int, n+m)
	for i, s := range cp {
		for j, c := range s {
			if c == '=' {
				continue
			}
			v, w := fa[i], fa[n+j]
			if v == w {
				Fprint(out, "No")
				return
			}
			if c == '>' {
				v, w = w, v
			}
			g[v] = append(g[v], w)
			inDeg[w]++
		}
	}
	queue := []int{}
	ans := make([]int, n+m)
	cnt := 0
	for i, deg := range inDeg {
		if deg == 0 {
			queue = append(queue, i)
			ans[i] = 1
			cnt++
		}
	}
	for len(queue) > 0 {
		var v int
		v, queue = queue[0], queue[1:]
		for _, w := range g[v] {
			inDeg[w]--
			if inDeg[w] == 0 {
				queue = append(queue, w)
				ans[w] = ans[v] + 1
				cnt++
			}
		}
	}
	if cnt < n+m {
		Fprint(out, "No")
		return
	}
	Fprintln(out, "Yes")
	for i, o := range fa {
		if i == n {
			Fprintln(out)
		}
		Fprint(out, ans[o], " ")
	}
}

//func main() {
//	Sol1131D(os.Stdin, os.Stdout)
//}
