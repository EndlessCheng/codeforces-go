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
		a := make([]int, n+5)
		for i := 3; i <= n+2; i++ {
			Fscan(in, &a[i])
		}
		a[2] = 1
		a[n+4] = a[n+1]
		end := n + 1
		for i := n; i > 3 && gcd(a[i], a[i+1]) <= gcd(a[i+1], a[i+2]); i-- {
			end = i
		}
		// 枚举删除 a[i]
		for i := 3; i <= n+2 && gcd(a[i-3], a[i-2]) <= gcd(a[i-2], a[i-1]); i++ {
			g := gcd(a[i-1], a[i+1])
			if i+1 >= end && gcd(a[i-2], a[i-1]) <= g && g <= gcd(a[i+1], a[i+2]) {
				Fprintln(out, "YES")
				continue o
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { cf1980D(bufio.NewReader(os.Stdin), os.Stdout) }
