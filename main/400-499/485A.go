package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func cf485A(in io.Reader, out io.Writer) {
	var n, m uint
	Fscan(in, &n, &m)
	if n%(m>>bits.TrailingZeros(m)) == 0 {
		Fprint(out, "Yes")
	} else {
		Fprint(out, "No")
	}
}

//func main() { cf485A(os.Stdin, os.Stdout) }
