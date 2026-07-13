package main

import (
	"bytes"
	"math/bits"
)

func createGrid(k int) []string {
	w := bits.Len(uint(k))
	m := w * 2
	n := w + 3

	a := make([][]byte, m)
	for i := range a {
		a[i] = bytes.Repeat([]byte{'#'}, n)
		a[i][n-1] = '.'
	}

	for j := range w {
		i := j * 2
		a[i][j] = '.'
		a[i][j+1] = '.'
		a[i+1][j] = '.'
		a[i+1][j+1] = '.'
	}

	for i := range w {
		if k>>i&1 > 0 {
			for j := i + 2; j < n-1; j++ {
				a[i*2][j] = '.'
			}
		}
	}

	ans := make([]string, m)
	for i, row := range a {
		ans[i] = string(row)
	}
	return ans
}
