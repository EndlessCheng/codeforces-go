package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1373F(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
		}
		var s, l, i int
		for ; i < n*2; i++ {
			cap, cost := b[i%n], a[i%n]
			if s+cap < cost {
				s = 0
				l = i + 1
			} else {
				s = min(s+cap-cost, cap)
			}
		}
		if i-l >= n {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1373F(bufio.NewReader(os.Stdin), os.Stdout) }
