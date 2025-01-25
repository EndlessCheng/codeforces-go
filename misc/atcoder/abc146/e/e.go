package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, k, ans int
	Fscan(in, &n, &k)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] = (s[i] + s[i-1] - 1) % k // 题目保证元素都是正数
	}

	cnt := map[int]int{}
	for i, v := range s {
		ans += cnt[v]
		cnt[v]++
		if i >= k-1 {
			cnt[s[i-k+1]]--
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
