package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type item struct{ l, r, d int }

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]item, n)
		for i := range a {
			Fscan(in, &a[i].l, &a[i].r, &a[i].d)
		}
		p := sort.Search(1<<31, func(x int) bool {
			cnt := 0
			for _, p := range a {
				if p.l > x {
					continue
				}
				if p.r > x {
					p.r = x
				}
				cnt ^= (p.r-p.l)/p.d&1 ^ 1
			}
			return cnt == 1
		})
		if p == 1<<31 {
			Fprintln(out, "There's no weakness.")
			continue
		}
		cnt := 0
		for _, t := range a {
			if t.l <= p && p <= t.r && (p-t.l)%t.d == 0 {
				cnt++
			}
		}
		Fprintln(out, p, cnt)
	}
}

func main() { run(os.Stdin, os.Stdout) }
