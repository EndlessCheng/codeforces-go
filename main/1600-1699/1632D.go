package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1632D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	type result struct{ v, l, r int }
	set := []result{}
	var n, v, ans int
	Fscan(in, &n)
o:
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		if v == 1 {
			set = nil
			ans++
			Fprint(out, ans, " ")
			continue
		}
		for j, p := range set {
			g := gcd(p.v, v)
			if l := i + 1 - g; p.l <= l && l < p.r {
				set = nil
				ans++
				Fprint(out, ans, " ")
				continue o
			}
			set[j].v = g
		}
		set = append(set, result{v, i, i + 1})
		k := 0
		for _, q := range set[1:] {
			if set[k].v != q.v {
				k++
				set[k] = q
			} else {
				set[k].r = q.r
			}
		}
		set = set[:k+1]
		Fprint(out, ans, " ")
	}
}

//func main() { CF1632D(os.Stdin, os.Stdout) }
