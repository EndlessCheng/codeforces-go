package main

import (
	"bufio"
	. "fmt"
	"io"
)

// 注：也可以用 strings.Count，毕竟自带 RK

// github.com/EndlessCheng/codeforces-go
func CF1003F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	text := make([]int, n)  // 单词编号数组
	sum := make([]int, n+1) // 单词长度前缀和
	allSum := n - 1         // 总长度
	id := map[string]int{}
	for i := range text {
		var s string
		Fscan(in, &s)
		if _, has := id[s]; !has {
			id[s] = len(id)
		}
		text[i] = id[s]
		sum[i+1] = sum[i] + len(s)
		allSum += len(s)
	}

	// KMP: s 在 text 中的不相交出现次数
	count := func(s []int) (cnt int) {
		n := len(s)
		f := make([]int, n)
		c := 0
		for i := 1; i < n; i++ {
			b := s[i]
			for c > 0 && s[c] != b {
				c = f[c-1]
			}
			if s[c] == b {
				c++
			}
			f[i] = c
		}
		c = 0
		for _, b := range text {
			for c > 0 && s[c] != b {
				c = f[c-1]
			}
			if s[c] == b {
				c++
			}
			if c == n {
				cnt++
				c = 0
			}
		}
		return
	}

	ans := allSum
	for r := 1; r <= n; r++ {
		for l := 0; l < r; l++ {
			if c := count(text[l:r]); c > 1 {
				if s := allSum - c*(sum[r]-sum[l]-1); s < ans {
					ans = s
				}
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF1003F(os.Stdin, os.Stdout) }
