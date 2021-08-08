package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1063A(in io.Reader, out io.Writer) {
	var n int
	var s []byte
	Fscan(bufio.NewReader(in), &n, &s)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	Fprintf(out, "%s", s)
}

//func main() { CF1063A(os.Stdin, os.Stdout) }
