package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// 反思：这题能不能一次过？
// 0. 首先 O(26n) 当然是可以做的，但是我想试试 O(n) 的做法
// 想出 O(n) 的做法需要想明白一些细节：
// 1. 确认清楚目标是什么：确保所有 L 都被覆盖，即对位置 p，覆盖 [p-m+1,p]，最后检查 [0, n-m+1] 都被覆盖了（这是 L 的范围）
// 2. 遍历时，每次先计算出下一个没有被覆盖的位置，这里用 find(0) 来计算
// 3. 然后贪心地算出最靠右的覆盖 find(0) 的位置，若没有则当前字母全选，进入下一个字母

// github.com/EndlessCheng/codeforces-go
func CF724D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var m int
	var s []byte
	Fscan(in, &m, &s)
	n := len(s)
	pos := [26][]int{}
	for i, b := range s {
		b -= 'a'
		pos[b] = append(pos[b], i)
	}
	var fa []int
	initFa := func(n int) {
		fa = make([]int, n+1)
		for i := range fa {
			fa[i] = i
		}
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	mergeRange := func(l, r int) (merged bool) {
		if l < 0 {
			l = 0
		}
		for i := find(l); i < r; i = find(i + 1) {
			fa[i] = r
			merged = true
		}
		return
	}

	initFa(n)
	ans := make([]byte, 0, n)
outer:
	for i, ps := range pos {
		b := byte(i + 'a')
		left := len(ps)
		for j := 0; j < len(ps); j++ {
			check := find(0)
			found := false
			for ; j < len(ps); j++ {
				if ps[j]-m+1 > check {
					break
				}
				found = true
			}
			if !found {
				for _, p := range ps[j:] {
					mergeRange(p-m+1, p+1)
				}
				break
			}
			if j > 0 {
				j--
			}
			p := ps[j]
			if mergeRange(p-m+1, p+1) {
				ans = append(ans, b)
				left--
				if find(0) >= n-m+1 {
					break outer
				}
			}
		}
		ans = append(ans, bytes.Repeat([]byte{b}, left)...)
	}
	Fprint(out, string(ans))
}

//func main() {
//	CF724D(os.Stdin, os.Stdout)
//}
