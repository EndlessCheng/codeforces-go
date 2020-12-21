package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func exgcd(a, b int) (g, x, y int) {
	if b == 0 {
		return a, 1, 0
	}
	g, y, x = exgcd(b, a%b)
	y -= a / b * x
	return
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, a, c, b int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &c, &b)
		g, _, y := exgcd(a, b)
		if c%g != 0 {
			Fprintln(out, -1)
			continue
		}
		a /= g
		c /= g
		y = y * c % a
		if y > 0 {
			y -= a
		}
		Fprintln(out, -y)
	}
}

func main() { run(os.Stdin, os.Stdout) }
