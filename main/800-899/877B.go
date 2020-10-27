package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF877B(in io.Reader, out io.Writer) {
	var s []byte
	Fscan(bufio.NewReader(in), &s)
	n := len(s)
	sum := make([]int, n+1)
	for i, b := range s {
		sum[i+1] = sum[i]
		if b == 'a' {
			sum[i+1]++
		}
	}
	ans := sum[n]
	for i, b := range s {
		if b != 'b' {
			continue
		}
		for j := i; j < n; j++ {
			if s[j] != 'b' {
				continue
			}
			if s := sum[n] + j + 1 - i - 2*(sum[j+1]-sum[i]); s > ans {
				ans = s
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF877B(os.Stdin, os.Stdout) }
