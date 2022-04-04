package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1660A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, a, b int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b)
		if a == 0 {
			Fprintln(out, 1)
		} else {
			Fprintln(out, a+b*2+1)
		}
	}
}

//func main() { CF1660A(os.Stdin, os.Stdout) }
