package main

import (
	. "fmt"
	"io"
)

func cf1765K(in io.Reader, out io.Writer) {
	var n, v, s int
	Fscan(in, &n)
	mn := int(1e9)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			Fscan(in, &v)
			s += v
			if j == n-1-i {
				mn = min(mn, v)
			}
		}
	}
	Fprint(out, s-mn)
}

//func main() { cf1765K(bufio.NewReader(os.Stdin), os.Stdout) }
