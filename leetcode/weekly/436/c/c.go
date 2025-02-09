package main

// https://space.bilibili.com/206214
func countSubstrings(s string) (ans int64) {
	f := [10][9]int{}
	for _, c := range s {
		d := int(c - '0')
		for m := 1; m < 10; m++ { // 枚举模数 m
			// 滚动数组计算 f
			nf := [9]int{}
			nf[d%m] = 1
			for rem, fv := range f[m][:m] { // 枚举模 m 的余数 rem
				nf[(rem*10+d)%m] += fv // 刷表法
			}
			f[m] = nf
		}
		// 以 s[i] 结尾的，模 s[i] 余数为 0 的子串个数
		ans += int64(f[d][0])
	}
	return
}
