package main

import (
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF982A(in io.Reader, out io.Writer) {
	var n int
	var s string
	Fscan(in, &n, &s)
	s = "0" + s + "0"
	if strings.Contains(s, "000") || strings.Contains(s, "11") {
		Fprint(out, "No")
	} else {
		Fprint(out, "Yes")
	}
}

//func main() { CF982A(os.Stdin, os.Stdout) }
