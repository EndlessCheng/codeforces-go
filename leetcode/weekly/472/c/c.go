package main

import "strings"

// https://space.bilibili.com/206214
func lexGreaterPermutation1(s, target string) string {
	left := make([]int, 26)
	for i, b := range s {
		left[b-'a']++
		left[target[i]-'a']--
	}
	ans := []byte(target)

next:
	for i := len(s) - 1; i >= 0; i-- {
		b := target[i] - 'a'
		left[b]++
		for _, c := range left {
			if c < 0 { // 前面不能全部一样
				continue next
			}
		}

		// target[i] 增大到 j
		for j := b + 1; j < 26; j++ {
			if left[j] == 0 {
				continue
			}

			left[j]--
			ans[i] = 'a' + j
			ans = ans[:i+1]

			for k, c := range left {
				ch := string('a' + byte(k))
				ans = append(ans, strings.Repeat(ch, c)...)
			}
			return string(ans)
		}
		// 增大失败，继续枚举
	}
	return ""
}

func lexGreaterPermutation(s, target string) string {
	left := make([]int, 26)
	for i, b := range s {
		left[b-'a']++
		left[target[i]-'a']--
	}

	neg, mx := 0, byte(0)
	for i, cnt := range left {
		if cnt < 0 {
			neg++ // 统计 left 中的负数个数
		} else if cnt > 0 {
			mx = max(mx, byte(i))
		}
	}

	ans := []byte(target)
	for i := len(s) - 1; i >= 0; i-- {
		b := target[i] - 'a'
		left[b]++ // 撤销消耗

		if left[b] == 0 {
			neg--
		} else if left[b] == 1 {
			mx = max(mx, b)
		}

		// left 不能有负数 or 没有大于 target[i] 的字母
		if neg > 0 || b >= mx {
			continue
		}

		// target[i] 增大到 j
		for j := b + 1; ; j++ {
			if left[j] == 0 {
				continue
			}

			left[j]--
			ans[i] = 'a' + byte(j)
			ans = ans[:i+1]

			for k, c := range left {
				ch := string('a' + byte(k))
				ans = append(ans, strings.Repeat(ch, c)...)
			}
			return string(ans)
		}
	}
	return ""
}
