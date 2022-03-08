package main

import (
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF281A(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	Fprint(out, strings.Title(s))
}

//func main() { CF281A(os.Stdin, os.Stdout) }
