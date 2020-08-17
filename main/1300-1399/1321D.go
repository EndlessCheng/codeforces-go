package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1321D(_r io.Reader, _w io.Writer) {
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
	rg := make([][]int, n)
	for ; m > 0; m-- {
		v, w := read()-1, read()-1
		g[v] = append(g[v], w)
		rg[w] = append(rg[w], v)
	}
	k := read()
	path := make([]int, k)
	for i := range path {
		path[i] = read() - 1
	}

	dep := make([]int, n)
	vis := make([]bool, n)
	vis[path[k-1]] = true
	q := []int{path[k-1]}
	for len(q) > 0 {
		var v int
		v, q = q[0], q[1:]
		for _, w := range rg[v] {
			if !vis[w] {
				vis[w] = true
				dep[w] = dep[v] + 1
				q = append(q, w)
			}
		}
	}

	min, ex := 0, 0
	for i, v := range path[:k-1] {
		next := path[i+1]
		if dep[v] <= dep[next] {
			min++
			continue
		}
		for _, w := range g[v] {
			if w != next && dep[w] == dep[next] {
				ex++
				break
			}
		}
	}
	Fprint(out, min, min+ex)
}

//func main() { CF1321D(os.Stdin, os.Stdout) }
