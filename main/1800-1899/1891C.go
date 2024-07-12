package main

import (
	. "fmt"
	"io"
	"slices"
)

func cf1891C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		s := 0
		for i := range a {
			Fscan(in, &a[i])
			s += a[i]
		}
		slices.Sort(a)

		left := (s + 1) / 2
		ans := left
		for i, v := range a {
			if left < v {
				ans += n - i
				break
			}
			left -= v
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1891C(bufio.NewReader(os.Stdin), os.Stdout) }
