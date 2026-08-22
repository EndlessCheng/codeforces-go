package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2244D(in io.Reader, out io.Writer) {
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
		for i, v := range a {
			s += v
			if b[i] {
				ans += max(s, -s)
				s = 0
			}
		}
		Fprintln(out, ans+s)
	}
}

//func main() { cf2244D(bufio.NewReader(os.Stdin), os.Stdout) }
