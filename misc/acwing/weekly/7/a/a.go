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

	var T, h, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &h, &m)
		Fprintln(out, (24-h)*60-m)
	}
}

func main() { run(os.Stdin, os.Stdout) }
