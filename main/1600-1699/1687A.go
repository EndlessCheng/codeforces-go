package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1687A(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		s := 0
		for i := range a {
			Fscan(in, &a[i])
			if i < k {
				s += a[i]
			}
		}
		if k >= n {
			Fprintln(out, s+k*n-n*(n+1)/2)
			continue
		}
		ans := s
		for i, v := range a[k:] {
			s += v - a[i]
			ans = max(ans, s)
		}
		Fprintln(out, ans+k*(k-1)/2)
	}
}

//func main() { cf1687A(bufio.NewReader(os.Stdin), os.Stdout) }
