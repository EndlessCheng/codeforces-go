package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF912B(in io.Reader, out io.Writer) {
	var n, k uint64
	Fscan(in, &n, &k)
	if k == 1 {
		Fprint(out, n)
	} else {
		Fprint(out, int64(1)<<bits.Len64(n)-1)
	}
}

//func main() { CF912B(os.Stdin, os.Stdout) }
