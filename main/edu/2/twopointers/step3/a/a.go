package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	var p int64
	Fscan(in, &n, &p)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	a = append(a, a...)
	sum := make([]int64, 2*n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	mi := int64(2e18)
	for l := 0; l < n; l++ {
		for r := l + 1; r <= l+n; r++ {
			c := int64(r - l)
			if sum[r]-sum[l] < p {
				c += ((p-sum[r]+sum[l]-1)/sum[n] + 1) * int64(n)
			}
			if c < mi {
				k, mi = l, c
			}
		}
	}
	Fprint(out, k+1, mi)
}

func main() { run(os.Stdin, os.Stdout) }
