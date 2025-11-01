package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2110F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		ans, mx := 0, 0
		for i := range a {
			Fscan(in, &a[i])
			x := a[i]
			if x >= mx*2 {
				m := 0
				for _, y := range a[:i] {
					m = max(m, x%y+y)
				}
				ans = m
			} else {
				ans = max(ans, x%mx+mx%x)
			}
			Fprint(out, ans, " ")
			mx = max(mx, x)
		}
		Fprintln(out)
	}
}

//func main() { cf2110F(bufio.NewReader(os.Stdin), os.Stdout) }
