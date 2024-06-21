package main

import (
	. "fmt"
	"io"
	"slices"
)

func cf1759E(in io.Reader, out io.Writer) {
	var T, n, h int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &h)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)
		f := func(m []int) int {
			sum := h
			for i, v := range a {
				for sum <= v {
					if len(m) == 0 {
						return i
					}
					sum *= m[0]
					m = m[1:]
				}
				sum += v / 2
			}
			return n
		}
		Fprintln(out, max(f([]int{2, 2, 3}), f([]int{2, 3, 2}), f([]int{3, 2, 2})))
	}
}

//func main() { cf1759E(bufio.NewReader(os.Stdin), os.Stdout) }
