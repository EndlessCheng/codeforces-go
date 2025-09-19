package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1996D(in io.Reader, out io.Writer) {
	var T, n, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		ans := 0
		for a := 1; a < min(n, x-1); a++ {
			for b := 1; b <= min(n/a, x-a); b++ {
				ans += min((n-a*b)/(a+b), x-a-b)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1996D(bufio.NewReader(os.Stdin), os.Stdout) }
