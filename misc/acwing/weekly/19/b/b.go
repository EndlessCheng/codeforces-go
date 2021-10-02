package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, w, v, s, min, max int
	for Fscan(in, &n, &w); n > 0; n-- {
		Fscan(in, &v)
		s += v
		if s < min { // 记录前缀和的最小值和最大值
			min = s
		} else if s > max {
			max = s
		}
	}
	// 设初值为 x，则需要满足 x+min>=0 且 x+max<=w
	// 得 x∈[-min, w-max]，方案数即为区间长度
	ans := w - max + min + 1
	if ans < 0 {
		ans = 0
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
