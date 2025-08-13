package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1954B(in io.Reader, out io.Writer) {
	var T, n, v0, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &v0)
		ans := n
		pre := -1
		for i := 1; i < n; i++ {
			Fscan(in, &v)
			if v != v0 {
				ans = min(ans, i-pre-1)
				pre = i
			}
		}
		ans = min(ans, n-pre-1)
		if ans == n {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1954B(bufio.NewReader(os.Stdin), os.Stdout) }
