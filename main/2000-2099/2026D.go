package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf2026D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, v, q, l, r int
	Fscan(in, &n)
	sum := make([]int, n+1)
	iSum := make([]int, n+1)
	for i := range n {
		Fscan(in, &v)
		sum[i+1] = sum[i] + v
		iSum[i+1] = iSum[i] + (n-i)*v
	}

	iSumSum := make([]int, n+1)
	for i := range n {
		iSumSum[i+1] = iSumSum[i] + iSum[n] - iSum[i]
	}

	m := n*2 + 1
	f := func(k int) int {
		i := (m - int(math.Ceil(math.Sqrt(float64(m*m-k*8))))) / 2
		k -= (m - i) * i / 2
		return iSumSum[i] + iSum[i+k] - iSum[i] - (n-i-k)*(sum[i+k]-sum[i])
	}

	Fscan(in, &q)
	for range q {
		Fscan(in, &l, &r)
		Fprintln(out, f(r)-f(l-1))
	}
}

//func main() { cf2026D(bufio.NewReader(os.Stdin), os.Stdout) }
