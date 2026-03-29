package main

// https://space.bilibili.com/206214
const mx = 5001
const maxDigitSum = 31 // 4999 的数位和最大
var sumToNums [maxDigitSum + 1][]int

func init() {
	digSum := [mx]int{}
	for x := range digSum {
		// 去掉 x 的个位，问题变成 x/10 的数位和，即 digSum[x/10]
		digSum[x] = digSum[x/10] + x%10
		sumToNums[digSum[x]] = append(sumToNums[digSum[x]], x)
	}
}

func countArrays(digitSum []int) (ans int) {
	const mod = 1_000_000_007
	f := [mx]int{1} // f[x] 表示以 x 结尾的有效数组的个数
	pre := 0

	for _, cur := range digitSum {
		if cur > maxDigitSum {
			return 0
		}
		a := sumToNums[pre]
		j, m := 0, len(a)
		sum := 0
		for _, x := range sumToNums[cur] {
			// 有效数组的前一个数只要 <= x 就行
			for ; j < m && a[j] <= x; j++ {
				sum += f[a[j]]
			}
			// sum 现在就是以 x 结尾的有效数组的个数
			f[x] = sum % mod
		}
		pre = cur // 记录上一个数位和
	}

	for _, x := range sumToNums[pre] {
		ans += f[x]
	}
	return ans % mod
}

//const mx = 5001
//const maxDigitSum = 31 // 4999 的数位和最大
//var digSum [mx]int
//
//func init() {
//	// 预处理数位和
//	for x := range digSum {
//		// 去掉 x 的个位，问题变成 x/10 的数位和，即 digitSum[x/10]
//		digSum[x] = digSum[x/10] + x%10
//	}
//}
//
//func countArrays(digitSum []int) int {
//	const mod = 1_000_000_007
//	sum := [mx]int{}
//	for i := range sum {
//		sum[i] = 1
//	}
//	for _, ds := range digitSum {
//		if ds > maxDigitSum {
//			return 0
//		}
//		for x := range mx {
//			// 如果 digSum[x] != ds，那么 f[x] = 0，否则 f[x] = sum[i]
//			// 把 f[x] 的值填到 sum[x] 中，那么只需要把 digSum[x] != ds 的 sum[x] 置为 0
//			if digSum[x] != ds {
//				sum[x] = 0
//			}
//			if x > 0 {
//				sum[x] = (sum[x] + sum[x-1]) % mod
//			}
//		}
//	}
//	return sum[mx-1]
//}
