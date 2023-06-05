package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1838C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		for i := 1; i < n; i += 2 {
			for j := 1; j <= m; j++ {
				Fprint(out, i*m+j, " ")
			}
			Fprintln(out)
		}
		for i := 0; i < n; i += 2 {
			for j := 1; j <= m; j++ {
				Fprint(out, i*m+j, " ")
			}
			Fprintln(out)
		}
	}
}

//func main() { CF1838C(os.Stdin, os.Stdout) }
