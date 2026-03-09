package main

// https://space.bilibili.com/206214
func smallestBalancedIndex1(nums []int) int {
	n := len(nums)
	sum := make([]int, n) // sum[i] 表示 [0,i-1] 之和
	for i, x := range nums[:n-1] {
		sum[i+1] = sum[i] + x
	}

	mul := 1 // [i+1,n-1] 之积
	for i := n - 1; i > 0; i-- {
		if sum[i] == mul { // [0,i-1] 之和等于 [i+1,n-1] 之积
			return i
		}
		// 如果 mul*nums[i] > sum[i-1]，那么继续向左遍历，mul 越来越大（或者不变），sum 越来越小，不可能找到答案
		// 为避免乘法溢出，改成等价的除法
		if mul > sum[i-1]/nums[i] {
			break
		}
		mul *= nums[i]
	}
	return -1
}

func smallestBalancedIndex2(nums []int) int {
	n := len(nums)
	sum := 0
	for _, x := range nums[:n-1] {
		sum += x
	}

	mul := 1
	for i := n - 1; i > 0; i-- {
		if sum == mul {
			return i
		}
		sum -= nums[i-1]
		if mul > sum/nums[i] {
			break
		}
		mul *= nums[i]
	}
	return -1
}

func smallestBalancedIndex(nums []int) int {
	sum, mul := 0, 1
	l, r := 0, len(nums)-1
	for l < r {
		if sum < mul {
			sum += nums[l]
			l++
		} else {
			if mul > 1e14/nums[r] {
				return -1
			}
			mul *= nums[r]
			r--
		}
	}
	if sum == mul {
		return l
	}
	return -1
}
