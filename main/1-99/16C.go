package __99

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF16C(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var a, b, x, y int
	Fscan(in, &a, &b, &x, &y)
	g := gcd(x, y)
	x /= g
	y /= g
	k := a / x
	if b/y < k {
		k = b / y
	}
	if k > 0 {
		Fprint(out, x*k, y*k)
	} else {
		Fprint(out, 0, 0)
	}
}

//func main() { CF16C(os.Stdin, os.Stdout) }
