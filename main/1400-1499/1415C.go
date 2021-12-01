package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1415C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, p, k, add, del int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &p, &k, &s, &add, &del)
		p--
		ans := int(1e9)
		for i := n - 1; i >= p && i >= n-k; i-- {
			res := (i - p) * del
			for j := i; j >= p; j -= k {
				if s[j] == '0' {
					res += add
				}
				if res < ans {
					ans = res
				}
				res -= k * del
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1415C(os.Stdin, os.Stdout) }
