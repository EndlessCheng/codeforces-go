package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF10C(_r io.Reader, _w io.Writer) { // 方便测试，见 10C_test.go
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	ans := int64(0)
	cnt := [9]int64{}
	for i := 1; i <= n; i++ {
		cnt[i%9]++
		ans -= int64(n / i)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			ans += cnt[i] * cnt[j] * cnt[i*j%9]
		}
	}
	Fprint(out, ans)
}

//func main() { CF10C(os.Stdin, os.Stdout) }
