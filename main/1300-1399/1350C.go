package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1350C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	const mx int = 2e5
	lpf := [mx + 1]int{1: 1}
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
	var n, v int
	Fscan(in, &n)
	ps := [mx][]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		for v > 1 {
			p, e := lpf[v], 1
			for v /= p; lpf[v] == p; v /= p {
				e++
			}
			ps[p] = append(ps[p], e)
		}
	}
	ans := int64(1)
	for p, es := range ps {
		if len(es) >= n-1 {
			sort.Ints(es)
			for e := es[len(es)-n+1]; e > 0; e-- {
				ans *= int64(p)
			}
		}
	}
	Fprint(_w, ans)
}

//func main() { CF1350C(os.Stdin, os.Stdout) }
