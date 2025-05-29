package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1974C(in io.Reader, out io.Writer) {
	var T, n, x, y, z int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x, &y)
		ans := 0
		c12 := map[[2]int]int{}
		c13 := map[[2]int]int{}
		c23 := map[[2]int]int{}
		c3 := map[[3]int]int{}
		for range n - 2 {
			Fscan(in, &z)
			ans += c12[[2]int{x, y}] + c13[[2]int{x, z}] + c23[[2]int{y, z}] - c3[[3]int{x, y, z}]*3
			c12[[2]int{x, y}]++
			c13[[2]int{x, z}]++
			c23[[2]int{y, z}]++
			c3[[3]int{x, y, z}]++
			x, y = y, z
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1974C(bufio.NewReader(os.Stdin), os.Stdout) }
