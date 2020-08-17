package main

import (
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1368B(in io.Reader, out io.Writer) {
	var k int64
	Fscan(in, &k)
	n := 10
	for {
		s := int64(1)
		for i := 0; i < 10; i++ {
			s *= int64((n + i) / 10)
		}
		if s >= k {
			break
		}
		n++
	}
	for i, b := range "codeforces" {
		Fprint(out, strings.Repeat(string(b), (n+9-i)/10))
	}
}

//func main() { CF1368B(os.Stdin, os.Stdout) }
