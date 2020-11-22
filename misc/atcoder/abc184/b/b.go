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

	var n, x int
	var s []byte
	Fscan(in, &n, &x, &s)
	for _, b := range s {
		if b == 'o' {
			x++
		} else if x > 0 {
			x--
		}
	}
	Fprint(out, x)
}

func main() { run(os.Stdin, os.Stdout) }
