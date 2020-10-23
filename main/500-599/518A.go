package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF518A(in io.Reader, out io.Writer) {
	var s []byte
	var t string
	Fscan(in, &s, &t)
	for i := len(s) - 1; ; i-- {
		if s[i] < 'z' {
			s[i]++
			break
		}
		s[i] = 'a'
	}
	if string(s) == t {
		Fprint(out, "No such string")
	} else {
		Fprint(out, string(s))
	}
}

//func main() { CF518A(os.Stdin, os.Stdout) }
