package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF166E(in io.Reader, out io.Writer) {
	const p int64 = 1e9 + 7
	var n int
	Fscan(in, &n)
	var d, abc int64 = 1, 0
	for i := 0; i < n; i++ {
		d, abc = abc*3%p, (d+abc*2)%p
	}
	Fprint(out, d)
}

//func main() { CF166E(os.Stdin, os.Stdout) }
