package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF733D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ sz, i int }
	sort3 := func(a ...int) (x, y, z int) { sort.Ints(a); return a[0], a[1], a[2] }

	var n, a, b, c, max int
	var mxID []interface{}
	Fscan(in, &n)
	mp := map[[2]int]pair{}
	for i := 1; i <= n; i++ {
		Fscan(in, &a, &b, &c)
		a, b, c = sort3(a, b, c)
		if a > max {
			max = a
			mxID = []interface{}{i}
		}
		for _, sz := range [3][3]int{{a, b, c}, {a, c, b}, {b, c, a}} {
			a, b, c := sz[0], sz[1], sz[2]
			p, ok := mp[[2]int{a, b}]
			if p.i == i {
				continue
			}
			if ok {
				if a, _, _ := sort3(a, b, c+p.sz); a > max {
					max = a
					mxID = []interface{}{p.i, i}
				}
			}
			if c > p.sz {
				mp[[2]int{a, b}] = pair{c, i}
			}
		}
	}
	Fprintln(out, len(mxID))
	Fprint(out, mxID...)
}

//func main() {
//	CF733D(os.Stdin, os.Stdout)
//}
