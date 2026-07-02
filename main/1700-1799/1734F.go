package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1734F(in io.Reader, out io.Writer) {
	type pair struct{ i, j int }
	memo := map[pair]int{}
	var f func(int, int) int
	f = func(i, j int) (res int) {
		if i == 0 || j == 0 {
			return 0
		}
		p := pair{i, j}
		if v, ok := memo[p]; ok {
			return v
		}
		if i%2 == 0 {
			res = f(i/2, j/2) + f(i/2, (j+1)/2)
		} else {
			res = j - f(i/2, (j+1)/2) - f((i+1)/2, j/2)
		}
		memo[p] = res
		return
	}

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		Fprintln(out, f(n, m))
	}
}

//func main() { cf1734F(bufio.NewReader(os.Stdin), os.Stdout) }
