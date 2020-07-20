package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1385B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, v int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		vis := make([]bool, n+1)
		for n *= 2; n > 0; n-- {
			Fscan(in, &v)
			if !vis[v] {
				vis[v] = true
				Fprint(out, v, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { CF1385B(os.Stdin, os.Stdout) }
