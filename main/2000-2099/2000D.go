package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2000D(in io.Reader, out io.Writer) {
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		sum := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &sum[i])
			sum[i] += sum[i-1]
		}
		Fscan(in, &s)

		ans := 0
		l, r := 0, n-1
		for l < r {
			if s[l] != 'L' {
				l++
			} else if s[r] != 'R' {
				r--
			} else {
				ans += sum[r+1] - sum[l]
				l++
				r--
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2000D(bufio.NewReader(os.Stdin), os.Stdout) }
