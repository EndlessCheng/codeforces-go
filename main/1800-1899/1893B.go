package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

func cf1893B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, m)
		for i := range b {
			Fscan(in, &b[i])
		}
		slices.SortFunc(b, func(a, b int) int { return b - a })

		j := 0
		for _, v := range a {
			for ; j < m && b[j] >= v; j++ {
				Fprint(out, b[j], " ")
			}
			Fprint(out, v, " ")
		}
		for _, v := range b[j:] {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1893B(bufio.NewReader(os.Stdin), os.Stdout) }
