package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1362B(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		has := [1024]bool{}
		for i := range a {
			Fscan(in, &a[i])
			has[a[i]] = true
		}
		if n%2 > 0 {
			Fprintln(out, -1)
			continue
		}
		ans := int(1e9)
	o:
		for i := 1; i < n; i++ {
			xor := a[0] ^ a[i]
			for _, v := range a {
				if !has[v^xor] {
					continue o
				}
			}
			ans = min(ans, xor)
		}
		if ans == 1e9 {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1362B(bufio.NewReader(os.Stdin), os.Stdout) }
