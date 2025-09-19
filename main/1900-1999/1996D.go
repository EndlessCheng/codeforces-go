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
		for a := 1; ; a++ {
			c := min((n-a*a)/(a*2), x-a*2)
			if c <= 0 {
				break
			}
			ans += c * 3
			if c >= a {
				ans -= 2
			}
			for b := a + 1; ; b++ {
				d := min((n-a*b)/(a+b), x-a-b) - b
				if d <= 0 {
					break
				}
				ans += d * 6
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1996D(bufio.NewReader(os.Stdin), os.Stdout) }
