package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf147B(in io.Reader, out io.Writer) {
	var n, m, ans int
	Fscan(in, &n, &m)
	mat := func() [][]int {
		a := make([][]int, n)
		for i := range a {
			a[i] = make([]int, n)
			for j := range a[i] {
				a[i][j] = -1e9
			}
		}
		return a
	}

	const mx = 9
	f := [mx][][]int{}
	for i := range f {
		f[i] = mat()
	}
	// 至多 vs 恰好，初始化 f[0][i][0] = 0 是至多 2^k 条边，不写就是恰好 2^k 条边
	for v, r := range f[0] {
		r[v] = 0
	}
	for range m {
		var v, w, wt, wt2 int
		Fscan(in, &v, &w, &wt, &wt2)
		v--
		w--
		f[0][v][w] = max(f[0][v][w], wt)
		f[0][w][v] = max(f[0][w][v], wt2)
	}

	for i := 1; i < mx; i++ {
		for v := range n {
			for w := range n {
				for k := range n {
					f[i][v][w] = max(f[i][v][w], f[i-1][v][k]+f[i-1][k][w])
				}
			}
		}
	}

	cur := mat()
	for i, r := range cur {
		r[i] = 0
	}
o:
	for i := mx - 1; i >= 0; i-- {
		res := mat()
		for v := range n {
			for w := range n {
				for k := range n {
					res[v][w] = max(res[v][w], cur[v][k]+f[i][k][w])
				}
			}
		}
		for v, r := range res {
			if r[v] > 0 {
				continue o
			}
		}
		ans |= 1 << i
		cur = res
	}
	Fprint(out, (ans+1)%(1<<mx))
}

//func main() { cf147B(bufio.NewReader(os.Stdin), os.Stdout) }
