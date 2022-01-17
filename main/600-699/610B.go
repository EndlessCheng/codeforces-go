package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF610B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, c, mxC int
	Fscan(in, &n)
	a := make([]int, n)
	min := int(2e9)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] < min {
			min = a[i]
		}
	}
	for i := 0; i < n*2; i++ {
		if a[i%n] == min {
			c = 0
		} else if c++; c > mxC {
			mxC = c
		}
	}
	ans := int64(min) * int64(n)
	if mxC > 0 {
		ans += int64(mxC)
	}
	Fprint(out, ans)
}

//func main() { CF610B(os.Stdin, os.Stdout) }
