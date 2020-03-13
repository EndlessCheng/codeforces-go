package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF740D(_r io.Reader, _w io.Writer) {
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

	n := read()
	a := make([]int64, n)
	for i := range a {
		a[i] = int64(read())
	}
	type pair struct{ to, l int }
	g := make([][]pair, n)
	for w := 1; w < n; w++ {
		v := read() - 1
		g[v] = append(g[v], pair{w, read()})
	}

	const mx = 19
	pa := make([][mx]int, n)
	dep := make([]int64, n)
	var f func(int, int64)
	f = func(v int, d int64) {
		dep[v] = d
		for _, e := range g[v] {
			w := e.to
			pa[w][0] = v
			f(w, d+int64(e.l))
		}
	}
	pa[0][0] = -1
	f(0, 0)
	for k := 0; k+1 < mx; k++ {
		for v := range pa {
			if p := pa[v][k]; p != -1 {
				pa[v][k+1] = pa[p][k]
			} else {
				pa[v][k+1] = -1
			}
		}
	}

	diff := make([]int, n)
	for v, val := range a {
		down := v
	outer:
		for {
			dv := dep[v]
			for i, p := range pa[v] {
				if p == -1 || dv-dep[p] > val {
					if i == 0 {
						break outer
					}
					break
				}
				v = p
			}
			val -= dv - dep[v]
		}
		diff[v]--
		diff[down]++
	}

	ans := make([]interface{}, n)
	var f2 func(int) int
	f2 = func(v int) (sum int) {
		for _, e := range g[v] {
			sum += f2(e.to)
		}
		ans[v] = sum
		return sum + diff[v]
	}
	f2(0)
	Fprint(out, ans...)
}

//func main() { CF740D(os.Stdin, os.Stdout) }
