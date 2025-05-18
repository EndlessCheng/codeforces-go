package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1166C(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] < 0 {
			a[i] = -a[i]
		}
	}
	slices.Sort(a)

	l := 0
	for i, v := range a {
		for a[l]*2 < v {
			l++
		}
		ans += i - l
	}
	Fprint(out, ans)
}

//func main() { cf1166C(bufio.NewReader(os.Stdin), os.Stdout) }
