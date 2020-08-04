package __99

import (
	"bufio"
	. "fmt"
	"io"
)

func exgcd7C(a, b int64) (gcd, x, y int64) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, y, x = exgcd7C(b, a%b)
	y -= a / b * x
	return
}

// github.com/EndlessCheng/codeforces-go
func Sol7C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var a, b, c int64
	Fscan(in, &a, &b, &c)
	c = -c
	gcd, x, y := exgcd7C(a, b)
	if c%gcd != 0 {
		Fprint(out, -1)
	} else {
		c /= gcd
		Fprint(out, c*x, c*y)
	}
}

//func main() {
//	Sol7C(os.Stdin, os.Stdout)
//}
