package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2071B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	bad := []int{1, 8, 49, 288, 1681, 9800, 57121, 332928}
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if slices.Contains(bad, n) {
			Fprintln(out, -1)
			continue
		}
		a := make([]int, n)
		for i := range a {
			a[i] = i + 1
		}
		for _, v := range bad {
			if v > n {
				break
			}
			a[v-1], a[v] = a[v], a[v-1]
		}
		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2071B(bufio.NewReader(os.Stdin), os.Stdout) }
