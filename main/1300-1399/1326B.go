package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1326B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, max int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		v += max
		Fprint(out, v, " ")
		if v > max {
			max = v
		}
	}
}

//func main() { CF1326B(os.Stdin, os.Stdout) }
