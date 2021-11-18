package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF525B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var s []byte
	var m, v int
	Fscan(in, &s, &m)
	n := len(s)
	c := make([]int8, n)
	for ; m > 0; m-- {
		Fscan(in, &v)
		c[v-1] ^= 1
	}
	r := int8(0)
	for i, c := range c[:n/2] {
		if r ^= c; r > 0 {
			s[i], s[n-1-i] = s[n-1-i], s[i]
		}
	}
	Fprintf(out, "%s", s)
}

//func main() { CF525B(os.Stdin, os.Stdout) }
