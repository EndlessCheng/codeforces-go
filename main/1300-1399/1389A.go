package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1389A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l, &r)
		if 2*l > r {
			Fprintln(out, -1, -1)
		} else {
			Fprintln(out, l, 2*l)
		}
	}
}

//func main() { CF1389A(os.Stdin, os.Stdout) }
