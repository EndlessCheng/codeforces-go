package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF991D(in io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	var s1, s2 []byte
	Fscan(bufio.NewReader(in), &s1, &s2)
	const inf int = 1e9
	shape := [][3]int{{3, 1}, {1, 3}, {2, 3}, {3, 2}} // 依次对应题目描述的 4 种 L 形状

	f := [4]int{-inf, -inf, -inf, 0} // 第 -1 列视作都是 X
	for i, b := range s1 {
		cur := int(b>>6 | s2[i]>>6<<1) // 第 i 列的 X
		nf := [4]int{-inf, -inf, -inf, -inf}
		nf[cur] = max(max(max(f[0], f[1]), f[2]), f[3]) // 不填 L
		for _, p := range shape { // 填 L，枚举 L 形状
			for pre := 0; pre < 4; pre++ { // 枚举第 i-1 列
				if p[0]&pre == 0 && p[1]&cur == 0 { // 可以填 L
					nf[p[1]|cur] = max(nf[p[1]|cur], f[pre]+1)
				}
			}
		}
		f = nf
	}
	Fprint(out, max(max(max(f[0], f[1]), f[2]), f[3]))
}

//func main() { CF991D(os.Stdin, os.Stdout) }
