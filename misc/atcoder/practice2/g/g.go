package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for ; m > 0; m-- {
		var v, w int
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
	}

	allScc := [][]any{}
	dfn := make([]int, len(g))
	dfsClock := 0
	st := []int{}
	var tarjan func(int) int
	tarjan = func(v int) int {
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		st = append(st, v)
		for _, w := range g[v] {
			if dfn[w] == 0 {
				lowW := tarjan(w)
				lowV = min(lowV, lowW)
			} else {
				lowV = min(lowV, dfn[w])
			}
		}
		if dfn[v] == lowV {
			scc := []any{}
			for {
				w := st[len(st)-1]
				st = st[:len(st)-1]
				dfn[w] = math.MaxInt
				scc = append(scc, w)
				if w == v {
					break
				}
			}
			allScc = append(allScc, scc)
		}
		return lowV
	}
	for i, ts := range dfn {
		if ts == 0 {
			tarjan(i)
		}
	}
	Fprintln(out, len(allScc))
	for i := len(allScc) - 1; i >= 0; i-- {
		Fprint(out, len(allScc[i]), " ")
		Fprintln(out, allScc[i]...)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
