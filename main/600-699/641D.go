package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf641D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)

	p1 := make([]float64, n)
	for i := range p1 {
		Fscan(in, &p1[i])
	}
	for i := 1; i < n; i++ {
		p1[i] += p1[i-1]
	}

	p2 := make([]float64, n)
	for i := range p2 {
		Fscan(in, &p2[i])
	}
	for i := n - 2; i >= 0; i-- {
		p2[i] += p2[i+1]
	}

	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n-1; i++ {
		b := p1[i] + 1 - p2[i+1]
		d := math.Sqrt(b*b - p1[i]*4)
		x[i] = (b + d) / 2
		y[i] = (b - d) / 2
	}

	x[n-1] = 1
	y[n-1] = 1
	for i := n - 1; i >= 1; i-- {
		x[i] -= x[i-1]
		y[i] -= y[i-1]
	}

	for _, v := range x {
		Fprintf(out, "%.6f ", v)
	}
	Fprintln(out)
	for _, v := range y {
		Fprintf(out, "%.6f ", v)
	}
}

//func main() { cf641D(bufio.NewReader(os.Stdin), os.Stdout) }
