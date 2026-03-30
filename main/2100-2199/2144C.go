package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2144C(in io.Reader, out io.Writer) {
	const mod = 998244353
	var T, n, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 1
		preV, preW := 0, 0
		for _, v := range a {
			Fscan(in, &w)
			if v >= preV && w >= preW && v >= preW && w >= preV {
				ans = ans * 2 % mod
			}
			preV, preW = v, w
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2144C(bufio.NewReader(os.Stdin), os.Stdout) }
