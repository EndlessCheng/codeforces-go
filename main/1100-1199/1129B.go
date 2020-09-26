package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1129B(in io.Reader, out io.Writer) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	Fprint(out, "2000\n-1")
	var k int
	Fscan(in, &k)
	k += 2000
	for i := 1; i < 2000; i++ {
		v := min(1e6, k)
		Fprint(out, " ", v)
		k -= v
	}
}

//func main() { CF1129B(os.Stdin, os.Stdout) }
