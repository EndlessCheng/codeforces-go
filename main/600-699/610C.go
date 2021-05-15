package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF610C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var k int
	Fscan(in, &k)
	n := 1 << k
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if bits.OnesCount(uint(i&j))&1 > 0 {
				Fprint(out, "*")
			} else {
				Fprint(out, "+")
			}
		}
		Fprintln(out)
	}
}

//func main() { CF610C(os.Stdin, os.Stdout) }
