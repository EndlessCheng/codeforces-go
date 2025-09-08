package main

import (
	. "fmt"
	"io"
)

// 位置为 r 的僵尸将会猜对，当且仅当，所有真实数字的总和 S 模 n 的结果恰好等于该僵尸的位置 r。

// https://github.com/EndlessCheng
func cf690A3(in io.Reader, out io.Writer) {
	var T, n, r, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &r)
		s := 0
		for range n - 1 {
			Fscan(in, &x)
			s += x
		}
		Fprintln(out, (r-s%n+n)%n+1)
	}
}

//func main() { cf690A3(bufio.NewReader(os.Stdin), os.Stdout) }
