package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1371D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		c, k := k/n, k%n
		if k == 0 {
			Fprintln(out, 0)
		} else {
			Fprintln(out, 2)
		}
		for i := 0; i < n; i++ {
			s := bytes.Repeat([]byte{'0'}, n)
			for d := 0; d < c; d++ {
				s[(i+d)%n] = '1'
			}
			if k > 0 {
				s[(i+c)%n] = '1'
				k--
			}
			Fprintln(out, string(s))
		}
	}
}

//func main() { CF1371D(os.Stdin, os.Stdout) }
