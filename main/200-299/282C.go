package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF282C(in io.Reader, out io.Writer) {
	var s, t string
	Fscan(bufio.NewReader(in), &s, &t)
	if len(s) == len(t) && strings.Contains(s, "1") == strings.Contains(t, "1") {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}
}

//func main() { CF282C(os.Stdin, os.Stdout) }
