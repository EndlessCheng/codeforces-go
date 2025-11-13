package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1982C(in io.Reader, out io.Writer) {
	var T, n, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &l, &r)
		a := make([]int, n)
		ans, s, left := 0, 0, 0
		for i := range a {
			Fscan(in, &a[i])
			s += a[i]
			for s > r {
				s -= a[left]
				left++
			}
			if s >= l {
				ans++
				s = 0
				left = i + 1
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1982C(bufio.NewReader(os.Stdin), os.Stdout) }
