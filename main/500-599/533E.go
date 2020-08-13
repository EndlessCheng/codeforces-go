package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF533E(in io.Reader, out io.Writer) {
	var n, ans int
	var s, t string
	Fscan(bufio.NewReader(in), &n, &s, &t)
	for i := range s {
		if s[i] != t[i] {
			s, t = s[i:], t[i:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != t[i] {
			s, t = s[:i+1], t[:i+1]
			break
		}
	}
	if s[1:] == t[:len(t)-1] {
		ans++
	}
	if t[1:] == s[:len(s)-1] {
		ans++
	}
	Fprint(out, ans)
}

//func main() { CF533E(os.Stdin, os.Stdout) }
