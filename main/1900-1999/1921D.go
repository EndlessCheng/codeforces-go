package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1921D(in io.Reader, out io.Writer) {
	var T, n, m int
	f := func(n int) []int {
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}
		slices.Sort(a[1:])
		for i := 2; i <= n; i++ {
			a[i] += a[i-1]
		}
		return a
	}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		sa, sb := f(n), f(m)
		ans := 0
		for i := range n + 1 {
			ans = max(ans, sb[m]-sb[m-i]-sb[n-i]+sa[n]-sa[i]*2)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1921D(bufio.NewReader(os.Stdin), os.Stdout) }
