package main

// https://space.bilibili.com/206214
func sumOfPower(nums []int, k int) int {
	const mod = 1_000_000_007
	f := make([]int, k+1)
	f[0] = 1
	sum := 0
	for _, x := range nums {
		sum = min(sum+x, k)
		for j := sum; j >= 0; j-- {
			if j >= x {
				f[j] = (f[j]*2 + f[j-x]) % mod
			} else {
				f[j] = f[j] * 2 % mod
			}
		}
	}
	return f[k]
}

func sumOfPower2(nums []int, k int) (ans int) {
	const mod = 1_000_000_007
	n := len(nums)
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[0][0] = 1
	for i, x := range nums {
		for j := k; j >= x; j-- {
			for c := i + 1; c > 0; c-- {
				f[j][c] = (f[j][c] + f[j-x][c-1]) % mod
			}
		}
	}
	pow2 := 1
	for i := n; i > 0; i-- {
		ans = (ans + f[k][i]*pow2) % mod
		pow2 = pow2 * 2 % mod
	}
	return
}
