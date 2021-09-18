package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, ans int
	Fscan(in, &n, &m)
	cnt := make([][2]int, m)
	for ; n > 0; n-- {
		c := [2]int{}
		for j := 0; j < m; j++ {
			Fscan(in, &v)
			c[v]++
			cnt[j][v]++
		}
		ans += 1<<c[0] + 1<<c[1] - 2 // 去掉空集
	}
	for _, c := range cnt {
		ans += 1<<c[0] + 1<<c[1] - c[0] - c[1] - 2 // 去掉只有一个元素的集合，以及空集
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
