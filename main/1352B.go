package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1352B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, k int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &k)
		if k > n {
			Fprintln(out, "NO")
			continue
		}
		ans := []interface{}{}
		left := n - k
		if left&1 == 0 {
			ans = append(ans, left+1)
			for k--; k > 0; k-- {
				ans = append(ans, 1)
			}
		} else {
			if 2*k>n {
				Fprintln(out, "NO")
				continue
			}
			left = n - 2*k
			if left&1 == 0 {
				ans = append(ans, left+2)
				for k--; k > 0; k-- {
					ans = append(ans, 2)
				}
			} else {
				Fprintln(out, "NO")
				continue
			}
		}
		Fprintln(out, "YES")
		Fprintln(out, ans...)
	}
}

//func main() { CF1352B(os.Stdin, os.Stdout) }
