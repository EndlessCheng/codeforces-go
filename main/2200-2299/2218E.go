package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2218E(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		ans := 0
		for i := range a {
			Fscan(in, &a[i])
			for _, w := range a[:i] {
				ans = max(ans, a[i]^w)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2218E(bufio.NewReader(os.Stdin), os.Stdout) }
