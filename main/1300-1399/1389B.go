package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1389B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var T, n, k, mxL int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &mxL)
		k++
		a := make([]int, n)
		s := make([]int, n+1)
		for i := range a {
			Fscan(in, &a[i])
			s[i+1] = s[i] + a[i]
		}
		ans := s[k]
		for i := 1; i < k; i++ {
			for j := 1; j <= mxL && i+j*2 <= k; j++ {
				if i+j*2 < k {
					ans = max(ans, s[k-j*2]+(a[i]+a[i-1])*j)
				} else {
					ans = max(ans, s[i]+(a[i]+a[i-1])*j)
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1389B(os.Stdin, os.Stdout) }
