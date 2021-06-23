package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF52A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, v int
	c := [4]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		c[v]++
	}
	Fprint(out, min(c[1]+min(c[2], c[3]), c[2]+c[3]))
}

//func main() { CF52A(os.Stdin, os.Stdout) }
