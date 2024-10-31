package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2033C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		f0, f1 := 0, 0
		for i := 1; i < n/2; i++ {
			a, b, c, d := a[i-1], a[i], a[n-1-i], a[n-i]
			e0 := eq33(a, b) + eq33(c, d)
			e1 := eq33(a, c) + eq33(b, d)
			f0, f1 = min(f0+e0, f1+e1), min(f0+e1, f1+e0)
		}
		ans := min(f0, f1) + eq33(a[n/2-1], a[n/2])
		if n%2 > 0 {
			ans += eq33(a[n/2], a[n/2+1])
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2033C(bufio.NewReader(os.Stdin), os.Stdout) }
func eq33(v, w int) int { if v == w { return 1 }; return 0 }
