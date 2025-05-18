package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1270B(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var T, n, pre, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &pre)
		l, r := 0, 0
		for i := 1; i < n; i++ {
			Fscan(in, &v)
			if abs(v-pre) > 1 {
				l, r = i, i+1
			}
			pre = v
		}
		if r == 0 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
			Fprintln(out, l, r)
		}
	}
}

//func main() { cf1270B(bufio.NewReader(os.Stdin), os.Stdout) }
