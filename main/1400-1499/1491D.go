package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1491D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var q, v, w int
o:
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &v, &w)
		if v > w {
			Fprintln(out, "NO")
			continue
		}
		for i, c := 0, 0; i < 30; i++ {
			c += v>>i&1 - w>>i&1
			if c < 0 {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1491D(os.Stdin, os.Stdout) }
