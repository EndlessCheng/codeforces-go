package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF901A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var h, p2 int
	Fscan(in, &h)
	c := make([]int, h+1)
	for i := range c {
		Fscan(in, &c[i])
		if i > 0 && c[i] > 1 && c[i-1] > 1 {
			p2 = i
		}
	}

	if p2 == 0 {
		Fprint(out, "perfect")
		return
	}

	Fprintln(out, "ambiguous")
	v := 0
	for _, c := range c {
		for j := 0; j < c; j++ {
			Fprint(out, v, " ")
		}
		v += c
	}
	Fprintln(out)
	v = 0
	for i, c := range c {
		j := 0
		if i == p2 {
			Fprint(out, v-1, " ")
			j = 1
		}
		for ; j < c; j++ {
			Fprint(out, v, " ")
		}
		v += c
	}
}

//func main() { CF901A(os.Stdin, os.Stdout) }
