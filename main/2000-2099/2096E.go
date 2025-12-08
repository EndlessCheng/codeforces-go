package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2096E(in io.Reader, out io.Writer) {
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		var inv, d, r int
		for _, c := range s {
			if c == 'P' {
				r++
				continue
			}
			inv += r
			if d%2 != r%2 {
				d++
			} else if d > 0 {
				d--
			}
		}
		// 每两个无法配对的 B 可共用 1 次最小收益操作（消除 1 个逆序对），最后剩余的 1 个 B 需要 1 次最小收益操作（再消除 1 个）
		d = (d + 1) / 2
		Fprintln(out, (inv+d)/2)
	}
}

//func main() { cf2096E(bufio.NewReader(os.Stdin), os.Stdout) }
