package main

import (
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF805B(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	Fprint(out, strings.Repeat("aabb", n/4)+"aabb"[:n&3])
}

//func main() { CF805B(os.Stdin, os.Stdout) }
