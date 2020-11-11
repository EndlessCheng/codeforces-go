package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF246B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, s int
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		s += v
	}
	if s%n == 0 {
		Fprint(out, n)
	} else {
		Fprint(out, n-1)
	}
}

//func main() { CF246B(os.Stdin, os.Stdout) }
