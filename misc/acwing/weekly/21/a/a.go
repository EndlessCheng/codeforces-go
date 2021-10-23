package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var T, k, l int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &k, &l)
		for v := k; v != l; v *= k {
			if v > l {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

func main() { run(os.Stdin, os.Stdout) }
