package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2115A(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var T, n int
	vis := [5001]int{}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		allG, mn, c := 0, int(1e9), 0
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			allG = gcd(allG, a[i])
			if a[i] < mn {
				mn, c = a[i], 1
			} else if a[i] == mn {
				c++
			}
			vis[a[i]] = T
		}
		if mn == allG {
			Fprintln(out, n-c)
			continue
		}

		ans := n - 1
		q := slices.Clone(a)
	o:
		for {
			ans++
			tmp := q
			q = nil
			for _, v := range tmp {
				for _, w := range a {
					g := gcd(v, w)
					if g == allG {
						break o
					}
					if vis[g] != T {
						vis[g] = T
						q = append(q, g)
					}
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2115A(bufio.NewReader(os.Stdin), os.Stdout) }
