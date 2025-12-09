package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1375E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ans := [][2]int{}
	b := slices.Clone(a)
	for i := n - 1; i >= 0; i-- {
		id := []int{}
		for j, v := range a[:i] {
			if v > a[i] {
				id = append(id, j)
			}
		}
		slices.SortFunc(id, func(i, j int) int { return a[i] - a[j] })
		for _, j := range id {
			ans = append(ans, [2]int{j, i})
			b[j], b[i] = b[i], b[j]
		}
	}

	Fprintln(out, len(ans))
	for _, p := range ans {
		Fprintln(out, p[0]+1, p[1]+1)
	}
}

//func main() { cf1375E(bufio.NewReader(os.Stdin), os.Stdout) }
