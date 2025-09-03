package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1373C(in io.Reader, out io.Writer) {
	T, s := 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		ans := len(s)
		c, mn := 0, 0
		for i, b := range s {
			c += int(b&2 - 1)
			if c < mn {
				mn = c
				ans += i + 1
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1373C(bufio.NewReader(os.Stdin), os.Stdout) }
