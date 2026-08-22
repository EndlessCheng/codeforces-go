package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2244D(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var T, n, m, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]bool, n)
		for range m {
			Fscan(in, &v)
			b[v-1] = true
		}

		ans, s := 0, 0
		first := true
		for i := n - 1; i >= 0; i-- {
			if b[i] {
				if first {
					ans += s
					first = false
				} else {
					ans += abs(s)
				}
				s = 0
			}
			s += a[i]
		}
		ans += abs(s)
		Fprintln(out, ans)
	}
}

//func main() { cf2244D(bufio.NewReader(os.Stdin), os.Stdout) }
