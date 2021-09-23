package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1242A(in io.Reader, out io.Writer) {
	var n int64
	Fscan(in, &n)
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			if n > 1 {
				Fprint(out, 1)
			} else {
				Fprint(out, i)
			}
			return
		}
	}
	if n > 1 {
		Fprint(out, n)
	} else {
		Fprint(out, 1)
	}
}

//func main() { CF1242A(os.Stdin, os.Stdout) }
