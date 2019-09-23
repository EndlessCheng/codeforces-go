package cf

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
	"sort"
)

// https://codeforces.com/contest/257/problem/C

var (
	// for test
	reader io.Reader = os.Stdin
	writer io.Writer = os.Stdout
)

func minF(a, b float64) float64 {
	if a <= b {
		return a
	}
	return b
}

func Ans257C() {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, x, y int
	Fscan(in, &n)
	if n == 1 {
		Fprintln(out, 0)
		return
	}

	degs := make([]float64, n*2)
	for i := 0; i < n; i++ {
		Fscan(in, &x, &y)
		deg := math.Atan2(float64(y), float64(x)) * 180.0 / math.Pi
		if deg < 0 {
			deg = 360.0 + deg
		}
		degs[i] = deg
		degs[i+n] = 360.0 + deg
	}
	sort.Float64s(degs)

	minAngle := 360.0
	for i := 0; i < n; i++ {
		angle := degs[n-1+i] - degs[i]
		minAngle = minF(minAngle, angle)
	}
	Fprintln(out, minAngle)
}
