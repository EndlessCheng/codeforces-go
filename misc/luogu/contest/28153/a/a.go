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

	s := 50
	var v int
	for i := 0; i < 7; i++ {
		Fscan(in, &v)
		s += v
	}
	Fscan(in, &v)
	s += 5 * v
	if Fscan(in, &v); s >= v {
		Fprintln(out, "AKIOI")
	} else {
		Fprintln(out, "AFO")
	}
}

func main() { run(os.Stdin, os.Stdout) }
