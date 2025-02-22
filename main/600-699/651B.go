package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf651B(in io.Reader, out io.Writer) {
	cnt := [1001]int{}
	var n, v int
	Fscan(in, &n)
	for range n {
		Fscan(in, &v)
		cnt[v]++
	}
	Fprint(out, n-slices.Max(cnt[:]))
}

//func main() { cf651B(os.Stdin, os.Stdout) }
