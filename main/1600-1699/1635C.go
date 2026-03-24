package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1635C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		if a[n-2] > a[n-1] || a[n-1] < 0 && !slices.IsSorted(a) {
			Fprintln(out, -1)
			continue
		}
		Fprintln(out, n-2)
		for i := n - 2; i > 0; i-- {
			Fprintln(out, i, i+1, n)
		}
	}
}

//func main() { cf1635C(bufio.NewReader(os.Stdin), os.Stdout) }
