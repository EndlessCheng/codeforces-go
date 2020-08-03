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

	var n, d, ans, x, y int
	for Fscan(in, &n, &d); n > 0; n-- {
		if Fscan(in, &x, &y); x*x+y*y <= d*d {
			ans++
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
