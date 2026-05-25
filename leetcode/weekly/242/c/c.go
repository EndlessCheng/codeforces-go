package main

// github.com/EndlessCheng/codeforces-go
func canReach1(s string, minJump, maxJump int) bool {
	n := len(s)
	sum := make([]int, n+1) // f 的前缀和
	sum[1] = 1              // f[0] = true
	for j := 1; j < n; j++ {
		sum[j+1] = sum[j]
		if j >= minJump && s[j] == '0' && sum[j-minJump+1] > sum[max(j-maxJump, 0)] {
			sum[j+1]++ // f[j] = true
		}
	}
	return sum[n] > sum[n-1] // f[n-1] == true
}

func canReach2(s string, minJump, maxJump int) bool {
	n := len(s)
	diff := make([]int, n+1)
	// 一开始在起点 0，把 [0, 0] 加一
	diff[0] = 1
	diff[1] = -1

	sumD := 0
	for i, ch := range s {
		sumD += diff[i]
		// 现在 sumD 是下标 i 的标记次数
		if sumD > 0 && ch == '0' { // i 可以跳到
			// [i+minJump, i+maxJump] 可以跳到
			diff[min(i+minJump, n)]++
			diff[min(i+maxJump+1, n)]--
		}
	}
	return sumD > 0 && s[n-1] == '0' // n-1 可以跳到
}

func canReach(s string, minJump, maxJump int) bool {
	n := len(s)
	canReaches := make([]bool, n)
	canReaches[0] = true
	j := 1
	for i, ch := range s {
		if ch == '0' && canReaches[i] {
			// 注意 j 只会增大，不会减小，所以总体时间复杂度是 O(n)
			for j = max(j, i+minJump); j <= min(i+maxJump, n-1); j++ {
				canReaches[j] = true
			}
		}
	}
	return s[n-1] == '0' && canReaches[n-1]
}
