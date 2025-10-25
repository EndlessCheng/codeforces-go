package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
const mod93 = 1_000_000_007

type matrix93 [][]int

func newMatrix93(n, m int) matrix93 {
	a := make(matrix93, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix93) mul(b matrix93) matrix93 {
	c := newMatrix93(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod93
			}
		}
	}
	return c
}

func (a matrix93) powMul(n int, f0 matrix93) matrix93 {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func f93(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 4
	}

	m := newMatrix93(17, 17)
	f2 := newMatrix93(17, 1)
	f2[16][0] = 4 // 长度等于 1 的方案数
	// 0123 = 白黑红黄
	for i := range 4 {
		for j := range 4 {
			if j == i || i+j == 3 {
				continue
			}
			for k := range 4 {
				if k != j && j+k != 3 && (j != 0 || i+k != 3) {
					m[i*4+j][j*4+k] = 1
					m[16][j*4+k]++ // 计算 m 矩阵 j*4+k 列之和
				}
			}
			f2[i*4+j][0] = 1
			f2[16][0]++ // 长度等于 2 的方案数
		}
	}
	m[16][16] = 1

	return m.powMul(n-2, f2)[16][0]
}

// (f(n) - 回文方案数) / 2 + 回文方案数 = (f(n) + 回文方案数) / 2
// 回文方案数：长度恰好等于偶数不可能（中间两个颜色相同），此时 <= n 变成 <= n-1
// n 为奇数时，根据左半就能确定右半，所以只有前 ceil(n/2) 个颜色需要计算方案数
func solve93(n int) int {
	const inv2 = (mod93 + 1) / 2
	return (f93(n) + f93((n+1)/2)) * inv2
}

func cf93D(in io.Reader, out io.Writer) {
	var l, r int
	Fscan(in, &l, &r)
	ans := solve93(r) - solve93(l-1)
	Fprintln(out, (ans%mod93+mod93)%mod93)
}

//func main() { cf93D(os.Stdin, os.Stdout) }
