package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2065D(in io.Reader, out io.Writer) {
	var T, n, m, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		ans := 0
		a := make([]int, n)
		for i := range a {
			for j := m; j > 0; j-- {
				Fscan(in, &v)
				ans += v * j
				a[i] += v
			}
		}
		slices.Sort(a)
		for i, s := range a {
			ans += s * i * m
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2065D(bufio.NewReader(os.Stdin), os.Stdout) }
