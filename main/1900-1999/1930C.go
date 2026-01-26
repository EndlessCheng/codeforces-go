package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1930C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			a[i] += i + 1
		}

		slices.Sort(a)
		cur := int(2e9)
		for i := n - 1; i >= 0; i-- {
			cur = min(cur-1, a[i])
			Fprint(out, cur, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1930C(bufio.NewReader(os.Stdin), os.Stdout) }
