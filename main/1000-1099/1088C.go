package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1088C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v int
	Fscan(in, &n)
	Fprintln(out, n)
	Fprintln(out, 1, n, int(5e5))
	for i := 1; i < n; i++ {
		Fscan(in, &v)
		Fprintln(out, 2, i, v+5e5-i)
	}
}

//func main() { CF1088C(os.Stdin, os.Stdout) }
