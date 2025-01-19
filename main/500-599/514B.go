package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf514B(in io.Reader, out io.Writer) {
	var n, x0, y0, x, y, ans int
	Fscan(in, &n, &x0, &y0)
	has := map[float64]bool{}
	for range n {
		Fscan(in, &x, &y)
		if x == x0 {
			ans = 1
		} else {
			has[float64(y-y0)/float64(x-x0)] = true
		}
	}
	Fprintln(out, ans+len(has))
}

//func main() { cf514B(bufio.NewReader(os.Stdin), os.Stdout) }
