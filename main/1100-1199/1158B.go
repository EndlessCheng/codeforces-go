package main

import (
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1158B(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	s := strings.Repeat("0", (n-k)/2) + "1"
	Fprint(out, strings.Repeat(s, n/len(s)), s[:n%len(s)])
}

//func main() { CF1158B(os.Stdin, os.Stdout) }
