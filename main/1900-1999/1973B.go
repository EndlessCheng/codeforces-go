package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1973B(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 1
		for i := 0; i < 20; i++ {
			pre := -1
			for j, v := range a {
				if v>>i&1 > 0 {
					ans = max(ans, j-pre)
					pre = j
				}
			}
			if pre >= 0 {
				ans = max(ans, n-pre)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1973B(bufio.NewReader(os.Stdin), os.Stdout) }
