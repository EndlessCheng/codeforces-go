package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	var v, p int64
	Fscan(in, &n, &p)
	sum := make([]int64, 2*n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		sum[i+1] = sum[i] + v
	}
	for i := n; i < 2*n; i++ {
		sum[i+1] = sum[i] + sum[i+1-n] - sum[i-n]
	}
	dup, minL := p/sum[n]*int64(n), n
	p %= sum[n]
	for l := 0; l < n; l++ {
		// 2*n-l 也可以写成 n
		if d := sort.Search(2*n-l, func(d int) bool { return sum[l+d]-sum[l] >= p }); d < minL {
			k, minL = l, d
		}
	}
	Fprint(out, k+1, dup+int64(minL))
}

func main() { run(os.Stdin, os.Stdout) }
