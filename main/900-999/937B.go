package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf937B(in io.Reader, out io.Writer) {
	var p, y int
	Fscan(in, &p, &y)
o:
	for x := y; x > p; x-- {
		for i := 2; i <= p && i*i <= x; i++ {
			if x%i == 0 {
				continue o
			}
		}
		Fprint(out, x)
		return
	}
	Fprint(out, -1)
}

//func main() { cf937B(os.Stdin, os.Stdout) }
