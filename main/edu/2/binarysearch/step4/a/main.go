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
	const eps = 1e-8 // 由于任何 ±1 带来的均值变动至少是 1/n，eps 取 1e-8 绰绰有余
	in := bufio.NewReader(_r)

	var n, d, pl, pr int
	Fscan(in, &n, &d)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	l, r := 0.0, 1e2
o:
	for t := int(math.Log2((r - l) / eps)); t > 0; t-- {
		x := (l + r) / 2
		s := make([]float64, n+1)
		for i, v := range a {
			s[i+1] = s[i] + float64(v) - x
		}
		minI := 0
		for i, v := range s[:n+1-d] {
			if v < s[minI] {
				minI = i
			}
			if s[i+d] >= s[minI] {
				pl, pr = minI+1, i+d
				l = x
				continue o
			}
		}
		r = x
	}
	Fprint(out, pl, pr)
}

func main() { run(os.Stdin, os.Stdout) }
