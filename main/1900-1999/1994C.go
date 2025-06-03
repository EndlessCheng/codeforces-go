package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1994C(in io.Reader, out io.Writer) {
	var T, n, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ans := n * (n + 1) / 2
		f := make([]int, n+1)
		s := 0
		r := n - 1
		for l := n - 1; l >= 0; l-- {
			s += a[l]
			for s > x {
				s -= a[r]
				r--
			}
			if r < n-1 {
				f[l] = f[r+2] + 1
				ans -= f[l]
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1994C(bufio.NewReader(os.Stdin), os.Stdout) }
