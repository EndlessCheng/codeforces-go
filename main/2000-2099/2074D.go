package main

import (
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf2074D(in io.Reader, out io.Writer) {
	var T, n, m, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		x := make([]int, n)
		for i := range x {
			Fscan(in, &x[i])
		}
		maxY2 := map[int]int{}
		for _, x := range x {
			Fscan(in, &r)
			for i := x - r; i <= x+r; i++ {
				maxY2[i] = max(maxY2[i], r*r-(x-i)*(x-i))
			}
		}
		ans := 0
		for _, y2 := range maxY2 {
			ans += 1 + int(math.Sqrt(float64(y2)))*2
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2074D(bufio.NewReader(os.Stdin), os.Stdout) }
