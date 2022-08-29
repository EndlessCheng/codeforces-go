package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1551D2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		ans := make([][]byte, n)
		for i := range ans {
			ans[i] = make([]byte, m)
		}
		if n&1 > 0 {
			k -= m / 2
			if k < 0 {
				Fprintln(out, "NO")
				continue
			}
			for j := 1; j < m; j += 2 {
				c := 'f' ^ byte(j/2&1)
				ans[n-1][j-1] = c
				ans[n-1][j] = c
			}
		} else if m&1 > 0 {
			if k > m/2*n {
				Fprintln(out, "NO")
				continue
			}
			for i := 1; i < n; i += 2 {
				c := 'f' ^ byte(i/2&1)
				ans[i-1][m-1] = c
				ans[i][m-1] = c
			}
		}
		if k&1 > 0 {
			Fprintln(out, "NO")
			continue
		}
		for i := 1; i < n; i += 2 {
			for j := 1; j < m; j += 2 {
				if k > 0 {
					k -= 2
					c := 'b' ^ byte(j/2&1)
					ans[i-1][j-1] = c
					ans[i-1][j] = c
					ans[i][j-1] = c ^ 1
					ans[i][j] = c ^ 1
				} else {
					c := 'd' ^ byte(i/2&1)
					ans[i-1][j-1] = c
					ans[i-1][j] = c ^ 1
					ans[i][j-1] = c
					ans[i][j] = c ^ 1
				}
			}
		}
		Fprintln(out, "YES")
		for _, r := range ans {
			Fprintf(out, "%s\n", r)
		}
	}
}

//func main() { CF1551D2(os.Stdin, os.Stdout) }
