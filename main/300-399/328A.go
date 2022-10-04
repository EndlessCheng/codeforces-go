package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF328A(in io.Reader, out io.Writer) {
	var a, b, c, d int
	Fscan(in, &a, &b, &c, &d)
	if a+c == b*2 && b+d == c*2 {
		Fprint(out, d*2-c)
	} else if a*c == b*b && b*d == c*c && d*d%c == 0 {
		Fprint(out, d*d/c)
	} else {
		Fprint(out, 42)
	}
}

//func main() { CF328A(os.Stdin, os.Stdout) }
