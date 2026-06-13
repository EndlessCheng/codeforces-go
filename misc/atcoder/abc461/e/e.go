package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
type fenwick []int

func (t fenwick) update(i, v int) {
	for ; i < len(t); i += i & -i {
		t[i] += v
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, i, ans int
	Fscan(in, &n, &q)
	tRow := make([]int, n+1)
	tCol := make([]int, n+1)
	row := make(fenwick, q+1)
	col := make(fenwick, q+1)
	for t := q; t > 0; t-- {
		Fscan(in, &op, &i)
		if op == 1 {
			lastT := tRow[i]
			tRow[i] = t
			if lastT == 0 {
				ans += n
			} else {
				ans += col.pre(lastT)
				row.update(lastT, -1)
			}
			row.update(t, 1)
		} else {
			lastT := tCol[i]
			tCol[i] = t
			if lastT == 0 {
				lastT = q
			} else {
				col.update(lastT, -1)
			}
			ans -= row.pre(lastT)
			col.update(t, 1)
		}
		Fprintln(out, ans)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
