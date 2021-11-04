package main

import (
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1004B(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	Fprint(out, strings.Repeat("10", n/2)+"1"[:n&1])
}

//func main() { CF1004B(os.Stdin, os.Stdout) }
