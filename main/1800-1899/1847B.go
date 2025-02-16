package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1847B(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans, and := 0, -1
		for range n {
			Fscan(in, &v)
			and &= v
			if and == 0 {
				ans++
				and = -1
			}
		}
		Fprintln(out, max(ans, 1))
	}
}

//func main() { cf1847B(bufio.NewReader(os.Stdin), os.Stdout) }
