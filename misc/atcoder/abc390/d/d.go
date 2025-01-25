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
	var dfs func(int, int)
	dfs = func(i, xor int) {
		if i == n {
			ans[xor] = true
			return
		}
		v := a[i]
		// v 单独组成一个集合
		b = append(b, v)
		dfs(i+1, xor^v)
		b = b[:len(b)-1]
		// v 加到前面的集合中
		for j := range b {
			b[j] += v
			dfs(i+1, xor^(b[j]-v)^b[j])
			b[j] -= v
		}
	}
	dfs(0, 0)
	Fprint(out, len(ans))
}

func main() { run(os.Stdin, os.Stdout) }
