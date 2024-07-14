package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func p3173(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	horizontalCut := make([]int, n-1)
	for i := range horizontalCut {
		Fscan(in, &horizontalCut[i])
	}
	verticalCut := make([]int, m-1)
	for i := range verticalCut {
		Fscan(in, &verticalCut[i])
	}
	slices.SortFunc(horizontalCut, func(a, b int) int { return b - a })
	slices.SortFunc(verticalCut, func(a, b int) int { return b - a })
	ans := 0
	i, j := 0, 0
	for i < n-1 || j < m-1 {
		if j == m-1 || i < n-1 && horizontalCut[i] > verticalCut[j] {
			ans += horizontalCut[i] * (j + 1)
			i++
		} else {
			ans += verticalCut[j] * (i + 1)
			j++
		}
	}
	Fprint(out, ans)
}

//func main() { p3173(bufio.NewReader(os.Stdin), os.Stdout) }
