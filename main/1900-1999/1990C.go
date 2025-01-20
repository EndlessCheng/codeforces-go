package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1990C(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		vis := make([]int8, n+1)
		ans, mx := 0, 0
		for i := range a {
			Fscan(in, &v)
			ans += v
			if vis[v] == 1 {
				mx = max(mx, v)
			} else {
				vis[v] = 1
			}
			a[i] = mx
		}

		mx = 0
		for i, v := range a {
			if vis[v] == 2 {
				mx = max(mx, v)
			} else {
				vis[v] = 2
			}
			ans += v + mx*(n-i)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1990C(bufio.NewReader(os.Stdin), os.Stdout) }
