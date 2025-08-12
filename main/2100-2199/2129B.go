package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2129B(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 0
		for i, v := range a {
			l, r := 0, 0
			for _, w := range a[:i] {
				if w > v {
					l++
				}
			}
			for _, w := range a[i+1:] {
				if w > v {
					r++
				}
			}
			ans += min(l, r)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2129B(bufio.NewReader(os.Stdin), os.Stdout) }
