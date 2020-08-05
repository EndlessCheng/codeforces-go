package main

import (
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, x, y int
	Fscan(in, &n, &x, &y)
	if x > y {
		x, y = y, x
	}
	ans := x
	n--
	ans += sort.Search(n*x, func(t int) bool { return t/x+t/y >= n })
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
