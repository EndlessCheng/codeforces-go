package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1716A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n == 1 {
			Fprintln(out, 2)
		} else {
			Fprintln(out, (n+2)/3)
		}
	}
}

//func main() { CF1716A(os.Stdin, os.Stdout) }
