package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF538B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	Fscan(in, &s)
	mx := 0
	for _, b := range s {
		if int(b&15) > mx {
			mx = int(b & 15)
		}
	}
	Fprintln(out, mx)
	for i := 0; i < mx; i++ {
		p := false
		for _, b := range s {
			if int(b&15) > i {
				Fprint(out, 1)
				p = true
			} else if p {
				Fprint(out, 0)
			}
		}
		Fprint(out, " ")
	}
}

//func main() { CF538B(os.Stdin, os.Stdout) }
