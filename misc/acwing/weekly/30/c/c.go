package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var a, b, l, r, q int
	Fscan(in, &a, &b, &q)
	g := gcd(a, b)
	ds := []int{}
	for d := 1; d*d <= g; d++ {
		if g%d == 0 {
			ds = append(ds, d)
			if d*d < g {
				ds = append(ds, g/d)
			}
		}
	}
	sort.Ints(ds)

	for ; q > 0; q-- {
		Fscan(in, &l, &r)
		d := ds[sort.SearchInts(ds, r+1)-1] // <= r 的最大因子
		if d >= l {
			Fprintln(out, d)
		} else {
			Fprintln(out, -1)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
