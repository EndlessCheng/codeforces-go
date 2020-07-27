package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, i, v, l, r int
	Fscan(in, &n, &q)
	tree := make([]int64, n+1)
	add := func(i, v int) {
		for ; i <= n; i += i & -i {
			tree[i] += int64(v)
		}
	}
	query := func(l, r int) (s int64) {
		for ; r > l; r &= r - 1 {
			s += tree[r]
		}
		for ; l > r; l &= l - 1 {
			s -= tree[l]
		}
		return
	}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		add(i+1, a[i])
	}
	for ; q > 0; q-- {
		if Fscan(in, &op); op == 1 {
			Fscan(in, &i, &v)
			add(i+1, v-a[i])
			a[i] = v
		} else {
			Fscan(in, &l, &r)
			Fprintln(out, query(l, r))
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
