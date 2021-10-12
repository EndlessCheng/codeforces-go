package main

import (
	. "fmt"
	"io"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func CF1177B(in io.Reader, out io.Writer) {
	var k int64
	Fscan(in, &k)
	for i, p10 := int64(1), int64(10); ; i++ {
		if i*p10 > k {
			Fprintf(out, "%c", strconv.FormatInt(k/i, 10)[k%i])
			return
		}
		k += p10
		p10 *= 10
	}
}

//func main() { CF1177B(os.Stdin, os.Stdout) }
