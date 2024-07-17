package main

import (
	. "fmt"
	"io"
)

func cf1372B(in io.Reader, out io.Writer) {
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		for d := 2; d*d <= n; d++ {
			if n%d == 0 {
				Fprintln(out, n/d, n-n/d)
				continue o
			}
		}
		Fprintln(out, 1, n-1)
	}
}

//func main() { cf1372B(os.Stdin, os.Stdout) }
