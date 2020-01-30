package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1295B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	moveToRange := func(st, d, l, r int) (firstPos int, ok bool) {
		switch {
		case st < l:
			if d < 0 {
				return
			}
			return l + ((st-l)%d+d)%d, true
		case st <= r:
			return st, true
		default:
			if d > 0 {
				return
			}
			return r + ((st-r)%d+d)%d, true
		}
	}

	var t, n, x int
	var s []byte
	solve := func() (ans int) {
		cnts := map[int]int{0: 1}
		d, l, r := 0, 0, 0
		for i, b := range s {
			if b == '0' {
				d++
				if d > r {
					r = d
				}
			} else {
				d--
				if d < l {
					l = d
				}
			}
			if i+1 < n {
				cnts[d]++
			}
		}
		if d == 0 {
			if c, ok := cnts[x]; ok && c > 0 {
				return -1
			}
		} else if x, ok := moveToRange(x, -d, l, r); ok {
			for ; x >= l && x <= r; x -= d {
				ans += cnts[x]
			}
		}
		return
	}
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &x, &s)
		Fprintln(out, solve())
	}
}

//func main() {
//	CF1295B(os.Stdin, os.Stdout)
//}
