package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF246A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	if n < 3 {
		Fprint(out, -1)
	} else {
		for i := n; i > 0; i-- {
			Fprint(out, i, " ")
		}
	}
}

//func main() { CF246A(os.Stdin, os.Stdout) }
