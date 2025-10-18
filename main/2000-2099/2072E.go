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
	var T, k, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &k)
		a := []int{}
		for ; k > 0; y++ {
			m := (int(math.Sqrt(float64(k*8+1))) + 1) / 2
			for range m {
				a = append(a, y)
			}
			k -= m * (m - 1) / 2
		}
		Fprintln(out, len(a))
		for x, y := range a {
			Fprintln(out, x, y)
		}
	}
}

//func main() { cf2072E(bufio.NewReader(os.Stdin), os.Stdout) }
