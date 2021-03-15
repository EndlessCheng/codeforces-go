package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF475D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n, x, q int
	Fscan(in, &n)
	cnt := map[int]int64{}
	type pair struct{ v, l, r int }
	set := []pair{}
	for i := 0; i < n; i++ {
		Fscan(in, &x)
		for j, p := range set {
			set[j].v = gcd(p.v, x)
		}
		set = append(set, pair{x, i, i + 1})
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
		for _, p := range set {
			cnt[p.v] += int64(p.r - p.l)
		}
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &x)
		Fprintln(out, cnt[x])
	}
}

//func main() { CF475D(os.Stdin, os.Stdout) }
