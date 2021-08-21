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

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		for i := 2; i <= n; i++ {
			Fprint(out, i, " ")
		}
		Fprintln(out, 1)
	}
}

func main() { run(os.Stdin, os.Stdout) }
