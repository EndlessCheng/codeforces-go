package main

import (
	"bufio"
	. "fmt"
	"io"
)

// todo 2-SAT 可否？

// github.com/EndlessCheng/codeforces-go
func CF1400C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, x int
	var t []byte
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &t, &x)
		n := len(t)
		s := make([]byte, n)
		for i, b := range t {
			if b == '0' {
				if i >= x {
					s[i-x] = '0'
				}
				if i+x < n {
					s[i+x] = '0'
				}
			}
		}
		for i, b := range s {
			if b == 0 {
				s[i] = '1'
			}
		}
		for i, b := range t {
			if b == '1' && !(i >= x && s[i-x] == '1' || i+x < n && s[i+x] == '1') {
				Fprintln(out, -1)
				continue o
			}
		}
		Fprintln(out, string(s))
	}
}

//func main() { CF1400C(os.Stdin, os.Stdout) }
