package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1620C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var k, x int64
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &x, &s)
		if x == 1 {
			Fprintln(out, strings.ReplaceAll(s, "*", ""))
			continue
		}
		x--
		d := []int{}
		for i, c := n-1, 0; i >= 0; i-- {
			if s[i] == '*' {
				c++
				if i == 0 || s[i-1] != '*' {
					k := k*int64(c) + 1
					d = append(d, int(x%k))
					x /= k
					c = 0
				}
			}
		}
		ans := []byte{}
		j := len(d) - 1
		for i, b := range s {
			if b == 'a' {
				ans = append(ans, 'a')
			} else if i == n-1 || s[i+1] == 'a' {
				ans = append(ans, bytes.Repeat([]byte{'b'}, d[j])...)
				j--
			}
		}
		Fprintf(out, "%s\n", ans)
	}
}

//func main() { CF1620C(os.Stdin, os.Stdout) }
