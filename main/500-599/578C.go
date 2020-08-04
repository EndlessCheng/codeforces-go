package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func Sol578C(in io.Reader, out io.Writer) {
	maxSubArrayAbsSum := func(a []float64) float64 {
		min, max, abs := math.Min, math.Max, math.Abs
		curMaxSum, maxSum := a[0], a[0]
		curMinSum, minSum := a[0], a[0]
		for _, v := range a[1:] {
			curMaxSum = max(curMaxSum+v, v)
			maxSum = max(maxSum, curMaxSum)
			curMinSum = min(curMinSum+v, v)
			minSum = min(minSum, curMinSum)
		}
		return max(abs(maxSum), abs(minSum))
	}
	ternarySearch := func(l, r float64, f func(x float64) float64) float64 {
		step := int(math.Log((r-l)/1e-12) / math.Log(1.5))
		for i := 0; i < step; i++ {
			m1 := l + (r-l)/3
			m2 := r - (r-l)/3
			v1, v2 := f(m1), f(m2)
			if v1 < v2 {
				r = m2
			} else {
				l = m1
			}
		}
		return (l + r) / 2
	}

	in = bufio.NewReader(in)
	var n int
	Fscan(in, &n)
	a := make([]float64, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	f := func(x float64) float64 {
		b := make([]float64, n)
		for i, v := range a {
			b[i] = v - x
		}
		return maxSubArrayAbsSum(b)
	}
	x := ternarySearch(-1e4, 1e4, f)
	Fprintf(out, "%.15f", f(x))
}

//func main() {
//	Sol578C(os.Stdin, os.Stdout)
//}
