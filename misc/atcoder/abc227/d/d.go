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
	var n, s, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
	}
	sort.Ints(a)
	for i := n - 1; ; i-- {
		if a[i] <= s/k {
			Fprint(out, s/k)
			return
		}
		s -= a[i]
		k--
	}
}

func main() { run(os.Stdin, os.Stdout) }
