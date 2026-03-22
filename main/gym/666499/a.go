package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cfA(in io.Reader, out io.Writer) {
	var T, n, c, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &c, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)
		for _, v := range a {
			d := min(c-v, k)
			if d < 0 {
				break
			}
			v += d
			k -= d
			c += v
		}
		Fprintln(out, c)
	}
}

//func main() { cfA(bufio.NewReader(os.Stdin), os.Stdout) }
