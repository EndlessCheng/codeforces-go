package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2118E(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := [][]int{}
		for y := range m {
			for x := range n {
				dx := abs(x - n/2)
				dy := abs(y - m/2)
				a = append(a, []int{max(dy, dx), dx, dy, x, y})
			}
		}
		slices.SortFunc(a, func(a, b []int) int { return slices.Compare(a, b) })
		for _, t := range a {
			Fprintln(out, t[3]+1, t[4]+1)
		}
	}
}

//func main() { cf2118E(os.Stdin, os.Stdout) }
