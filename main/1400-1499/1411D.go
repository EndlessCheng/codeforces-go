package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1411D(in io.Reader, out io.Writer) {
	var s string
	var w01, w10, sum int64
	Fscan(bufio.NewReader(in), &s, &w01, &w10)
	tar := '1'
	if w01 < w10 {
		w01, w10, tar = w10, w01, '0'
	}
	c1 := 0
	// '?' as '0'
	for i, b := range s {
		if b == tar { // '1'
			sum += int64(i-c1) * w01
			c1++
		} else { // '0' or '?'
			sum += int64(c1) * w10
		}
	}
	ans := sum
	for i, b := range s {
		if b == '?' { // '1'
			sum += int64(i-c1)*w01 + int64(len(s)-1-c1-i)*w10
			if sum < ans {
				ans = sum
			}
			c1++
		}
	}
	Fprint(out, ans)
}

//func main() { CF1411D(os.Stdin, os.Stdout) }
