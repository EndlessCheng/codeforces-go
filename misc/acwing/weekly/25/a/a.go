package main

import (
	. "fmt"
	"io"
	"os"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	const S = "I hate that I love that "
	n := 0
	Fscan(in, &n)
	s := strings.Repeat(S, n/2) + S[:n&1*12]
	Fprint(out, s[:len(s)-5]+"it")
}

func main() { run(os.Stdin, os.Stdout) }
