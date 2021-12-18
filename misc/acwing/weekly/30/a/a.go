package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, h, v, ans int
	for Fscan(in, &n, &h); n > 0; n-- {
		if Fscan(in, &v); v > h {
			ans += 2
		} else {
			ans++
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
