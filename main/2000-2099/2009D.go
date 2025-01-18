package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2009D(in io.Reader, out io.Writer) {
	var T, n, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		has := make([][2]bool, n+1)
		cnt := 0
		for range n {
			Fscan(in, &x, &y)
			has[x][y] = true
			if has[x][0] && has[x][1] {
				cnt++
			}
		}

		ans := 0
		for i := 1; i < n; i++ {
			if has[i][0] && has[i-1][1] && has[i+1][1] {
				ans++
			}
			if has[i][1] && has[i-1][0] && has[i+1][0] {
				ans++
			}
		}
		Fprintln(out, ans+cnt*(n-2))
	}
}

//func main() { cf2009D(bufio.NewReader(os.Stdin), os.Stdout) }
