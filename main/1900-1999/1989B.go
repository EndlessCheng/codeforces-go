package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1989B(in io.Reader, out io.Writer) {
	f := func(s, t []byte) int {
		cnt := 0
		for _, b := range s {
			if t[cnt] != b {
				continue
			}
			cnt++
			if cnt == len(t) {
				break
			}
		}
		return cnt
	}

	var T int
	var s, t []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &t)
		mx := 0
		for i := range t {
			mx = max(mx, f(s, t[i:]))
		}
		Fprintln(out, len(s)+len(t)-mx)
	}
}

//func main() { cf1989B(bufio.NewReader(os.Stdin), os.Stdout) }
