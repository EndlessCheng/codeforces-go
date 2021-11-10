package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1037A(in io.Reader, out io.Writer) {
	var n uint
	Fscan(in, &n)
	Fprint(out, bits.Len(n))
}

//func main() { CF1037A(os.Stdin, os.Stdout) }
