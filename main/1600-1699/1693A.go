package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1693A(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s := 0
		ok := true
		vis0 := false
		for range n {
			Fscan(in, &v)
			s += v
			if s < 0 || vis0 && s > 0 {
				ok = false
			}
			if s == 0 {
				vis0 = true
			}
		}
		if ok && s == 0 {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

//func main() { cf1693A(bufio.NewReader(os.Stdin), os.Stdout) }
