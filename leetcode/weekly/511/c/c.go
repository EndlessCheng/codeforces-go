package main

import "strings"

// https://space.bilibili.com/206214
func transformStr1(s string, strs []string) []bool {
	total0 := strings.Count(s, "0")
	ans := make([]bool, len(strs))

next:
	for idx, str := range strs {
		cnt0 := strings.Count(str, "0")
		cntQ := strings.Count(str, "?")
		// str 中的 '0' 的个数在闭区间 [cnt0, cnt0+cntQ] 中，total0 必须在这个范围内
		if total0 < cnt0 || total0 > cnt0+cntQ {
			continue
		}

		// 把前 total0-cnt0 个 '?' 改成 '0'
		t := []byte(str)
		for i, ch := range t {
			if cnt0 == total0 {
				break
			}
			if ch == '?' {
				t[i] = '0'
				cnt0++
			}
		}

		// 判断能否把 s 变成 t
		i, j := 0, 0
		for range total0 {
			// 找下一个 s[i] = '0'
			for s[i] != '0' {
				i++
			}

			// 找下一个 t[j] = '0'
			for t[j] != '0' {
				j++
			}

			// s 中的 '0' 无法右移，所以无法把 s 变成 t
			if i < j {
				continue next
			}

			i++
			j++
		}

		ans[idx] = true
	}

	return ans
}

func transformStr(s string, strs []string) []bool {
	total0 := strings.Count(s, "0")
	ans := make([]bool, len(strs))

next:
	for idx, t := range strs {
		cnt0 := strings.Count(t, "0")
		cntQ := strings.Count(t, "?")
		// str 中的 '0' 的个数在闭区间 [cnt0, cnt0+cntQ] 中，total0 必须在这个范围内
		if total0 < cnt0 || total0 > cnt0+cntQ {
			continue
		}

		// 判断能否把 s 变成 t
		i, j := 0, 0
		for range total0 {
			// 找下一个 s[i] = '0'
			for s[i] != '0' {
				i++
			}

			// 找下一个 t[j] = '0'
			for t[j] == '1' || t[j] == '?' && cnt0 == total0 {
				j++
			}

			// s 中的 '0' 无法右移，所以无法把 s 变成 t
			if i < j {
				continue next
			}

			if t[j] == '?' {
				cnt0++
			}

			i++
			j++
		}

		ans[idx] = true
	}

	return ans
}
