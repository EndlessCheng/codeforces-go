package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF482B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	qs := make([]struct{ l, r, q int }, m)
	for i := range qs {
		Fscan(in, &qs[i].l, &qs[i].r, &qs[i].q)
		qs[i].l--
	}

	ans := make([]int, n)
	for mask := 1; mask < 1<<30; mask <<= 1 {
		d := make([]int, n+1)
		for _, q := range qs {
			if q.q&mask > 0 {
				d[q.l]++
				d[q.r]--
			}
		}
		p0 := -1
		for i, v := range d[:n] {
			d[i+1] += v
			if v > 0 {
				ans[i] |= mask
			} else {
				p0 = i
			}
			d[i] = p0 // 复用一下节省空间，存储左侧最近 0 的位置
		}
		for _, q := range qs {
			if q.q&mask == 0 && d[q.r-1] < q.l {
				Fprint(out, "NO")
				return
			}
		}
	}
	Fprintln(out, "YES")
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF482B(os.Stdin, os.Stdout) }
