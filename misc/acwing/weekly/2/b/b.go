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

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		if n == 1 {
			Fprintln(out, 0)
			continue
		}
		sort.Ints(a)
		for i := n - 2; k > 0; i-- {
			a[n-1] += a[i]
			k--
		}
		Fprintln(out, a[n-1])
	}
}

func main() { run(os.Stdin, os.Stdout) }
