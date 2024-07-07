package main

import (
	. "fmt"
	"io"
)

func cf58B(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	for p := 2; p*p <= n; p++ {
		for ; n%p == 0; n /= p {
			Fprint(out, n, " ")
		}
	}
	if n > 1 {
		Fprint(out, n, " ")
	}
	Fprint(out, 1)
}

//func main() { cf58B(os.Stdin, os.Stdout) }
