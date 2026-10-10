package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1834F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([]int, n+1)
	diff := [2][]int{}
	for i := range diff {
		diff[i] = make([]int, n*2+1)
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}

	do := func(r int) {
		for i := 1; i <= n; i++ {
			if i <= a[i] {
				if a[i]-i+1 < n-i+1 {
					diff[r][a[i]-i+1]++
					diff[r][n-i+1]--
				}
			} else {
				diff[r][0]++
				diff[r][n-i+1]--
				diff[r][n-i+a[i]+1]++
			}
		}
		sum := 0
		for i := range len(diff[r]) {
			sum += diff[r][i]
			diff[r][i] = sum
		}
	}

	do(0)
	slices.Reverse(a[1:])
	do(1)

	var q, r, p, t, k int
	Fscan(in, &q)
	Fprintln(out, diff[r][p])
	for range q {
		Fscan(in, &t)
		if t != 3 {
			Fscan(in, &k)
		}
		if t == 1 {
			p += n - k
		} else if t == 2 {
			p += k
		} else {
			r ^= 1
			p = n - p
		}
		p %= n
		Fprintln(out, diff[r][p])
	}
}

//func main() { cf1834F(bufio.NewReader(os.Stdin), os.Stdout) }
