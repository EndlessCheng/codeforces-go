package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF977A(in io.Reader, out io.Writer) {
	var n, k int
	for Fscan(in, &n, &k); k > 0; k-- {
		if n%10 > 0 {
			n--
		} else {
			n /= 10
		}
	}
	Fprint(out, n)
}

//func main() { CF977A(os.Stdin, os.Stdout) }
