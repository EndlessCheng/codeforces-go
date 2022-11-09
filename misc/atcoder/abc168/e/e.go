package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
const mod int = 1e9 + 7

func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, x, y, cntO, left int
	type pair struct{ x, y int }
	cnt := map[pair]int{}
	cnt2 := map[pair]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &x, &y)
		if x == 0 && y == 0 {
			cntO++
			continue
		}
		// 标准化
		if y < 0 {
			x = -x
			y = -y
		} else if y == 0 {
			x = abs(x)
		}
		g := gcd(abs(x), y)
		x /= g
		y /= g
		if x > 0 {
			cnt[pair{x, y}]++ // 在第一象限或 x 正半轴的向量
		} else {
			cnt2[pair{y, -x}]++ // 在第二象限或 y 正半轴的向量 => 有可能与 cnt 中的向量垂直
			left++
		}
	}

	ans := 1
	for p, c := range cnt {
		ans = ans * (pow(2, c) + pow(2, cnt2[p]) - 1) % mod // 单独选 p + 单独选与 p 垂直的（注意空集重复统计了一次）
		left -= cnt2[p]
	}
	Fprint(out, (ans*pow(2, left)+cntO-1+mod)%mod) // 最后单独选 cnt2 中的，-1 表示减去空集
}

func main() { run(os.Stdin, os.Stdout) }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
