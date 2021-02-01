package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1149A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v int
	c := [3]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		c[v]++
	}
	if c[1] == 0 {
		Fprint(out, strings.Repeat("2 ", c[2]))
	} else if c[2] == 0 {
		Fprint(out, strings.Repeat("1 ", c[1]))
	} else {
		Fprint(out, "2 1 ", strings.Repeat("2 ", c[2]-1), strings.Repeat("1 ", c[1]-1))
	}
}

//func main() { CF1149A(os.Stdin, os.Stdout) }
