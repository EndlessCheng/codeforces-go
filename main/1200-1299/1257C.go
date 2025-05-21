package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1257C(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		last := make([]int, n+1)
		ans := n + 1
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			if last[v] > 0 {
				ans = min(ans, i-last[v]+1)
			}
			last[v] = i
		}
		if ans > n {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1257C(bufio.NewReader(os.Stdin), os.Stdout) }
