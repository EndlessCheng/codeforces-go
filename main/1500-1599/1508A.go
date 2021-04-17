package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1508A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var a, b, c string
	var s []byte
	f := func(a, b string) {
		c0 := strings.Count(a, "0")
		if c0 >= n != (strings.Count(b, "0") >= n) {
			return
		}
		tar := byte('0')
		if c0 < n {
			tar = '1'
		}
		s = nil
		i, j := 0, 0
		for i < 2*n && j < 2*n {
			if a[i] == b[j] {
				s = append(s, a[i])
				i++
				j++
			} else if a[i] != tar {
				s = append(s, a[i])
				i++
			} else {
				s = append(s, b[j])
				j++
			}
		}
		s = append(s, a[i:]...)
		s = append(s, b[j:]...)
	}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &a, &b, &c)
		f(a, b)
		f(a, c)
		f(b, c)
		Fprintln(out, string(s))
	}
}

//func main() { CF1508A(os.Stdin, os.Stdout) }
