package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf853D(in io.Reader, out io.Writer) {
	var n, v, s, c1 int
	Fscan(in, &n)
	for range n {
		Fscan(in, &v)
		v /= 1000
		s += v
		c1 += 2 - v
	}
	off := s * 10 / 11
	if s <= 11 {
		off = s - v
	}
	if c1 == 0 && off&1 > 0 {
		off--
	}
	Fprint(out, s*1000-off*100)
}

//func main() { cf853D(bufio.NewReader(os.Stdin), os.Stdout) }
