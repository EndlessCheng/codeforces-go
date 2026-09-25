package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf643E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var q, op, v int
	Fscan(in, &q)
	fa := make([]int, q+2)
	f := make([][]float64, q+2)
	for i := range f {
		f[i] = make([]float64, 50)
	}
	id := 1
	for range q {
		Fscan(in, &op, &v)
		if op == 1 {
			id++
			fa[id] = v
			t1, t2 := 0., 1.
			now := v
			for i := 1; i < 50 && now != 0; i++ {
				t := f[now][i]
				f[now][i] = 1 - (1-f[now][i])/(1-t1/2)*(1-t2/2)
				t1 = t
				t2 = f[now][i]
				now = fa[now]
			}
		} else {
			ans := 0.
			for i := 1; i < 50; i++ {
				ans += f[v][i]
			}
			Fprintf(out, "%.6f\n", ans)
		}
	}
}

//func main() { cf643E(bufio.NewReader(os.Stdin), os.Stdout) }
