package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, v, max int
	for Fscan(in, &n); n > 0; n-- {
		if Fscan(in, &v); v > max {
			max = v
		}
	}
	Fprint(out, max^v)
}

func main() { run(os.Stdin, os.Stdout) }
