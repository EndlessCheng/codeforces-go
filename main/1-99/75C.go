package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF75C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var l, r, q int
	Fscan(in, &l, &r, &q)
	g := gcd(l, r)
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
		if i := sort.SearchInts(ds, r+1); i > 0 && ds[i-1] >= l {
			Fprintln(out, ds[i-1])
		} else {
			Fprintln(out, -1)
		}
	}
}

//func main() { CF75C(os.Stdin, os.Stdout) }
