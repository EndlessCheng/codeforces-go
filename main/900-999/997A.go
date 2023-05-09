package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func CF997A(in io.Reader, out io.Writer) {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var x, y int
	var s string
	Fscan(bufio.NewReader(in), &s, &x, &y, &s)
	zero := strings.Count("1"+s, "10")
	if zero == 0 {
		Fprint(out, 0)
	} else {
		Fprint(out, min(x, y)*(zero-1)+y)
	}
}

//func main() { CF997A(os.Stdin, os.Stdout) }
