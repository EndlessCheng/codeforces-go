package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF735D(_r io.Reader, _w io.Writer) {
	isP := func(n int) bool {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	var n int
	if Fscan(_r, &n); isP(n) {
		Fprint(_w, 1)
	} else if n%2 == 0 || isP(n-2) {
		Fprint(_w, 2)
	} else {
		Fprint(_w, 3)
	}
}

//func main() { CF735D(os.Stdin, os.Stdout) }
