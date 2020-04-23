package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type tuple struct{ l, r, d int }

	solve := func(_case int) {
		var n int
		Fscan(in, &n)
		a := make([]tuple, n)
		for i := range a {
			Fscan(in, &a[i].l, &a[i].r, &a[i].d)
		}
		p := sort.Search(1<<31, func(x int) bool {
			cnt := 0
			for _, t := range a {
				if t.l > x {
					continue
				}
				if t.r > x {
					t.r = x
				}
				cnt += 1 + (t.r-t.l)/t.d
			}
			return cnt&1 == 1
		})
		if p == 1<<31 {
			Fprintln(out, "There's no weakness.")
		} else {
			cnt := 0
			for _, t := range a {
				if t.l <= p && p <= t.r && (p-t.l)%t.d == 0 {
					cnt++
				}
			}
			Fprintln(out, p, cnt)
		}
	}

	var t int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		solve(_case)
	}
}

func main() { run(os.Stdin, os.Stdout) }
