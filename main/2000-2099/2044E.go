package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2044E(in io.Reader, out io.Writer) {
	var T, k, lx, rx, ly, ry int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &k, &lx, &rx, &ly, &ry)
		ans := 0
		for pk := 1; pk <= ry; pk *= k {
			l := max(lx, (ly-1)/pk+1)
			r := min(rx, ry/pk)
			ans += max(r-l+1, 0)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2044E(bufio.NewReader(os.Stdin), os.Stdout) }
