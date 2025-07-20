package main

// https://space.bilibili.com/206214
const mx = 100_000

var np = [mx]bool{true, true} // 0 和 1 不是质数

func init() {
	for i := 2; i*i < mx; i++ {
		if !np[i] {
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func splitArray(nums []int) (ans int64) {
	for i, x := range nums {
		if np[i] {
			ans -= int64(x)
		} else {
			ans += int64(x)
		}
	}
	if ans < 0 {
		ans = -ans
	}
	return
}
