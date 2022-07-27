package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF494A(in io.Reader, out io.Writer) {
	var s string
	Fscan(bufio.NewReader(in), &s)
	cnt, d, minD := 0, 0, len(s)
	for _, b := range s {
		if b == '(' {
			d++
			continue
		}
		d--
		if d < 0 {
			Fprint(out, -1)
			return
		}
		if b == '#' {
			cnt++
			minD = d
		} else if d < minD {
			minD = d
		}
	}
	if minD < d {
		Fprint(out, -1)
	} else {
		Fprint(out, strings.Repeat("1\n", cnt-1), d+1)
	}
}

//func main() { CF494A(os.Stdin, os.Stdout) }
