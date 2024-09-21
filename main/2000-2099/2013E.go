package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2013E(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s := map[int]bool{}
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			s[v] = true
		}
		ans := 0
		for ; len(s) > 1; n-- {
			mn := int(1e9)
			for v := range s {
				mn = min(mn, v)
			}
			ans += mn
			delete(s, mn)
			ns := map[int]bool{}
			for v := range s {
				ns[gcd(v, mn)] = true
			}
			s = ns
		}
		for v = range s {}
		Fprintln(out, ans+n*v)
	}
}

//func main() { cf2013E(bufio.NewReader(os.Stdin), os.Stdout) }
