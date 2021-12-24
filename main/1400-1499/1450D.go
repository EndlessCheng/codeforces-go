package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1450D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		c := make([]int, n)
		ans := bytes.Repeat([]byte{'0'}, n)
		ans[0] = '1'
		for i := range a {
			Fscan(in, &a[i])
			a[i]--
			c[a[i]]++
			if c[a[i]] > 1 {
				ans[0] = '0'
			}
		}
		l, r := 0, n-1
		for i, c := range c {
			if c == 0 {
				break
			}
			ans[n-1-i] = '1'
			if c > 1 || a[l] != i && a[r] != i {
				break
			}
			if a[l] == i {
				l++
			} else {
				r--
			}
		}
		Fprintf(out, "%s\n", ans)
	}
}

//func main() { CF1450D(os.Stdin, os.Stdout) }
