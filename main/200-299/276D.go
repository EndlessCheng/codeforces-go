package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF276D(in io.Reader, out io.Writer) {
	var l, r uint64
	Fscan(in, &l, &r)
	Fprint(out, int64(1)<<bits.Len64(r^l)-1)
}

//func main() { CF276D(os.Stdin, os.Stdout) }
