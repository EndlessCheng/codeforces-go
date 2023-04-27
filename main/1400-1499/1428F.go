package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1428F(in io.Reader, out io.Writer) {
	var n, ones int
	var s string
	Fscan(bufio.NewReader(in), &n, &s)
	var ans, sum int64
	last := make([]int, n+1)
	for i := range last {
		last[i] = -1
	}
	for i, c := range s {
		if c == '1' {
			// 1111011
			// 11110111
			//  | 从这里到 i 的左开右闭区间的（以 i 为右端点的）子串都多了 1
			ones++
			sum += int64(i - last[ones])
		} else {
			// 更新上一个连续 ones 个 1 的起始位置
			for ; ones > 0; ones-- {
				last[ones] = i - ones
			}
		}
		ans += sum
	}
	Fprint(out, ans)
}

//func main() { CF1428F(os.Stdin, os.Stdout) }
