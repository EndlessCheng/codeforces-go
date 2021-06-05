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
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		for x := 0; 3*x <= n; x++ {
			for y := 0; 3*x+5*y <= n; y++ {
				if d := n - 3*x - 5*y; d%7 == 0 {
					Fprintln(out, x, y, d/7)
					continue o
				}
			}
		}
		Fprintln(out, -1)
	}
}

func main() { run(os.Stdin, os.Stdout) }
