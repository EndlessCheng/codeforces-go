package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1480A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		for i, b := range s {
			if i&1 == 0 {
				if b > 'a' {
					b = 'a'
				} else {
					b = 'b'
				}
			} else {
				if b < 'z' {
					b = 'z'
				} else {
					b = 'y'
				}
			}
			s[i] = b
		}
		Fprintln(out, string(s))
	}
}

//func main() { CF1480A(os.Stdin, os.Stdout) }
