package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1343C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 0
		for i := 0; i < n; {
			mx := a[i]
			for i++; i < n && a[i] > 0 == (mx > 0); i++ {
				mx = max(mx, a[i])
			}
			ans += mx
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1343C(bufio.NewReader(os.Stdin), os.Stdout) }
