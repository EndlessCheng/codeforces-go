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

	var T, a, b, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &k)
		Fprintln(out, k/2*(a-b)+k&1*a)
	}
}

func main() { run(os.Stdin, os.Stdout) }
