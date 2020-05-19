package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var a, b, n int
	Fscan(in, &a, &b, &n)
	if n >= b {
		n = b - 1
	}
	Fprint(out, a*n/b-a*(n/b))
}

func main() { run(os.Stdin, os.Stdout) }
