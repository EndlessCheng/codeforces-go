package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2137E(in io.Reader, out io.Writer) {
	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		cnt := make([]int, n+1)
		for range n {
			Fscan(in, &v)
			cnt[v]++
		}

		m := 0
		for cnt[m] == 1 {
			m++
		}

		mex := m
		for cnt[mex] > 0 {
			mex++
		}

		if k == 1 {
			ans := mex * n
			for i, c := range cnt[:mex] {
				if c == 1 {
					ans += i - mex
				}
			}
			Fprintln(out, ans)
		} else if mex > m {
			Fprintln(out, m*(m-1)/2+(m+k%2)*(n-m))
		} else if n-m == 1 {
			Fprintln(out, n*(n-1)/2)
		} else {
			Fprintln(out, m*(m-1)/2+(m+1-k%2)*(n-m))
		}
	}
}

//func main() { cf2137E(bufio.NewReader(os.Stdin), os.Stdout) }
