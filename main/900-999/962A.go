package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF962A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, tot, s int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		tot += a[i]
	}
	tot = (tot + 1) / 2
	for i, v := range a {
		s += v
		if s >= tot {
			Fprint(out, i+1)
			return
		}
	}
}

//func main() { CF962A(os.Stdin, os.Stdout) }
