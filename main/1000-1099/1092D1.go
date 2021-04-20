package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1092D1(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	s := []int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		if v &= 1; len(s) > 0 && v == s[len(s)-1] {
			s = s[:len(s)-1]
		} else {
			s = append(s, v)
		}
	}
	if len(s) < 2 {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}
}

//func main() { CF1092D1(os.Stdin, os.Stdout) }
