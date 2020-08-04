package __99

import (
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF20B(in io.Reader, out io.Writer) {
	var a, b, c int64
	Fscan(in, &a, &b, &c)
	if a == 0 {
		if b == 0 {
			if c == 0 {
				Fprint(out, -1)
			} else {
				Fprint(out, 0)
			}
		} else {
			Fprintf(out, "1\n%.5f", float64(-c)/float64(b))
		}
	} else {
		d := b*b - 4*a*c
		if d < 0 {
			Fprint(out, 0)
		} else if d == 0 {
			Fprintf(out, "1\n%.5f", float64(-b)/float64(2*a))
		} else {
			dd := math.Sqrt(float64(d))
			x, y := (float64(-b)-dd)/float64(2*a), (float64(-b)+dd)/float64(2*a)
			if x > y {
				x, y = y, x
			}
			Fprintf(out, "2\n%.5f %.5f", x, y)
		}
	}
}

//func main() { CF20B(os.Stdin, os.Stdout) }
