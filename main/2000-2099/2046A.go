package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2046A(in io.Reader, out io.Writer) {
	var T, n, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 0
		mx := int(-1e18)
		for _, x := range a {
			Fscan(in, &y)
			ans += max(x, y)
			mx = max(mx, min(x, y))
		}
		Fprintln(out, ans+mx)
	}
}

//func main() { cf2046A(bufio.NewReader(os.Stdin), os.Stdout) }
