package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1023A(in io.Reader, out io.Writer) {
	var n, m int
	var s, t string
	Fscan(bufio.NewReader(in), &n, &m, &s, &t)
	i := strings.IndexByte(s, '*')
	if i < 0 && s == t || i >= 0 && len(s)-1 <= len(t) && strings.HasPrefix(t, s[:i]) && strings.HasSuffix(t, s[i+1:]) {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}
}

//func main() { CF1023A(os.Stdin, os.Stdout) }
