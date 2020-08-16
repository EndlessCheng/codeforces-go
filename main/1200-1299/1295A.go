package main

import (
	"bufio"
	. "fmt"
	"io"
	. "strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1295A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		if n&1 == 0 {
			Fprintln(out, Repeat("1", n/2))
		} else {
			Fprintln(out, "7"+Repeat("1", (n-3)/2))
		}
	}
}

//func main() {
//	CF1295A(os.Stdin, os.Stdout)
//}
