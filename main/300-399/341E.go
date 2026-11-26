package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf341E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}

	nz := func() int {
		cnt := 0
		for i := 1; i <= n; i++ {
			if a[i] != 0 {
				cnt++
			}
		}
		return cnt
	}

	if nz() < 2 {
		Fprintln(out, -1)
		return
	}

	ans := [][2]int{}
	for {
		if nz() == 2 {
			break
		}

		p, q, r := 0, 0, 0
		for i := 1; i <= n; i++ {
			if a[i] != 0 && (p == 0 || a[i] < a[p]) {
				p = i
			}
		}

		for i := 1; i <= n; i++ {
			if i != p && a[i] != 0 {
				if q == 0 {
					q = i
				} else {
					r = i
				}
			}
		}

		if a[q] > a[r] {
			q, r = r, q
		}

		d := a[q] / a[p]
		for i := 0; 1<<i <= d; i++ {
			if d>>i&1 != 0 {
				ans = append(ans, [2]int{p, q})
				a[q] -= a[p]
			} else {
				ans = append(ans, [2]int{p, r})
				a[r] -= a[p]
			}
			a[p] <<= 1
		}
	}

	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprintln(out, v[0], v[1])
	}
}

//func main() { cf341E(bufio.NewReader(os.Stdin), os.Stdout) }
