package main

import (
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1342B(_r io.Reader, _w io.Writer) {
	var t int
	var s string
	for Fscan(_r, &t); t > 0; t-- {
		Fscan(_r, &s)
		n := len(s)
		if strings.Count(s, "0") == n || strings.Count(s, "1") == n {
			Fprintln(_w, s)
		} else {
			Fprintln(_w, strings.Repeat("01", n))
		}
	}
}

//func main() { CF1342B(os.Stdin, os.Stdout) }
