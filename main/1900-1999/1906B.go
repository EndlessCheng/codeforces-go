package main

import (
	. "fmt"
	"io"
)

func cf1906B(in io.Reader, out io.Writer) {
	var T, n int
	var a string
	f := func() (c1 int) {
		Fscan(in, &a)
		s := 0
		for _, b := range a {
			s ^= int(b & 1)
			c1 += s
		}
		return
	}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		c1 := f()
		c2 := f()
		if c1 == c2 || c1 == n+1-c2 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1906B(bufio.NewReader(os.Stdin), os.Stdout) }
