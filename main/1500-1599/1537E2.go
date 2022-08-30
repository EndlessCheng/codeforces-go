package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1537E2(in io.Reader, out io.Writer) {
	var n, k int
	var s string
	Fscan(bufio.NewReader(in), &n, &k, &s)
	sz := 1
	for i := range s {
		if s[i] > s[i%sz] {
			break
		}
		if s[i] < s[i%sz] {
			sz = i + 1
		}
	}
	Fprint(out, strings.Repeat(s[:sz], k/sz+1)[:k])
}

//func main() { CF1537E2(os.Stdin, os.Stdout) }
