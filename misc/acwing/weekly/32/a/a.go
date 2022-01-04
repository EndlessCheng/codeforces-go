package main

import (
	. "fmt"
	"io"
	"os"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	if strings.Contains(s, "1111111") || strings.Contains(s, "0000000") {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}
}

func main() { run(os.Stdin, os.Stdout) }
