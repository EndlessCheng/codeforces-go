package main

// github.com/EndlessCheng/codeforces-go
func sumOddLengthSubarrays2(a []int) (ans int) {
	for i := range a {
		s := 0
		for j := i; j < len(a); j++ {
			s += a[j]
			if (j-i)&1 == 0 {
				ans += s
			}
		}
	}
	return
}

// O(n) 做法，计算每个数的贡献：
// 对于一个数 a[i]，若其出现在一个长度为奇数的子数组中
// 则有两种情况：a[i] 左右均有奇数个数，或左右均有偶数个数
func sumOddLengthSubarrays(a []int) (ans int) {
	n := len(a)
	for i, v := range a {
		ans += ((i+1)/2*((n-i)/2) + (i/2+1)*((n-1-i)/2+1)) * v
	}
	return
}
