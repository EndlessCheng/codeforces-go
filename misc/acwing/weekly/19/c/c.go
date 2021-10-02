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
	var n, k, v, ans int
	min, max := int(2e5), 0
	cnt := [2e5 + 1]int{}
	for Fscan(in, &n, &k); n > 0; n-- {
		Fscan(in, &v)
		cnt[v]++
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	if min == max { // 无需操作
		Fprint(out, 0)
		return
	}

	// 从第 max 层开始往下取石子
	for i, c := max, 0; i > min; i-- {
		if c+cnt[i] > k { // 贪心：当石子个数累计超过 k 时才进行一轮操作，取出高度大于 i 的剩余石子
			ans++
			c = cnt[i]
		} else {
			c += cnt[i]
		}
		cnt[i-1] += cnt[i] // 累加到低一级的石子堆个数上
	}
	Fprint(out, ans+1) // 最后剩余的 c 枚石子再进行一轮操作
}

func main() { run(os.Stdin, os.Stdout) }
