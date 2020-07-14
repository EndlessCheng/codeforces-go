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

	var n, q, op, l, r, i int
	var v int64
	Fscan(in, &n, &q)
	tree := make([]int64, n+1)
	add := func(i int, val int64) {
		for ; i <= n; i += i & -i {
			tree[i] += val
		}
	}
	update := func(l, r int, val int64) { add(l, val); add(r+1, -val) }
	query := func(i int) (res int64) {
		for ; i > 0; i &= i - 1 {
			res += tree[i]
		}
		return
	}
	for ; q > 0; q-- {
		Fscan(in, &op)
		if op == 1 {
			Fscan(in, &l, &r, &v)
			update(l+1, r, v)
		} else {
			Fscan(in, &i)
			Fprintln(out, query(i+1))
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
