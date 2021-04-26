package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF476D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	Fprintln(out, (n*6-1)*k)
	for i := 0; i < 6*n; i += 6 {
		Fprintln(out, (i+1)*k, (i+2)*k, (i+3)*k, (i+5)*k)
	}
}

//func main() { CF476D(os.Stdin, os.Stdout) }
