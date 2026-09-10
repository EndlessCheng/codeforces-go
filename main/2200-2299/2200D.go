package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2200D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x, &y)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		b := a[x:y]
		i := slices.Index(b, slices.Min(b))
		b = append(slices.Clone(b[i:]), b[:i]...)

		a = append(a[:x], a[y:]...)
		i = 0
		for i < len(a) && a[i] < b[0] {
			i++
		}
		a = slices.Insert(a, i, b...)

		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2200D(bufio.NewReader(os.Stdin), os.Stdout) }
