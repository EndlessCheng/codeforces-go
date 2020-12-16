package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF359D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ mi, g int }
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, v int
	Fscan(in, &n)
	st := make([][19]pair, n)
	for i := range st {
		Fscan(in, &v)
		st[i][0] = pair{v, v}
	}
	for j := 1; 1<<j <= n; j++ {
		for i := 0; i+1<<j <= n; i++ {
			p, q := st[i][j-1], st[i+1<<(j-1)][j-1]
			st[i][j] = pair{min(p.mi, q.mi), gcd(p.g, q.g)}
		}
	}
	query := func(l, r int) bool {
		k := bits.Len(uint(r-l)) - 1
		p, q := st[l][k], st[r-1<<k][k]
		return min(p.mi, q.mi) == gcd(p.g, q.g)
	}

	ans := sort.Search(n+1, func(sz int) bool {
		for r := sz; r <= n; r++ {
			if query(r-sz, r) {
				return false
			}
		}
		return true
	}) - 1
	ls := []int{}
	for r := ans; r <= n; r++ {
		if query(r-ans, r) {
			ls = append(ls, r-ans+1)
		}
	}
	Fprintln(out, len(ls), ans-1)
	for _, l := range ls {
		Fprint(out, l, " ")
	}
}

//func main() { CF359D(os.Stdin, os.Stdout) }
