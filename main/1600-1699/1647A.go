package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1647A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		if Fscan(in, &n); n%3 == 0 {
			Fprintln(out, strings.Repeat("21", n/3))
		} else if n%3 == 1 {
			Fprintln(out, strings.Repeat("12", n/3)+"1")
		} else {
			Fprintln(out, strings.Repeat("21", n/3)+"2")
		}
	}
}

//func main() { CF1647A(os.Stdin, os.Stdout) }
