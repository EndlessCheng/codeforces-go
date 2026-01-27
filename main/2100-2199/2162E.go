package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2162E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		has := make([]bool, n+2)
		for i := range a {
			Fscan(in, &a[i])
			has[a[i]] = true
		}

		x := 1
		for has[x] {
			x++
		}

		if x <= n {
			z := a[n-1]
			y := 1
			for y == x || y == z {
				y++
			}
			a[0] = x
			a[1] = y
			a[2] = z
		}

		for i := range k {
			Fprint(out, a[i%3], " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2162E(bufio.NewReader(os.Stdin), os.Stdout) }
