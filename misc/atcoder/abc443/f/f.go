package main

import (
	. "fmt"
	"io"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	from := make([][10]int, n)
	q := []int{}
	for d := 1; d < 10; d++ {
		r := d % n
		from[r][d] = -1
		q = append(q, r*10+d)
	}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		r, d := p/10, p%10
		if r == 0 {
			ans := []byte{'0' + byte(d)}
			for from[r][d] >= 0 {
				p = from[r][d]
				r, d = p/10, p%10
				ans = append(ans, '0'+byte(d))
			}
			slices.Reverse(ans)
			Fprintf(out, "%s", ans)
			return
		}
		for d2 := range 10 {
			if d2 < d {
				continue
			}
			r2 := (r*10 + d2) % n
			if from[r2][d2] == 0 {
				from[r2][d2] = p
				q = append(q, r2*10+d2)
			}
		}
	}
	Fprint(out, -1)
}

func main() { run(os.Stdin, os.Stdout) }
