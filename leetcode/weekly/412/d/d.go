package main

import (
	"slices"
	"strconv"
)

// https://space.bilibili.com/206214
var pow10 = [...]int{1, 10, 100, 1000, 10000, 100000, 1000000}

func countPairs(nums []int) int {
	slices.Sort(nums)
	ans := 0
	cnt := make(map[int]int)
	for _, x := range nums {
		st := map[int]struct{}{x: {}} // 不交换
		a := []int{}
		for v := x; v > 0; v /= 10 {
			a = append(a, v%10)
		}
		m := len(a)
		for i := 0; i < m; i++ {
			for j := i + 1; j < m; j++ {
				if a[i] == a[j] { // 小优化
					continue
				}
				y := x + (a[j]-a[i])*(pow10[i]-pow10[j])
				st[y] = struct{}{} // 交换一次
				a[i], a[j] = a[j], a[i]
				for p := i + 1; p < m; p++ {
					for q := p + 1; q < m; q++ {
						st[y+(a[q]-a[p])*(pow10[p]-pow10[q])] = struct{}{} // 交换两次
					}
				}
				a[i], a[j] = a[j], a[i]
			}
		}
		for x := range st {
			ans += cnt[x]
		}
		cnt[x]++
	}
	return ans
}

//

func countPairs2(nums []int) (ans int) {
	slices.Sort(nums)
	cnt := map[int]int{}
	for _, x := range nums {
		set := map[int]struct{}{x: {}} // 不交换
		s := []byte(strconv.Itoa(x))
		m := len(s)
		for i := range s {
			for j := i + 1; j < m; j++ {
				s[i], s[j] = s[j], s[i]
				set[atoi2(s)] = struct{}{} // 交换一次
				for p := i + 1; p < m; p++ {
					for q := p + 1; q < m; q++ {
						s[p], s[q] = s[q], s[p]
						set[atoi2(s)] = struct{}{} // 交换两次
						s[p], s[q] = s[q], s[p]
					}
				}
				s[i], s[j] = s[j], s[i]
			}
		}
		for x := range set {
			ans += cnt[x]
		}
		cnt[x]++
	}
	return
}

// 手动转 int 快一些
func atoi2(s []byte) (res int) {
	for _, b := range s {
		res = res*10 + int(b&15)
	}
	return
}
