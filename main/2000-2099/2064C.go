package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2064C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		negS := 0
		for i := range a {
			Fscan(in, &a[i])
			if a[i] < 0 {
				negS -= a[i]
			}
		}

		ans := negS
		posS := 0
		for _, v := range a {
			if v < 0 {
				negS += v
			} else {
				posS += v
			}
			ans = max(ans, posS+negS)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2064C(bufio.NewReader(os.Stdin), os.Stdout) }
