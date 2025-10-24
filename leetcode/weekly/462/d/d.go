package main

import (
	"bytes"
	"math"
	"math/bits"
	"slices"
	"sort"
	"strconv"
)

// https://space.bilibili.com/206214
// 从 a 中选一个字典序最小的、元素和等于 target 的子序列
// a 已经从小到大排序
// 无解返回 nil
func zeroOneKnapsack(a []int, target int) []int {
	n := len(a)
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, target+1)
	}
	f[n][0] = true

	// 倒着 DP，这样后面可以正着（从小到大）选
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		for j := range f[i] {
			if j < v {
				f[i][j] = f[i+1][j]
			} else {
				f[i][j] = f[i+1][j] || f[i+1][j-v]
			}
		}
	}

	if !f[0][target] {
		return nil
	}

	ans := []int{}
	j := target
	for i, v := range a {
		if j >= v && f[i+1][j-v] {
			ans = append(ans, v)
			j -= v
		}
	}
	return ans
}

func specialPalindrome(num int64) int64 {
	s := strconv.FormatInt(num, 10)
	m := len(s)
	mid := (m - 1) / 2

	const mx = 10
	cnt := make([]int, mx)
	for _, d := range s[:mid+1] {
		cnt[d-'0'] += 2
	}
	valid := func() bool {
		for i, c := range cnt {
			if c > 0 && c != i {
				return false
			}
		}
		return true
	}

	// 首先，单独处理中间位置
	tmp := []byte(s[:m/2])
	slices.Reverse(tmp)
	pal, _ := strconv.ParseInt(s[:mid+1]+string(tmp), 10, 64)
	if m%2 == 0 {
		// 不修改
		if pal > num && valid() {
			return pal
		}
	} else {
		// 修改正中间
		cnt[s[mid]-'0'] -= 2
		for j := s[mid] - '0'; j < mx; j++ {
			cnt[j]++
			if pal > num && valid() {
				return pal
			}
			cnt[j]--
			pal += int64(math.Pow10(m / 2))
		}
	}

	// 下面正式开始枚举

	// 生成答案
	buildAns := func(t []byte, missing []int, midD byte) int64 {
		for _, v := range missing {
			cnt[v*2] = -v * 2 // 用负数表示可以随便填的数
		}

		for k, c := range cnt {
			if c > 0 {
				c = k - c
			} else {
				c = -c
				cnt[k] = 0 // 还原
			}
			d := []byte{'0' + byte(k)}
			t = append(t, bytes.Repeat(d, c/2)...) // 只考虑左半
		}

		right := slices.Clone(t)
		slices.Reverse(right)

		if midD > 0 {
			t = append(t, '0'+midD)
		}

		t = append(t, right...)

		ans, _ := strconv.ParseInt(string(t), 10, 64)
		return ans
	}

	// 下标 i 填 j 且正中间填 midD（如果 m 是偶数则 midD 是 0）
	solve := func(i int, j, midD byte) int64 {
		// 中间 [i+1, m-2-i] 需要补满 0 < cnt[k] < k 的数字 k，然后左半剩余数位可以随便填
		free := m/2 - 1 - i // 统计左半（不含正中间）可以随便填的数位个数
		odd := 0
		for k, c := range cnt {
			if k < c { // 不合法
				free = -1
				break
			}
			if c > 0 {
				odd += k % 2
				free -= (k - c) / 2
			}
		}
		if free < 0 || odd > m%2 { // 不合法，继续枚举
			return -1
		}

		// 对于可以随便填的数位，计算字典序最小的填法
		a := []int{}
		for k := 2; k < mx; k += 2 {
			if cnt[k] == 0 {
				a = append(a, k/2) // 左半需要填 k/2 个数
			}
		}
		missing := zeroOneKnapsack(a, free)
		if missing == nil {
			return -1
		}

		t := []byte(s[:i+1])
		t[i] = '0' + j
		return buildAns(t, missing, midD)
	}

	// 从右往左尝试
	for i := m/2 - 1; i >= 0; i-- {
		cnt[s[i]-'0'] -= 2 // 撤销

		// 增大 s[i] 为 j
		for j := s[i] - '0' + 1; j < mx; j++ {
			cnt[j] += 2
			if m%2 == 0 {
				ans := solve(i, j, 0)
				if ans != -1 {
					return ans
				}
			} else {
				ans := int64(math.MaxInt)
				// 枚举正中间填 d
				for d := byte(1); d < mx; d += 2 {
					cnt[d]++
					res := solve(i, j, d)
					if res != -1 {
						ans = min(ans, res)
					}
					cnt[d]--
				}
				if ans != math.MaxInt {
					return ans
				}
			}
			cnt[j] -= 2
		}
	}

	// 没找到，返回长为 m+1 的最小回文数
	return specialPalindrome(int64(math.Pow10(m)))
}

// 没找到，返回长为 m+1 的最小回文数
//a := make([]int, (mx-1)/2)
//for i := range a {
//	a[i] = i + 1
//}
//
//// m+1 是偶数
//if m%2 > 0 {
//	missing := zeroOneKnapsack(a, (m+1)/2)
//	return buildAns(nil, missing, 0)
//}
//
//// m+1 是奇数
//ans := int64(math.MaxInt)
//// 枚举正中间填 midD
//for midD := 1; midD < mx && midD/2 <= m/2; midD += 2 {
//	cnt[midD] = -midD
//	missing := zeroOneKnapsack(a, m/2-midD/2)
//	res := buildAns(nil, missing, byte(midD))
//	if res != -1 {
//		ans = min(ans, res)
//	}
//	cnt[midD] = 0
//}
//return ans

//

func specialPalindrome3(Num int64) int64 {
	subSum := make([]int, 1<<9)
	a := []int{1, 9, 8, 7, 6, 5, 4, 3, 2}
	mp := [10]int{}
	oddMask := 0
	for i, v := range a {
		mp[v] = i
		if v%2 > 0 {
			oddMask |= 1 << i
		}
		highBit := 1 << i
		for mask, s := range subSum[:highBit] {
			subSum[highBit|mask] = s + v
		}
	}
	for mask := 1; mask < 1<<9; mask++ {
		t := mask & oddMask
		if t&(t-1) > 0 {
			subSum[mask] = -1
		}
	}

	full := func(pal, odd int) int {
		v := pal
		if odd > 0 {
			pal = pal*10 + odd
		}
		for ; v > 0; v /= 10 {
			pal = pal*10 + v%10
		}
		return pal
	}

	cnt := [10]int{}
	extend := func(pal, mask, odd int) int {
		for j := 1; j < 10; j++ {
			need := 0
			if mask>>mp[j]&1 > 0 {
				need = j
			} else if cnt[j] > 0 {
				need = j - cnt[j]
			}
			for range need / 2 {
				pal = pal*10 + j
			}
			if need > 0 && j%2 > 0 {
				odd = j
			}
		}
		return full(pal, odd)
	}

	num := int(Num)
	s := strconv.Itoa(num)
	n := len(s)

	ans := 0
	var dfs func(int, int, int, bool) bool
	dfs = func(i, pal, odd int, limit bool) bool {
		if odd > 0 && n%2 == 0 {
			return false
		}

		if i == n/2 {
			if n%2 == 0 {
				for j, c := range cnt {
					if c > 0 && c != j {
						return false
					}
				}
				// 左半反转到右半
				pal = full(pal, 0)
				if pal > num {
					ans = pal
				}
				return pal > num
			} else { // else 可以省略，为了代码格式对齐保留
				if odd == 0 {
					odd = 1
				}
				cnt[odd]++
				defer func() { cnt[odd]-- }()
				for j, c := range cnt {
					if c > 0 && c != j {
						return false
					}
				}
				// 左半反转到右半
				pal = full(pal, odd)
				if pal > num {
					ans = pal
				}
				return pal > num
			}
		}

		if !limit {
			// 中间随便填
			left := n - i*2
			mask := 1<<9 - 1
			for j, c := range cnt {
				if c > 0 {
					left -= j - c
					mask ^= 1 << mp[j]
				}
			}
			if left < 0 {
				return false
			}

			// 枚举 mask 的子集，从大到小
			for sub, ok := mask, true; ok; ok = sub != mask {
				if (odd == 0 || sub&oddMask == 0) && subSum[sub] == left {
					ans = extend(pal, sub, odd)
					return true
				}
				sub = (sub - 1) & mask
			}
			return false
		}

		low := int(s[i] - '0')
		for v := low; v <= 9; v++ {
			if cnt[v]+2 > v || odd > 0 && v%2 > 0 && v != odd {
				continue
			}
			newOdd := odd
			if v%2 > 0 {
				newOdd = v
			}
			cnt[v] += 2
			if dfs(i+1, pal*10+v, newOdd, limit && v == low) {
				return true
			}
			cnt[v] -= 2
		}
		return false
	}

	if dfs(0, 0, 0, true) {
		return int64(ans)
	}

	// 没找到就取长为 n+1 的最小回文数
	for mask := 1<<9 - 1; ; mask-- {
		if subSum[mask] == n+1 {
			return int64(extend(0, mask, 0))
		}
	}
}

//

var specialNumbers []int

func init() {
	const oddMask = 0x155
	for mask := 1; mask < 512; mask++ {
		t := mask & oddMask
		if t&(t-1) > 0 { // 至少有两个奇数
			continue
		}

		perm := []int{}
		size := 0
		odd := 0
		for s := uint(mask); s > 0; s &= s - 1 {
			x := bits.TrailingZeros(s) + 1
			size += x
			for range x / 2 {
				perm = append(perm, x)
			}
			if x%2 > 0 {
				odd = x
			}
		}
		if size > 16 {
			continue
		}

		permutations(len(perm), len(perm), func(idx []int) bool {
			pal := 0
			for _, i := range idx {
				pal = pal*10 + perm[i]
			}
			v := pal
			if odd > 0 {
				pal = pal*10 + odd
			}
			// 反转 pal 的左半，拼在 pal 后面
			for ; v > 0; v /= 10 {
				pal = pal*10 + v%10
			}
			specialNumbers = append(specialNumbers, pal)
			return false
		})
	}
	slices.Sort(specialNumbers)
	specialNumbers = slices.Compact(specialNumbers)
}

func specialPalindrome2(n int64) int64 {
	i := sort.SearchInts(specialNumbers, int(n+1))
	return int64(specialNumbers[i])
}

func permutations(n, r int, do func(ids []int) (Break bool)) {
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	if do(ids[:r]) {
		return
	}
	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}
	for {
		i := r - 1
		for ; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				tmp := ids[i]
				copy(ids[i:], ids[i+1:])
				ids[n-1] = tmp
				cycles[i] = n - i
			} else {
				j := cycles[i]
				ids[i], ids[n-j] = ids[n-j], ids[i]
				if do(ids[:r]) {
					return
				}
				break
			}
		}
		if i == -1 {
			return
		}
	}
}

//

var size [512]int

func init() {
	const oddMask = 0x155
	for mask := 1; mask < 512; mask++ {
		t := mask & oddMask
		if t&(t-1) > 0 { // 至少有两个奇数
			continue
		}
		for s := uint(mask); s > 0; s &= s - 1 {
			size[mask] += bits.TrailingZeros(s) + 1
		}
	}
}

func specialPalindrome1(Num int64) int64 {
	num := int(Num)
	targetSize := len(strconv.Itoa(num))
	ans := math.MaxInt
	for mask := 1; mask < 512; mask++ {
		sz := size[mask]
		if sz != targetSize && sz != targetSize+1 {
			continue
		}

		perm := make([]int, 0, sz/2)
		odd := 0
		for s := uint(mask); s > 0; s &= s - 1 {
			x := bits.TrailingZeros(s) + 1
			for range x / 2 {
				perm = append(perm, x)
			}
			if x%2 > 0 {
				odd = x
			}
		}

		permutations(len(perm), len(perm), func(idx []int) (Break bool) {
			x := 0
			for _, i := range idx {
				x = x*10 + perm[i]
			}
			tmp := x
			if odd > 0 {
				x = x*10 + odd
			}
			// 反转 x 的左半，拼在 x 后面
			for v := tmp; v > 0; v /= 10 {
				x = x*10 + v%10
			}
			if x >= ans { // 最优性剪枝，不再继续枚举
				return true
			}
			if x > num { // 满足要求
				ans = x
				return true
			}
			return false
		})
	}
	return int64(ans)
}
