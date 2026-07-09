package main

import (
	. "fmt"
	"io"
	"math"
	"slices"
)

// https://github.com/EndlessCheng
func cf1025F(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	x := make([]int, n)
	y := make([]int, n)
	for i := range x {
		Fscan(in, &x[i], &y[i])
	}

	a := make([]float64, n-1)
	ans := 0
	for i := range n {
		m := 0
		for j := range n {
			if i != j {
				a[m] = math.Atan2(float64(y[j]-y[i]), float64(x[j]-x[i]))
				m++
			}
		}
		slices.Sort(a)
		k := 0
		for j := 0; j < m && a[j] <= 0; j++ {
			for k < m && a[k]-a[j] < math.Pi {
				k++
			}
			ans += (k - j - 1) * (m - k + j) * (k - j - 2) * (m - k + j - 1) / 4
		}
	}
	Fprint(out, ans)
}

//func main() { cf1025F(bufio.NewReader(os.Stdin), os.Stdout) }
