package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1980D(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		end := n - 2
		for i := n - 3; i > 0 && gcd(a[i], a[i+1]) <= gcd(a[i+1], a[i+2]); i-- {
			end = i
		}
		if end < 2 {
			Fprintln(out, "YES")
			continue
		}
		for i := 1; i < n && (i < 3 || gcd(a[i-3], a[i-2]) <= gcd(a[i-2], a[i-1])); i++ {
			if i+1 >= end && (i == 1 || i == n-1 || gcd(a[i-2], a[i-1]) <= gcd(a[i-1], a[i+1])) &&
				(i >= n-2 || gcd(a[i-1], a[i+1]) <= gcd(a[i+1], a[i+2])) {
				Fprintln(out, "YES")
				continue o
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { cf1980D(bufio.NewReader(os.Stdin), os.Stdout) }
