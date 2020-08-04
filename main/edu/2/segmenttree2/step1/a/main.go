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

	var n, q, op, l, r, v int
	Fscan(in, &n, &q)
	tree := make([]int64, n+1)
	add := func(i, v int) {
		for ; i <= n; i += i & -i {
			tree[i] += int64(v)
		}
	}
	sum := func(i int) (s int64) {
		for ; i > 0; i &= i - 1 {
			s += tree[i]
		}
		return
	}
	update := func(l, r, v int) { add(l, v); add(r+1, -v) }
	for ; q > 0; q-- {
		if Fscan(in, &op, &l); op == 1 {
			Fscan(in, &r, &v)
			update(l+1, r, v)
		} else {
			Fprintln(out, sum(l+1))
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
