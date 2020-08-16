package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1225C(in io.Reader, out io.Writer) {
	var n, p int
	Fscan(in, &n, &p)
	for i := 1; i*p < n; i++ {
		if x := n - i*p; i <= x && i >= bits.OnesCount(uint(x)) {
			Fprint(out, i)
			return
		}
	}
	Fprint(out, -1)
}

//func main() { CF1225C(os.Stdin, os.Stdout) }
