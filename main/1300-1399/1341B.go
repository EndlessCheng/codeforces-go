package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1341B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, k int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sum := make([]int, n+1)
		for i := 1; i+1 < n; i++ {
			sum[i+1] = sum[i]
			if a[i-1] < a[i] && a[i] > a[i+1] {
				sum[i+1]++
			}
		}
		sum[n] = sum[n-1]
		ans, l := 0, 0
		for i := 0; i+k <= n; i++ {
			s := sum[i+k] - sum[i]
			if sum[i+1] > sum[i] {
				s--
			}
			if sum[i+k] > sum[i+k-1] {
				s--
			}
			if s > ans {
				ans = s
				l = i
			}
		}
		Fprintln(out, ans+1, l+1)
	}
}

//func main() { CF1341B(os.Stdin, os.Stdout) }
