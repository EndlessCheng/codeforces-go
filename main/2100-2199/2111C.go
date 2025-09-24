package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2111C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := int(1e18)
		for i := 0; i < n; {
			st := i
			v := a[st]
			for ; i < n && a[i] == v; i++ {
			}
			ans = min(ans, (n-i+st)*v)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2111C(bufio.NewReader(os.Stdin), os.Stdout) }
