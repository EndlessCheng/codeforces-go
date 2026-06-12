package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2097C(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var T, n, x, y, dx, dy int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x, &y, &dx, &dy)

		g := gcd(dx, dy)
		dx /= g
		dy /= g

		d := x*dy - y*dx
		if d%n != 0 {
			Fprintln(out, -1)
			continue
		}
		d /= n

		_, p, q := exgcd97(dy, dx)
		p *= d
		q *= -d

		if p <= 0 {
			v := (dx - p) / dx
			p += v * dx
			q += v * dy
		}
		if q <= 0 {
			v := (dy - q) / dy
			p += v * dx
			q += v * dy
		}

		v := min((p-1)/dx, (q-1)/dy)
		p -= v * dx
		q -= v * dy

		Fprintln(out, p-1+q-1+(p+q)/2+abs(p-q)/2)
	}
}

//func main() { cf2097C(bufio.NewReader(os.Stdin), os.Stdout) }

func exgcd97(a, b int) (gcd, x, y int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, y, x = exgcd97(b, a%b)
	y -= a / b * x
	return
}
