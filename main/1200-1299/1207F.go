package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf1207F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 500001
	a := [mx]int32{}
	var q, op, x uint32
	var y int32
	Fscan(in, &q)
	B := uint32(math.Sqrt(mx / 5))
	sum := make([][]int32, B)
	for i := range sum {
		sum[i] = make([]int32, i)
	}
	for range q {
		Fscan(in, &op, &x, &y)
		if op == 1 {
			a[x] += y
			for i := uint32(1); i < B; i++ {
				sum[i][x%i] += y
			}
		} else if x < B {
			Fprintln(out, sum[x][y])
		} else {
			s := int32(0)
			for i := uint32(y); i < mx; i += x {
				s += a[i]
			}
			Fprintln(out, s)
		}
	}
}

//func main() { cf1207F(bufio.NewReader(os.Stdin), os.Stdout) }
