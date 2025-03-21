package main

import (
	"math/big"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func minZeroArray(nums []int, queries [][]int) int {
	ans := (slices.Max(nums) + 9) / 10
	m := len(queries)
	cnts := make([][11]int, m+1)
	for i, x := range nums {
		if x == 0 {
			continue
		}
		for k, q := range queries {
			cnts[k+1] = cnts[k]
			if q[0] <= i && i <= q[1] {
				cnts[k+1][q[2]]++
			}
		}
		ans += sort.Search(m+1-ans, func(mx int) bool {
			mx += ans
			p := new(big.Int)
			f := big.NewInt(1)
			for v, num := range cnts[mx] {
				for pow2 := 1; num > 0; pow2 *= 2 {
					k := min(pow2, num)
					f.Or(f, p.Lsh(f, uint(v*k)))
					if f.Bit(x) > 0 {
						return true
					}
					num -= k
				}
			}
			return false
		})
		if ans > m {
			return -1
		}
	}
	return ans
}

func minZeroArray42(nums []int, queries [][]int) int {
	m := len(queries)
	cnts := make([][][11]int, m+1)
	cnts[0] = make([][11]int, len(nums))
	for k, q := range queries {
		cnts[k+1] = slices.Clone(cnts[k])
		for i := q[0]; i <= q[1]; i++ {
			cnts[k+1][i][q[2]]++
		}
	}
	ans := sort.Search(m+1, func(mx int) bool {
		p := new(big.Int)
	next:
		for i, x := range nums {
			if x == 0 {
				continue
			}
			// 多重背包（二进制优化）
			f := big.NewInt(1)
			for v, num := range cnts[mx][i] {
				for pow2 := 1; num > 0; pow2 *= 2 {
					k := min(pow2, num)
					f.Or(f, p.Lsh(f, uint(v*k)))
					if f.Bit(x) > 0 {
						continue next
					}
					num -= k
				}
			}
			return false
		}
		return true
	})
	if ans <= m {
		return ans
	}
	return -1
}

func minZeroArray4(nums []int, queries [][]int) int {
	ans := sort.Search(len(queries)+1, func(mx int) bool {
		p := new(big.Int)
	next:
		for i, x := range nums {
			if x == 0 {
				continue
			}
			cnt := [11]int{}
			for _, q := range queries[:mx] {
				if q[0] <= i && i <= q[1] {
					cnt[q[2]]++
				}
			}
			// 多重背包（二进制优化）
			f := big.NewInt(1)
			for v, num := range cnt {
				for pow2 := 1; num > 0; pow2 *= 2 {
					k := min(pow2, num)
					f.Or(f, p.Lsh(f, uint(v*k)))
					if f.Bit(x) > 0 {
						continue next
					}
					num -= k
				}
			}
			return false
		}
		return true
	})
	if ans <= len(queries) {
		return ans
	}
	return -1
}

func minZeroArray3(nums []int, queries [][]int) (ans int) {
	p := new(big.Int)
	for i, x := range nums {
		if x == 0 {
			continue
		}
		f := big.NewInt(1)
		for k, q := range queries {
			if q[0] <= i && i <= q[1] {
				f.Or(f, p.Lsh(f, uint(q[2])))
			}
			if f.Bit(x) > 0 {
				ans = max(ans, k+1)
				break
			}
		}
		if f.Bit(x) == 0 {
			return -1
		}
	}
	return
}

func minZeroArray2(nums []int, queries [][]int) int {
	if !slices.ContainsFunc(nums, func(x int) bool { return x > 0 }) {
		return 0
	}

	f := make([]*big.Int, len(nums))
	for i := range f {
		f[i] = big.NewInt(1)
	}
	p := new(big.Int)
next:
	for k, q := range queries {
		val := uint(q[2])
		for i := q[0]; i <= q[1]; i++ {
			if f[i].Bit(nums[i]) == 0 {
				f[i].Or(f[i], p.Lsh(f[i], val))
			}
		}
		for i, x := range nums {
			if f[i].Bit(x) == 0 {
				continue next
			}
		}
		return k + 1
	}
	return -1
}

func minZeroArray1(nums []int, queries [][]int) int {
	if !slices.ContainsFunc(nums, func(x int) bool { return x > 0 }) {
		return 0
	}

	n := len(nums)
	f := make([][]bool, n)
	for i, x := range nums {
		f[i] = make([]bool, x+1)
		f[i][0] = true
	}
next:
	for k, q := range queries {
		val := q[2]
		for i := q[0]; i <= q[1]; i++ {
			if f[i][nums[i]] {
				continue
			}
			for j := nums[i]; j >= val; j-- {
				f[i][j] = f[i][j] || f[i][j-val]
			}
		}
		for i, x := range nums {
			if !f[i][x] {
				continue next
			}
		}
		return k + 1
	}
	return -1
}
