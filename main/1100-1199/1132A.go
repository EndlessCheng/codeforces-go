package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1132A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var a, b, c, d int
	Fscan(in, &a, &b, &c, &d)
	if a != d || c > 0 && a*d == 0 {
		Fprint(out, 0)
	} else {
		Fprint(out, 1)
	}
}

//func main() { CF1132A(os.Stdin, os.Stdout) }
