package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	f := func() bool {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		mn, mx := slices.Min(a), slices.Max(a)
		if mn == -mx {
			cn, cp := 0, 0
			for _, v := range a {
				if v == mn {
					cn++
				} else if v == mx {
					cp++
				}
			}
			if cn+cp == n {
				return abs(cn-cp) <= 1
			}
		}
		slices.SortFunc(a, func(a, b int) int { return abs(a) - abs(b) })
		for i := 1; i < n-1; i++ {
			if a[i]*a[i] != a[i-1]*a[i+1] {
				return false
			}
		}
		return true
	}
	for Fscan(in, &T); T > 0; T-- {
		if f() {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
func abs(x int) int { if x < 0 { return -x }; return x }
