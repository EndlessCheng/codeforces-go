package main

import (
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ans := map[int]bool{}
	b := []int{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			xor := 0
			for _, v := range b {
				xor ^= v
			}
			ans[xor] = true
			return
		}
		v := a[i]
		// v 单独组成一个集合
		b = append(b, v)
		dfs(i + 1)
		b = b[:len(b)-1]
		// v 加到前面的集合中
		for j := range b {
			b[j] += v
			dfs(i + 1)
			b[j] -= v
		}
	}
	dfs(0)
	Fprint(out, len(ans))
}

func main() { run(os.Stdin, os.Stdout) }
