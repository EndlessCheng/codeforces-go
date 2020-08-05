package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const eps = 1e-8
	var n, k int
	Fscan(in, &n, &k)
	a := make([]float64, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	l, r := 0.0, 1e7
	for t := int(math.Log2((r - l) / eps)); t > 0; t-- {
		sz := (l + r) / 2
		c := 0
		for _, v := range a {
			c += int(v / sz)
		}
		if c < k {
			r = sz
		} else {
			l = sz
		}
	}
	Fprintf(out, "%.8f", (l+r)/2)
}

func main() { run(os.Stdin, os.Stdout) }
