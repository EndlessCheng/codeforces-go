package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2112C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ans := 0
		for i := 2; i < n; i++ {
			tar := max(a[i], a[n-1]-a[i])
			l, r := 0, i-1
			for l < r {
				if a[l]+a[r] > tar {
					ans += r - l
					r--
				} else {
					l++
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2112C(bufio.NewReader(os.Stdin), os.Stdout) }
