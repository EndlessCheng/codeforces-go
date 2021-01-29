package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, r int
	Fscan(in, &n, &r)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	s := int64(0)
	for _, v := range a {
		s += int64(sort.SearchInts(a, v-r))
	}
	Fprint(out, s)
}

func main() { run(os.Stdin, os.Stdout) }
