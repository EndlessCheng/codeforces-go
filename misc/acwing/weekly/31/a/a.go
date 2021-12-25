package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var v, l, r, ans int
	x := int(1e9)
	for i := 0; i < 4; i++ {
		Fscan(in, &v)
		x = min(x, v)
	}
	Fscan(in, &l, &r)
	r = min(r, x-1)
	if r >= l {
		ans = r - l + 1
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
