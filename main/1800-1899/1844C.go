package main

import (
	. "fmt"
	"io"
)

func cf1844C(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		mx := int(-1e9)
		s := [2]int{}
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			mx = max(mx, v)
			s[i%2] += max(v, 0)
		}
		if mx < 0 {
			Fprintln(out, mx)
		} else {
			Fprintln(out, max(s[0], s[1]))
		}
	}
}

//func main() { cf1844C(bufio.NewReader(os.Stdin), os.Stdout) }
