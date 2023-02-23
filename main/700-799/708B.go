package main

import (
	. "fmt"
	"io"
	"math"
	"strings"
)

// https://space.bilibili.com/206214
func CF708B(in io.Reader, out io.Writer) {
	var c00, c01, c10, c11 int
	Fscan(in, &c00, &c01, &c10, &c11)
	f := func(c int) int {
		x := (1 + int(math.Sqrt(float64(1+8*int64(c))))) / 2
		if x*(x-1)/2 == c {
			return x
		}
		return -1
	}
	c0, c1 := f(c00), f(c11)
	if c0 < 0 || c1 < 0 {
		Fprint(out, "Impossible")
		return
	}

	if c01 == 0 && c10 == 0 {
		if c00 > 0 && c11 > 0 {
			Fprint(out, "Impossible")
		} else if c11 == 0 {
			Fprint(out, strings.Repeat("0", c0))
		} else {
			Fprint(out, strings.Repeat("1", c1))
		}
		return
	}

	if c01+c10 != c0*c1 {
		Fprint(out, "Impossible")
		return
	}
	left1, right0 := c10/c0, c10%c0
	ans := strings.Repeat("1", left1) + strings.Repeat("0", c0-right0)
	if right0 > 0 {
		ans += "1" + strings.Repeat("0", right0)
		c1--
	}
	ans += strings.Repeat("1", c1-left1)
	Fprint(out, ans)
}

//func main() { CF708B(os.Stdin, os.Stdout) }
