package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, v int
	g := []int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		v = -v
		j := sort.SearchInts(g, v+1)
		if j < len(g) {
			g[j] = v
		} else {
			g = append(g, v)
		}
	}
	Fprint(out, len(g))
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
