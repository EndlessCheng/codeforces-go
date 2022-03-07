package main

import (
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF112A(in io.Reader, out io.Writer) {
	var s, t string
	Fscan(in, &s, &t)
	Fprint(out, strings.Compare(strings.ToLower(s), strings.ToLower(t)))
}

//func main() { CF112A(os.Stdin, os.Stdout) }
