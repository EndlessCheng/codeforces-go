package main

import (
	"bufio"
	. "fmt"
	"io"
)

type fenwick27 []int

func (f fenwick27) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick27) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func cf627B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k, up1, up2, q, op, i, x int
	Fscan(in, &n, &k, &up2, &up1, &q)
	a := make([]int, n+1)
	b := make([]int, n+1)
	f1 := make(fenwick27, n+1)
	f2 := make(fenwick27, n+1)
	for ; q > 0; q-- {
		Fscan(in, &op, &i)
		if op == 1 {
			Fscan(in, &x)
			d := min(x, up1-a[i])
			a[i] += d
			f1.update(i, d)
			d = min(x, up2-b[i])
			b[i] += d
			f2.update(n+1-i, d)
		} else {
			Fprintln(out, f1.pre(i-1)+f2.pre(n+1-i-k))
		}
	}
}

//func main() { cf627B(bufio.NewReader(os.Stdin), os.Stdout) }
