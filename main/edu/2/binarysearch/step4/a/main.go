package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const eps = 1e-8 // 由于任何 ±1 带来的均值变动至少是 1/n，eps 取 1e-8 绰绰有余

	var n, d, pl, pr int
	Fscan(in, &n, &d)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	l, r := -1.0, 101.0
o:
	for t := int(math.Log2((r - l) / eps)); t > 0; t-- {
		mid := (l + r) / 2
		sum := make([]float64, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + float64(v) - mid
		}
		minI := 0
		for i, v := range sum[:n+1-d] {
			if v < sum[minI] {
				minI = i
			}
			if sum[i+d] >= sum[minI] {
				pl, pr = minI+1, i+d
				l = mid
				continue o
			}
		}
		r = mid
	}
	Fprint(out, pl, pr)
}

func main() { run(os.Stdin, os.Stdout) }
