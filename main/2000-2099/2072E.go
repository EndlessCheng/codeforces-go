package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf2072E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, k, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &k)
		a := [][2]int{}
		for ; k > 0; y++ {
			m := (int(math.Sqrt(float64(k*8+1))) + 1) / 2
			for range m {
				a = append(a, [2]int{x, y})
				x++
			}
			k -= m * (m - 1) / 2
		}
		Fprintln(out, len(a))
		for _, p := range a {
			Fprintln(out, p[0], p[1])
		}
	}
}

//func main() { cf2072E(bufio.NewReader(os.Stdin), os.Stdout) }
