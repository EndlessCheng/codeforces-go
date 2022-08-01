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

	var n int
	var s, v int64
	Fprint(out, 0)
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		s += v
		Fprint(out, " ", s)
	}
}

func main() { run(os.Stdin, os.Stdout) }
