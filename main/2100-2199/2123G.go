package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2123G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	const mx = 500001
	divisors := [mx][]int32{}
	for i := int32(1); i < mx; i++ {
		for j := i; j < mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}

	var T, n, m, q, op, j, x int
	pos := [mx]int{}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &q)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ds := divisors[m]
		cnt := make([]int, len(ds))
		for i, d := range ds {
			pos[d] = i
			d := int(d)
			for j := range n - 1 {
				if a[j]%d > a[j+1]%d {
					cnt[i]++
				}
			}
		}

		for range q {
			Fscan(in, &op, &j)
			if op == 1 {
				j--
				Fscan(in, &x)
				for i, d := range ds {
					d := int(d)
					if j > 0 && a[j-1]%d > a[j]%d {
						cnt[i]--
					}
					if j > 0 && a[j-1]%d > x%d {
						cnt[i]++
					}
					if j < n-1 && a[j]%d > a[j+1]%d {
						cnt[i]--
					}
					if j < n-1 && x%d > a[j+1]%d {
						cnt[i]++
					}
				}
				a[j] = x
			} else {
				g := gcd(j, m)
				if cnt[pos[g]]*g+a[n-1]%g < m {
					Fprintln(out, "YES")
				} else {
					Fprintln(out, "NO")
				}
			}
		}
	}
}

//func main() { cf2123G(bufio.NewReader(os.Stdin), os.Stdout) }
