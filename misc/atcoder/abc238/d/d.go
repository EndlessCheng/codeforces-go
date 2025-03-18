package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, and, s int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &and, &s)
		or := s - and
		if or >= 0 && or&and == and {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
