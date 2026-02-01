package main

import (
	"math"
	"math/big"
)

// https://space.bilibili.com/206214
type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x } // 如果乘法会溢出，用 detCmp
func (a vec) detCmp(b vec) int {
	v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
	w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
	return v.Cmp(w)
}

func minPartitionScore(nums []int, k int) int64 {
	n := len(nums)
	sum := make([]int, n+1)
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
	}

	for K := 1; K <= k; K++ {
		s := sum[K-1]
		q := []vec{{s, f[K-1] + s*s - s}}
		for i := K; i <= n-(k-K); i++ {
			s = sum[i]
			p := vec{-2 * s, 1}
			for len(q) > 1 && p.dot(q[0]) >= p.dot(q[1]) {
				q = q[1:]
			}

			v := vec{s, f[i] + s*s - s}
			f[i] = p.dot(q[0]) + s*s + s

			// 读者可以把 detCmp 改成 det 感受下这个算法的效率
			// 目前 det 也能过，可以试试 hack 一下
			for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).detCmp(v.sub(q[len(q)-1])) <= 0 {
				q = q[:len(q)-1]
			}
			q = append(q, v)
		}
	}

	return int64(f[n] / 2)
}

//func main() {
//	a := make([]int, 1000)
//	for i := range a {
//		a[i] = 1e4
//	}
//	fmt.Println(minPartitionScore(a, 2))
//}
