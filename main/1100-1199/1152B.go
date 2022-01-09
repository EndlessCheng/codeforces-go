package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1152B(in io.Reader, out io.Writer) {
	var n uint
	Fscan(in, &n)
	op := []int{}
	for {
		op = append(op, bits.TrailingZeros(n))
		n |= n&-n - 1
		if n&(n+1) == 0 {
			break
		}
		n++
	}
	Fprintln(out, len(op)*2-1)
	for _, v := range op {
		Fprint(out, v, " ")
	}
}

//func main() { CF1152B(os.Stdin, os.Stdout) }
