package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func f85(a []int) int {
	res, f := int(-1e18), 0
	for _, v := range a {
		f = max(f, 0) + v
		res = max(res, f)
	}
	return res
}

func cf1285B(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		s := 0
		for i := range a {
			Fscan(in, &a[i])
			s += a[i]
		}
		if max(f85(a[1:]), f85(a[:n-1])) < s {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1285B(bufio.NewReader(os.Stdin), os.Stdout) }
