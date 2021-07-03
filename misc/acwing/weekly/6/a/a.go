package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, v int
	c := [2]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		c[v&1]++
	}
	Fprint(out, c[c[1]&1])
}

func main() { run(os.Stdin, os.Stdout) }
