package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2207E2(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}

		ans := 1
		for i := 1; i <= n; i++ {
			if a[i] < n-i || a[i] > n || i > 1 && a[i] > a[i-1] {
				ans = 0
			}
			if i > 1 && a[i] == a[i-1] {
				ans = ans * (a[i] - n + i) % mod
			} else {
				ans = ans * i % mod
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2207E2(bufio.NewReader(os.Stdin), os.Stdout) }
