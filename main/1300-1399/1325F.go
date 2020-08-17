package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF1325F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	sq := int(math.Ceil(math.Sqrt(float64(n))))
	dep := make([]int, n)
	cycle := []interface{}{}
	outSet := make([]bool, n)
	var f func(v int) bool
	f = func(v int) bool {
		cycle = append(cycle, v+1)
		dep[v] = len(cycle)
		for _, w := range g[v] {
			if dep[w] == 0 {
				if f(w) {
					return true
				}
			} else if dep[v]-dep[w]+1 >= sq {
				Fprintln(out, 2)
				Fprintln(out, dep[v]-dep[w]+1)
				Fprint(out, cycle[dep[w]-1:dep[v]]...)
				return true
			}
		}
		if !outSet[v] {
			for _, w := range g[v] {
				outSet[w] = true
			}
		}
		cycle = cycle[:len(cycle)-1]
		return false
	}
	if f(0) {
		return
	}
	Fprintln(out, 1)
	for i := 0; sq > 0; i++ {
		if !outSet[i] {
			Fprint(out, i+1, " ")
			sq--
		}
	}
}

//func main() { CF1325F(os.Stdin, os.Stdout) }
