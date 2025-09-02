package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1279B(in io.Reader, out io.Writer) {
	var T, n, s, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		var ans, mx, mxI int
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			if v > mx {
				mx, mxI = v, i
			}
			s -= v
			if s >= -mx {
				ans = mxI
			}
		}
		if s >= 0 {
			Fprintln(out, 0)
		} else {
			Fprintln(out, ans)
		}
	}
}

//func main() { cf1279B(bufio.NewReader(os.Stdin), os.Stdout) }
