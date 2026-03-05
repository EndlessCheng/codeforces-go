package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1768F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([]int, n+1)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		f[i] = 1e18
	}
	f[1] = 0
	for i := 1; i <= n; i++ {
		m := n / a[i]
		for j := i - 1; j >= 1 && j >= i-m; j-- {
			f[i] = min(f[i], f[j]+a[i]*(i-j)*(i-j))
			if a[j] <= a[i] {
				break
			}
		}
		for j := i + 1; j <= n && j <= i+m; j++ {
			f[j] = min(f[j], f[i]+a[i]*(j-i)*(j-i))
			if a[j] <= a[i] {
				break
			}
		}
		Fprint(out, f[i], " ")
	}
}

//func main() { cf1768F(bufio.NewReader(os.Stdin), os.Stdout) }
