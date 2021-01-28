package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n+m)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
