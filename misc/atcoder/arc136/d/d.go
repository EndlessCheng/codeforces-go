package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	f := [1e6]int{}
	for i := range a {
		Fscan(in, &a[i])
		f[a[i]]++
	}
	// 注意循环顺序，需要保证先计算完 f[i-p10]，再计算 f[i] 
	for p10 := 1; p10 < 1e6; p10 *= 10 {
		for i := range f {
			if i/p10%10 > 0 {
				f[i] += f[i-p10]
			}
		}
	}
next:
	for _, x := range a {
		ans += f[999999-x]
		for ; x > 0; x /= 10 {
			if x%10 > 4 {
				continue next
			}
		}
		ans-- // 多统计了 x+x 的情况
	}
	Fprint(out, ans/2)
}

func main() { run(os.Stdin, os.Stdout) }
