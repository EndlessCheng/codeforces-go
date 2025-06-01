package main

import "math/big"

// https://space.bilibili.com/206214
func checkEqualPartitions2(nums []int, target int64) bool {
	tar := int(target)
	for s := 1; s < 1<<(len(nums)-1); s++ {
		mul1, mul2 := 1, 1
		for i, x := range nums {
			if s>>i&1 > 0 { // i 在集合 s 中
				mul1 = min(mul1*x, tar+1) // 与 tar+1 取 min，防止溢出
			} else { // i 在 s 的补集中
				mul2 = min(mul2*x, tar+1)
			}
		}
		if mul1 == tar && mul2 == tar {
			return true
		}
	}
	return false
}

func checkEqualPartitions1(nums []int, target int64) bool {
	tar := int(target)
	var dfs func(int, int, int) bool
	dfs = func(i, mul1, mul2 int) bool {
		if mul1 > tar || mul2 > tar {
			return false
		}
		if i == len(nums) {
			return mul1 == tar && mul2 == tar
		}
		return dfs(i+1, mul1*nums[i], mul2) || dfs(i+1, mul1, mul2*nums[i])
	}
	return dfs(0, 1, 1)
}

//

func calc(nums []int, target int) map[[2]int]struct{} {
	set := map[[2]int]struct{}{}
	var dfs func(int, int, int)
	dfs = func(i, a, b int) {
		if a > target || b > target {
			return
		}
		if i == len(nums) {
			g := gcd(a, b)
			set[[2]int{a / g, b / g}] = struct{}{} // 最简分数
			return
		}
		dfs(i+1, a*nums[i], b)
		dfs(i+1, a, b*nums[i])
	}
	dfs(0, 1, 1)
	return set
}

func checkEqualPartitions(nums []int, target int64) bool {
	prodAll := big.NewInt(1)
	for _, x := range nums {
		prodAll.Mul(prodAll, big.NewInt(int64(x)))
	}
	square := big.NewInt(target)
	square.Mul(square, square)
	if prodAll.Cmp(square) != 0 {
		return false
	}

	m := len(nums) / 2
	set1 := calc(nums[:m], int(target))
	set2 := calc(nums[m:], int(target))

	for p := range set1 {
		if _, ok := set2[p]; ok {
			return true
		}
	}
	return false
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
