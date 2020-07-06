package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w int64
	Fscan(in, &n)
	s := n * (n + 1) * (n + 2) / 6
	for m := n - 1; m > 0; m-- {
		Fscan(in, &v, &w)
		if v > w {
			v, w = w, v
		}
		s -= v * (n - w + 1)
	}
	Fprint(out, s)
}

func main() { run(os.Stdin, os.Stdout) }
