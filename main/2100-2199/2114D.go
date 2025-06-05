package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2114D(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]struct{ x, y int }, n)
		var xi, xj, yi, yj int
		for i := range a {
			Fscan(in, &a[i].x, &a[i].y)
			if a[i].x < a[xi].x {
				xi = i
			}
			if a[i].x > a[xj].x {
				xj = i
			}
			if a[i].y < a[yi].y {
				yi = i
			}
			if a[i].y > a[yj].y {
				yj = i
			}
		}
		if n == 1 {
			Fprintln(out, 1)
			continue
		}
		f := func(ban int) int {
			minX, maxX, minY, maxY := int(1e9), 0, int(1e9), 0
			for i, p := range a {
				if i == ban {
					continue
				}
				minX = min(minX, p.x)
				maxX = max(maxX, p.x)
				minY = min(minY, p.y)
				maxY = max(maxY, p.y)
			}
			dx, dy := maxX-minX+1, maxY-minY+1
			if dx*dy > n-1 {
				return dx * dy
			}
			return min((dx+1)*dy, dx*(dy+1))
		}
		Fprintln(out, min(f(xi), f(xj), f(yi), f(yj)))
	}
}

//func main() { cf2114D(bufio.NewReader(os.Stdin), os.Stdout) }
