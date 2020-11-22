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

	var a, b, c, d int
	Fscan(in, &a, &b, &c, &d)
	Fprint(out, a*d-b*c)
}

func main() { run(os.Stdin, os.Stdout) }
