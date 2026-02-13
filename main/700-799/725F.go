package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf725F(in io.Reader, out io.Writer) {
	var n, a1, b1, a2, b2 int
	Fscan(in, &n)
	a := []int{}
	ans := 0
	for range n {
		Fscan(in, &a1, &b1, &a2, &b2)
		if a1+b1 >= a2+b2 {
			a = append(a, a1+b1, a2+b2)
			ans += a1 + a2
		} else if a1 > b2 {
			ans += a1 - b2
		} else if a2 < b1 {
			ans += a2 - b1
		}
	}
	slices.Sort(a)
	for i := 0; i < len(a); i += 2 {
		ans -= a[i]
	}
	Fprint(out, ans)
}

//func main() { cf725F(bufio.NewReader(os.Stdin), os.Stdout) }
