package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1358A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, m int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m)
		Fprintln(out, (n*m+1)/2)
	}
}

//func main() { CF1358A(os.Stdin, os.Stdout) }
