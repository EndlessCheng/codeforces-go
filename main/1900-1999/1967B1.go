package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1967B1(in io.Reader, out io.Writer) {
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		ans := -1
		for b := 1; b <= min(m, n/b+1); b++ {
			ans += (n/b + 1) / b
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1967B1(bufio.NewReader(os.Stdin), os.Stdout) }
