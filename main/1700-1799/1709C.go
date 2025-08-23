package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1709C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	T, s := 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		c, q := 0, 0
		for _, b := range s {
			if b == '?' {
				q++
			} else {
				c += 1 - int(b%2*2)
			}
			if c+q == 1 {
				c, q = 1, 0
			}
		}
		if c == q || -c == q {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1709C(bufio.NewReader(os.Stdin), os.Stdout) }
