package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf416D(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i := 0; i < n; {
		ans++
		st := i
		x0, y0 := -1, -1
		k := int(2e9)
		expY := 0
		for ; i < n; i++ {
			y := a[i]
			expY += k
			if k < 2e9 {
				if expY < 1 || y > 0 && y != expY {
					break
				}
				continue
			}
			if y < 0 {
				continue
			}
			if x0 < 0 {
				x0, y0 = i, y
				continue
			}
			dx, dy := i-x0, y-y0
			k = dy / dx
			if dy%dx != 0 || y-(i-st)*k < 1 {
				break
			}
			expY = y
		}
	}
	Fprint(out, ans)
}

//func main() { cf416D(bufio.NewReader(os.Stdin), os.Stdout) }
