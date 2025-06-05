package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1438E(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		v := a[i]
		mx := 1 << bits.Len(uint(v))
		s := 0
		for j := i - 2; j >= 0 && s+a[j+1] < mx; j-- {
			s += a[j+1]
			if a[j] < v && v^a[j] == s {
				ans++
			}
		}
	}
	for i, v := range a {
		mx := 1 << bits.Len(uint(v))
		s := 0
		for j := i + 2; j < n && s+a[j-1] < mx; j++ {
			s += a[j-1]
			if a[j] <= v && v^a[j] == s {
				ans++
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf1438E(bufio.NewReader(os.Stdin), os.Stdout) }
