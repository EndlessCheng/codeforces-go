package main

// https://space.bilibili.com/206214
func maxTotal(nums []int, s string) int64 {
	f0, f1 := 0, 0
	for i, x := range nums {
		if s[i] == '0' {
			f1 = f0 + x
		} else {
			f0 = max(f0+x, f1)
			f1 += x
		}
	}
	return int64(f0)
}

func maxTotal2(nums []int, s string) int64 {
	n := len(s)
	ans, mn := 0, 0
	for i, x := range nums {
		if s[i] == '1' {
			ans += x
			mn = min(mn, x)                // 维护这一段的最小值
			if i == n-1 || s[i+1] == '0' { // 遍历到了段的末尾
				ans -= mn // 段的最小值不选
			}
		} else if i < n-1 && s[i+1] == '1' { // 0111..11 段的开头
			ans += x
			mn = x
		}
	}
	return int64(ans)
}
