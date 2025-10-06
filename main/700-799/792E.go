package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://gemini.google.com/app/4c413db204a68a63

// https://github.com/EndlessCheng
func cf792E(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	f := func(k int) bool {
		ans := 0
		for _, v := range a {
			if v%k == 0 {
				ans += v / k
			} else if v/k+v%k >= k-1 {
				ans += v/k + 1
			} else {
				return false
			}
		}
		Fprint(out, ans)
		return true
	}
	mn := slices.Min(a)
	for i := 1; !f(mn/i+1) && !f(mn/i); i++ {
	}
}

//func main() { cf792E(os.Stdin, os.Stdout) }
