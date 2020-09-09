package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF766B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	// 官方题解的解法二 https://codeforces.com/blog/entry/50294
	// 当数列的增长速度不低于斐波那契数列时，是无法组成三角形的。由于数列上界为 1e9，斐波那契数列最多有 44 项
	if n > 44 {
		Fprint(out, "YES")
		return
	}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	for i, v := range a[2:] {
		if a[i]+a[i+1] > v {
			Fprint(out, "YES")
			return
		}
	}
	Fprint(out, "NO")
}

//func main() { CF766B(os.Stdin, os.Stdout) }
