package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1973E(in io.Reader, out io.Writer) {
	f := func(x int) int {
		return x * (x - 1) / 2
	}
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		l, r := 0, 2*n
		s := map[int]bool{}
		for i := 1; i <= n; i++ {
			var x int
			Fscan(in, &x)
			if x != i {
				s[x+i] = true
				l = max(l, x+1)
				r = min(r, x+n)
			}
		}

		if len(s) == 0 {
			Fprintln(out, f(2*n+1))
			continue
		}

		ans := f(2*n) - f(l-1) - f(2*n-r)
		if len(s) == 1 {
			ans++
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1973E(bufio.NewReader(os.Stdin), os.Stdout) }
