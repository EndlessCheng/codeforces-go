package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf922C(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	for i := 1; i <= k; i++ {
		if (n+1)%i > 0 {
			Fprint(out, "No")
			return
		}
	}
	Fprint(out, "Yes")
}

//func main() { cf922C(os.Stdin, os.Stdout) }
