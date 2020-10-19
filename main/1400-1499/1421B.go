package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1421B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]string, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := [][2]int{}
		if a[0][1] == a[1][0] {
			if a[n-2][n-1] == a[0][1] {
				ans = append(ans, [2]int{n - 2, n - 1})
			}
			if a[n-1][n-2] == a[0][1] {
				ans = append(ans, [2]int{n - 1, n - 2})
			}
		} else if a[n-2][n-1] == a[n-1][n-2] {
			if a[n-2][n-1] == a[0][1] {
				ans = append(ans, [2]int{0, 1})
			}
			if a[n-2][n-1] == a[1][0] {
				ans = append(ans, [2]int{1, 0})
			}
		} else {
			ans = append(ans, [2]int{0, 1})
			if a[n-2][n-1] == a[1][0] {
				ans = append(ans, [2]int{n - 2, n - 1})
			} else {
				ans = append(ans, [2]int{n - 1, n - 2})
			}
		}
		Fprintln(out, len(ans))
		for _, p := range ans {
			Fprintln(out, p[0]+1, p[1]+1)
		}
	}
}

//func main() { CF1421B(os.Stdin, os.Stdout) }
