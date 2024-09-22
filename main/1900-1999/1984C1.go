package main

import (
	. "fmt"
	"io"
)

func cf1984C1(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s, minS := 0, 0
		for ; n > 0; n-- {
			Fscan(in, &v)
			s += v
			minS = min(minS, s)
		}
		Fprintln(out, s-minS*2)
	}
}

//func main() { cf1984C1(bufio.NewReader(os.Stdin), os.Stdout) }
