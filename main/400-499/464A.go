package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF464A(in io.Reader, out io.Writer) {
	n, u, s := 0, byte(0), []byte{}
	Fscan(in, &n, &u, &s)
	u += 'a'
	i := n - 1
	s[i]++
	for i < n {
		if s[i] == u {
			if i == 0 {
				Fprint(out, "NO")
				return
			}
			s[i] = 'a'
			i--
			s[i]++
		} else if i > 0 && s[i] == s[i-1] || i > 1 && s[i] == s[i-2] {
			s[i]++
		} else {
			i++
		}
	}
	Fprintf(out, "%s", s)
}

//func main() { CF464A(os.Stdin, os.Stdout) }
