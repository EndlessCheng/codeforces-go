package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1798D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)
		if a[0] == 0 {
			Fprintln(out, "No")
			continue
		}
		Fprintln(out, "Yes")
		s, i, j := 0, 0, n-1
		for i <= j {
			if s < 0 {
				s += a[j]
				Fprint(out, a[j], " ")
				j--
			} else {
				s += a[i]
				Fprint(out, a[i], " ")
				i++
			}
		}
		Fprintln(out)
	}
}

//func main() { cf1798D(bufio.NewReader(os.Stdin), os.Stdout) }
