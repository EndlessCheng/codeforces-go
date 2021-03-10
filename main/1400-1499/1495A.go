package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1495A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var T, n, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		var a, b []int
		for n *= 2; n > 0; n-- {
			Fscan(in, &x, &y)
			if x < 0 {
				x = -x
			}
			if y < 0 {
				y = -y
			}
			if x == 0 {
				a = append(a, abs(y))
			} else {
				b = append(b, abs(x))
			}
		}
		sort.Ints(a)
		sort.Ints(b)
		s := .0
		for i, y := range a {
			s += math.Hypot(float64(b[i]), float64(y))
		}
		Fprintf(out, "%.15f\n", s)
	}
}

//func main() { CF1495A(os.Stdin, os.Stdout) }
