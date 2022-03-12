package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1650A(in io.Reader, out io.Writer) {
	var T int
	var s, t string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &t)
		if len(s)&1 > 0 {
			for i := range s {
				if s[i] == t[0] && i&1 == 0 {
					Fprintln(out, "YES")
					continue o
				}
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { CF1650A(os.Stdin, os.Stdout) }
