package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1822D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n == 1 {
			Fprintln(out, 1)
		} else if n%2 > 0 {
			Fprintln(out, -1)
		} else {
			for i := 0; i < n; i += 2 {
				Fprint(out, n-i, i+1, " ")
			}
			Fprintln(out)
		}
	}
}

//func main() { CF1822D(os.Stdin, os.Stdout) }
