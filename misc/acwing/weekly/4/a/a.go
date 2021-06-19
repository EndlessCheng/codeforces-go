package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	var n int
	var s string
	Fscan(bufio.NewReader(_r), &n, &s)
	a := strings.Count(s, "A")
	b := n - a
	if a == b {
		Fprint(out, "T")
	} else if a > b {
		Fprint(out, "A")
	} else {
		Fprint(out, "B")
	}
}

func main() { run(os.Stdin, os.Stdout) }
