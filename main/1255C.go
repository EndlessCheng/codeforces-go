package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1255C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	orderP3 := [6][3]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}

	var n int
	Fscan(in, &n)
	mp := make(map[[2]int][]int, 6*(n-2))
	q := make([][3]int, n-2)
	for i := range q {
		Fscan(in, &q[i][0], &q[i][1], &q[i][2])
		for j, order := range orderP3 {
			pair := [2]int{q[i][order[0]], q[i][order[1]]}
			mp[pair] = append(mp[pair], 6*i+j)
		}
	}

	g := make([][]int, 6*(n-2))
	inDeg := make([]int, len(g))
	for i, qi := range q {
		for j, order := range orderP3 {
			if ws, ok := mp[[2]int{qi[order[1]], qi[order[2]]}]; ok {
				v := 6*i + j
				for _, w := range ws {
					if w/6 != i {
						g[v] = append(g[v], w)
						inDeg[w]++
					}
				}
			}
		}
	}

	fa := make([]int, len(inDeg)) // init fa
	for i := range fa {
		fa[i] = -1
	}
	var end int
	queue := []int{}
	vOrder := make([]int, len(inDeg))
	for i, deg := range inDeg {
		if deg == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		var v int
		v, queue = queue[0], queue[1:]
		for _, w := range g[v] {
			inDeg[w]--
			if inDeg[w] == 0 {
				queue = append(queue, w)
				vOrder[w] = vOrder[v] + 1
				if vOrder[w] == n-3 {
					end = w // record end
				}
				fa[w] = v // record fa
			}
		}
	}
	vs := make([]int, 0, n-2)
	for v := end; v != -1; v = fa[v] {
		vs = append(vs, v) // record vs from end to start
	}
	// reverse vs
	v := vs[len(vs)-1]
	q0, order0 := q[v/6], orderP3[v%6]
	Fprint(out, q0[order0[0]], q0[order0[1]])
	for i := len(vs) - 1; i >= 0; i-- {
		v = vs[i]
		qv, orderV := q[v/6], orderP3[v%6]
		Fprint(out, " ", qv[orderV[2]])
	}
}

//func main() {
//	Sol1255C(os.Stdin, os.Stdout)
//}
