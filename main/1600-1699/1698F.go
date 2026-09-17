package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1698F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+1)
		b := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}
		for i := 1; i <= n; i++ {
			Fscan(in, &b[i])
		}

		ans := [][2]int{}
		rev := func(l, r int) {
			ans = append(ans, [2]int{l, r})
			slices.Reverse(a[l : r+1])
		}
		check := func(x int) bool {
			for i := x + 1; i < n; i++ {
				if a[i] == b[x] && a[i+1] == a[x-1] {
					rev(x-1, i+1)
					return true
				}
			}
			return false
		}
		f := func(x int) bool {
			if check(x) {
				return true
			}
			for i := x + 1; i <= n; i++ {
				if a[i] == b[x] && a[i-1] == a[x-1] {
					for j := x - 1; j < i; j++ {
						for k := i; k <= n; k++ {
							if a[j] == a[k] {
								rev(j, k)
								return check(x)
							}
						}
					}
				}
			}
			return false
		}

		if a[1] != b[1] || a[n] != b[n] {
			Fprintln(out, "NO")
			continue
		}
		for i := 2; i < n; i++ {
			if a[i] != b[i] && !f(i) {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
		Fprintln(out, len(ans))
		for _, p := range ans {
			Fprintln(out, p[0], p[1])
		}
	}
}

//func main() { cf1698F(bufio.NewReader(os.Stdin), os.Stdout) }
