package main

import (
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	a := make([]int, 4)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	Fprint(out, a[3]-a[2], a[3]-a[1], a[3]-a[0])
}

func main() { run(os.Stdin, os.Stdout) }
