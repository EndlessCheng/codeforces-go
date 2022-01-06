package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF893D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, lim, v, low, up, ans int
	for Fscan(in, &n, &lim); n > 0; n-- {
		Fscan(in, &v)
		if v == 0 {
			if up < 0 {
				ans++
				low, up = 0, lim
			} else if low < 0 {
				low = 0
			}
		} else {
			if low += v; low > lim {
				Fprint(out, -1)
				return
			}
			if up += v; up > lim {
				up = lim
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF893D(os.Stdin, os.Stdout) }
