package main

import (
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF848A(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	if n == 0 {
		Fprint(out, "a")
		return
	}
	for c := 'a'; n > 0; c++ {
		x := 2
		for ; x*(x+1)/2 <= n; x++ {
		}
		n -= x * (x - 1) / 2
		Fprint(out, strings.Repeat(string(c), x))
	}
}

//func main() { CF848A(os.Stdin, os.Stdout) }
