package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1626B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		for i := n - 1; i > 0; i-- {
			if v := s[i]&15 + s[i-1]&15; v > 9 {
				s[i-1] = '1'
				s[i] = '0' + v%10
				goto o
			}
		}
		s[0] += s[1] & 15
		Fprintf(out, "%s", s[:1])
		s = s[2:]
	o:
		Fprintf(out, "%s\n", s)
	}
}

//func main() { CF1626B(os.Stdin, os.Stdout) }
