package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1443E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 14
	fac := [mx + 1]int{1}
	for i := 1; i <= mx; i++ {
		fac[i] = fac[i-1] * i
	}

	var n, q, op, l, r, k int
	Fscanln(in, &n, &q)
	m := min(n, mx)
	s := make([]int, m+1)
	valid := make([]int, m+1)
	kthPermutation := func(k int) {
		for i := 1; i <= m; i++ {
			valid[i] = 1
		}
		for i := 1; i <= m; i++ {
			order := k/fac[m-i] + 1
			k %= fac[m-i]
			for j := 1; j <= m; j++ {
				order -= valid[j]
				if order == 0 {
					s[i] = s[i-1] + j
					valid[j] = 0
					break
				}
			}
		}
	}
	pre := func(r int) (res int) {
		d := max(n-mx, 0)
		mid := min(r, d)
		r = max(r-d, 0)
		return mid*(mid+1)/2 + s[r] + r*d
	}

	for range q {
		Fscanln(in, &op, &l, &r)
		if op == 1 {
			kthPermutation(k)
			Fprintln(out, pre(r)-pre(l-1))
		} else {
			k = (k + l) % fac[m]
		}
	}
}

//func main() { cf1443E(bufio.NewReader(os.Stdin), os.Stdout) }
