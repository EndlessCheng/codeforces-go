package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1352C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, k int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &k)
		Fprintln(out, sort.Search(2.1e9, func(x int) bool { return x-x/n >= k }))
	}
}

//func main() { CF1352C(os.Stdin, os.Stdout) }
