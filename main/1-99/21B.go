package main

import (
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func gaussJordanElimination(a [][]float64) (sol []float64, infSol bool) {
	const eps = 1e-8
	n := len(a)
	row := 0
	for col := 0; col < n; col++ {
		pivot := row
		for i := row; i < n; i++ {
			if math.Abs(a[i][col]) > math.Abs(a[pivot][col]) {
				pivot = i
			}
		}
		if math.Abs(a[pivot][col]) < eps {
			continue
		}
		a[row], a[pivot] = a[pivot], a[row]
		for j := col + 1; j <= n; j++ {
			a[row][j] /= a[row][col]
		}
		for i := range a {
			if i != row {
				for j := col + 1; j <= n; j++ {
					a[i][j] -= a[i][col] * a[row][j]
				}
			}
		}
		row++
	}
	if row < n {
		for _, r := range a[row:] {
			if math.Abs(r[n]) > eps {
				return nil, false
			}
		}
		return nil, true
	}
	res := make([]float64, n)
	for i, r := range a {
		res[i] = r[n]
	}
	return res, false
}

func CF21B(in io.Reader, out io.Writer) {
	a := make([][]float64, 2)
	for i := range a {
		a[i] = make([]float64, 3)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	sol, inf := gaussJordanElimination(a)
	if inf {
		Fprint(out, -1)
	} else if sol == nil {
		Fprint(out, 0)
	} else {
		Fprint(out, 1)
	}
}

//func main() { CF21B(os.Stdin, os.Stdout) }
