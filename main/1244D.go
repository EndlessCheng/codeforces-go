package main

import (
	"bufio"
	. "fmt"
	"io"
	)

// github.com/EndlessCheng/codeforces-go
func Sol1244D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	c := [3][]int64{}
	for i := range c {
		c[i] = make([]int64, n)
		for j := range c[i] {
			Fscan(in, &c[i][j])
		}
	}
	g := make([][]int, n)
	deg := make([]int, n)
	for m := n - 1; m > 0; m-- {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		deg[v]++
		deg[w]++
		if deg[v] > 2 || deg[w] > 2 {
			Fprint(out, -1)
			return
		}
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	vs := make([]int, 0, n)
	for i, d := range deg {
		if d == 1 {
			vs = append(vs, i, g[i][0])
			for prev, v := i, g[i][0]; len(g[v]) == 2; {
				for _, w := range g[v] {
					if w != prev {
						vs = append(vs, w)
						prev, v = v, w
						break
					}
				}
			}
			break
		}
	}
	minAns, orderIdxAns := int64(1e18), 0
	orders := [6][3]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}
	for i, order := range orders {
		sum := int64(0)
		for j, v := range vs {
			sum += c[order[j%3]][v]
		}
		if sum < minAns {
			minAns = sum
			orderIdxAns = i
		}
	}
	Fprintln(out, minAns)
	ans := make([]interface{}, n)
	order := orders[orderIdxAns]
	for i, v := range vs {
		ans[v] = order[i%3] + 1
	}
	Fprint(out, ans...)
}

//func main() {
//	Sol1244D(os.Stdin, os.Stdout)
//}
