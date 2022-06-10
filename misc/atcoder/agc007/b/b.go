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

	var n, p int
	Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &p)
		a[p-1] = n*p + i
	}
	for _, v := range a {
		Fprint(out, v, " ")
	}
	Fprintln(out)
	for v := n * n; v > 0; v -= n {
		Fprint(out, v, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
