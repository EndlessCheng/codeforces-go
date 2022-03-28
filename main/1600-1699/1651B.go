package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1651B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n > 19 {
			Fprintln(out, "NO")
			continue
		}
		Fprintln(out, "YES")
		for v := 1; n > 0; n-- {
			Fprint(out, v, " ")
			v *= 3
		}
		Fprintln(out)
	}
}

//func main() { CF1651B(os.Stdin, os.Stdout) }
