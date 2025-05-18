package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1270B(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &v)
		mn, mx := v, v
		minL, maxL, l, r := 0, 0, -1, -1
		for i := 1; i < n; i++ {
			Fscan(in, &v)
			if v-i > mn {
				l, r = minL, i
			} else {
				mn, minL = v-i, i
			}
			if v+i < mx {
				l, r = maxL, i
			} else {
				mx, maxL = v+i, i
			}
		}
		if r < 0 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
			Fprintln(out, l+1, r+1)
		}
	}
}

//func main() { cf1270B(bufio.NewReader(os.Stdin), os.Stdout) }
