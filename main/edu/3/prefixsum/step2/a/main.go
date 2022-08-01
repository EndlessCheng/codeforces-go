package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, l, r int
	Fscan(in, &n)
	a := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		a[i] += a[i-1]
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &l, &r)
		Fprintln(out, a[r]-a[l-1])
	}
}

func main() { run(os.Stdin, os.Stdout) }
