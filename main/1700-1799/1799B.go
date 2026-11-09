package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1799B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		idx := -1
		for i := range a {
			Fscan(in, &a[i])
			if a[i] != a[0] {
				idx = i
			}
		}
		if idx < 0 {
			Fprintln(out, 0)
			continue
		}
		if slices.Min(a) == 1 {
			Fprintln(out, -1)
			continue
		}

		ans := [][2]int{}
		for {
			minI := slices.Index(a, slices.Min(a))
			maxI := slices.Index(a, slices.Max(a))
			if a[minI] == a[maxI] {
				break
			}
			ans = append(ans, [2]int{maxI, minI})
			a[maxI] = (a[maxI]-1)/a[minI] + 1
		}

		Fprintln(out, len(ans))
		for _, p := range ans {
			Fprintln(out, p[0]+1, p[1]+1)
		}
	}
}

//func main() { cf1799B(bufio.NewReader(os.Stdin), os.Stdout) }
