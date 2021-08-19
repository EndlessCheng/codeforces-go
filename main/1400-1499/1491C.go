package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1491C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		fa := make([]int, n+1)
		var find func(int) int
		find = func(x int) int {
			if fa[x] != x {
				fa[x] = find(fa[x])
			}
			return fa[x]
		}
		a := make([]int, n)
		for i := range a {
			fa[i] = i
			if Fscan(in, &a[i]); a[i] == 1 {
				fa[i]++
			}
		}
		fa[n] = n
		ans := int64(0)
		for i, s := range a[:n-1] {
			if s == 1 {
				continue
			}
			if i+s >= n {
				to := n - 1 - i
				ans += int64(s - to)
				s = to
			}
			for ; s > 1; s-- {
				ans++
				for j := i + s; j < n; {
					if a[j] > 1 {
						if a[j]--; a[j] == 1 {
							fa[j] = j + 1
						}
						j += a[j] + 1
					} else {
						j = find(j)
					}
				}
			}
		}
		Fprintln(out, ans+int64(a[n-1]-1))
	}
}

//func main() { CF1491C(os.Stdin, os.Stdout) }
