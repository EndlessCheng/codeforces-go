package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf938F(in io.Reader, out io.Writer) {
	var s []byte
	Fscan(in, &s)
	n := len(s)
	m := bits.Len(uint(n)) - 1
	f := make([]bool, 1<<m)
	f[0] = true
	for i := range n - 1<<m + 1 {
		for j, ok := range f {
			if !ok {
				continue
			}
			for k := range m {
				f[j|1<<k] = true
			}
		}

		mnC := byte(127)
		for j, ok := range f {
			if ok {
				mnC = min(mnC, s[i+j])
			}
		}

		for j, b := range s[i : i+1<<m] {
			if b != mnC {
				f[j] = false
			}
		}

		s[i] = mnC
	}
	Fprintf(out, "%s", s[:n-1<<m+1])
}

//func main() { cf938F(os.Stdin, os.Stdout) }
