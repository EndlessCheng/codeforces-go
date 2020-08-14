package main

import (
	"bufio"
	. "fmt"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func CF1146C() {
	out := bufio.NewWriter(os.Stdout)
	var qs [10][2][]int
	var f func(d, tp, l, r int)
	f = func(d, tp, l, r int) {
		if d == 10 {
			return
		}
		for i := l; i < r; i++ {
			qs[d][tp] = append(qs[d][tp], i)
		}
		f(d+1, 0, l, (l+r)/2)
		f(d+1, 1, (l+r)/2, r)
	}
	var t, n int
	for Scan(&t); t > 0; t-- {
		qs = [10][2][]int{}
		Scan(&n)
		f(0, 0, 1, 1+n/2)
		f(0, 1, 1+n/2, n+1)
		ans := 0
		for _, q := range qs {
			if len(q[0]) == 0 || len(q[1]) == 0 {
				break
			}
			Fprint(out, len(q[0]), len(q[1]))
			for _, a := range q {
				for _, v := range a {
					Fprint(out, " ", v)
				}
			}
			Fprintln(out)
			out.Flush()
			Scan(&n)
			if n > ans {
				ans = n
			}
		}
		Fprintln(out, -1, ans)
		out.Flush()
	}
}

//func main() { CF1146C() }
