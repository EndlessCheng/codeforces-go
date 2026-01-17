package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf989D(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var n, l, maxW, x, v, ans int
	Fscan(in, &n, &l, &maxW)
	var a, b []int
	for range n {
		Fscan(in, &x, &v)
		if v > 0 {
			a = append(a, x)
		} else {
			b = append(b, x)
		}
	}
	slices.Sort(a)
	slices.Sort(b)

	j := 0
	for _, x := range a {
		for j < len(b) && (b[j]-x+l)*maxW <= abs(x+b[j]+l) {
			j++
		}
		ans += len(b) - j
	}
	Fprint(out, ans)
}

//func main() { cf989D(bufio.NewReader(os.Stdin), os.Stdout) }
