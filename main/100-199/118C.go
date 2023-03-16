package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
func CF118C(in io.Reader, out io.Writer) {
	var n, k int
	var s, ans string
	Fscan(bufio.NewReader(in), &n, &k, &s)
	pos := ['9' + 1][]int{}
	for i, b := range s {
		pos[b] = append(pos[b], i)
	}

	minCost := math.MaxInt
	for c := byte('0'); c <= '9'; c++ {
		t := []byte(s)
		cost := 0
		left := k - len(pos[c])
		for d := byte(1); d <= 9; d++ {
			if c+d <= '9' { // 先改比 c 大的
				p := pos[c+d]
				for i := 0; i < len(p) && left > 0; i++ { // 正序
					t[p[i]] = c
					left--
					cost += int(d)
				}
			}
			if c-d >= '0' { // 再改比 c 小的
				p := pos[c-d]
				for i := len(p) - 1; i >= 0 && left > 0; i-- { // 逆序
					t[p[i]] = c
					left--
					cost += int(d)
				}
			}
		}
		res := string(t)
		if cost < minCost || cost == minCost && res < ans {
			minCost = cost
			ans = res
		}
	}
	Fprintln(out, minCost)
	Fprint(out, ans)
}

//func main() { CF118C(os.Stdin, os.Stdout) }
