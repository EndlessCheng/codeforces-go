package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1556A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, a, b int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b)
		if a&1 != b&1 {
			Fprintln(out, -1)
		} else if a == 0 && b == 0 {
			Fprintln(out, 0)
		} else if a == b {
			Fprintln(out, 1)
		} else {
			Fprintln(out, 2)
		}
	}
}

//func main() { CF1556A(os.Stdin, os.Stdout) }
