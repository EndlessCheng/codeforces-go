package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF665C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s []byte
	Fscan(in, &s)
	for i, b := range s {
		if i+1 < len(s) && b == s[i+1] {
			s[i+1] = b + 1
			if s[i+1] > 'z' {
				s[i+1] = 'a'
			}
			if i+2 < len(s) && s[i+1] == s[i+2] {
				s[i+1]++
				if s[i+1] > 'z' {
					s[i+1] = 'a'
				}
			}
		}
	}
	Fprint(out, string(s))
}

//func main() { CF665C(os.Stdin, os.Stdout) }
