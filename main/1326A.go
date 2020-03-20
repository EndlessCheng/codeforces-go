package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1326A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		if n == 1 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, "5"+strings.Repeat("3", n-1))
		}
	}
}

//func main() { CF1326A(os.Stdin, os.Stdout) }
