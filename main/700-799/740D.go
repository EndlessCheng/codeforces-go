package main

import (
	"bufio"
	. "fmt"
	"io"
)

// TODO https://www.luogu.com.cn/blog/qiyue7ACM/solution-cf739b
// 也可以用一个全局栈记录当前 DFS 链，然后在链上二分。这样做更简单但是适用范围较窄

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
	type edge struct{ to, l int }

	n := read()
	a := make([]int64, n)
	for i := range a {
		a[i] = int64(read())
	}
	const mx = 19
	pa := make([][mx]int, n)
	pa[0][0] = -1
	g := make([][]edge, n)
	for w := 1; w < n; w++ {
		v := read() - 1
		g[v] = append(g[v], edge{w, read()})
		pa[w][0] = v
	}
	for k := 0; k+1 < mx; k++ {
		for v := range pa {
			if p := pa[v][k]; p != -1 {
				pa[v][k+1] = pa[p][k]
			} else {
				pa[v][k+1] = -1
			}
		}
	}

	dep := make([]int64, n)
	var f func(int, int64)
	f = func(v int, d int64) {
		dep[v] = d
		for _, e := range g[v] {
			f(e.to, d+int64(e.l))
		}
	}
	f(0, 0)

	diff := make([]int, n)
	for v, val := range a {
		u := v
		dv := dep[v]
		for i := mx - 1; i >= 0; i-- {
			if p := pa[u][i]; p != -1 && dv-dep[p] <= val {
				u = p
			}
		}
		diff[v]++
		diff[u]--
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
