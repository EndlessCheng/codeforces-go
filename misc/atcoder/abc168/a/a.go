package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	var n int
	Fscan(_r, &n)
	switch n % 10 {
	case 2, 4, 5, 7, 9:
		Fprint(_w, "hon")
	case 0, 1, 6, 8:
		Fprint(_w, "pon")
	default:
		Fprint(_w, "bon")
	}
}

func main() { run(os.Stdin, os.Stdout) }
