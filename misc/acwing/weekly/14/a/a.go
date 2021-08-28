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

	var T, l1, r1, l2, r2 int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l1, &r1, &l2, &r2)
		a, b := l1, l2
		if a == b {
			if b < r2 {
				b++
			} else {
				b--
			}
		}
		Fprintln(out, a, b)
	}
}

func main() { run(os.Stdin, os.Stdout) }
