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

	var n, m int
	Fscan(in, &n, &m)
	used := make([]bool, n)
	for i := 1; i <= m; i++ {
		if 2*(n+1-2*i) != n && !used[2*i-1] {
			used[n+1-2*i] = true
			Fprintln(out, i, n+1-i)
		} else {
			used[n-2*i] = true
			Fprintln(out, i, n-i)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
