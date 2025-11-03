package main

import "strings"

// https://space.bilibili.com/206214
func lexGreaterPermutation1(s, target string) string {
	left := make([]int, 26)
	for i, b := range s {
		left[b-'a']++
		left[target[i]-'a']-- // 消耗 s 中的一个字母 target[i]
	}

next:
	for i := len(s) - 1; i >= 0; i-- {
		b := target[i] - 'a'
		left[b]++ // 撤销消耗
		for _, c := range left {
			if c < 0 { // [0,i-1] 无法做到全部一样
				continue next
			}
		}

		// 把 target[i] 增大到 j
		for j := b + 1; j < 26; j++ {
			if left[j] == 0 {
				continue
			}

			left[j]--
			ans := []byte(target[:i+1])
			ans[i] = 'a' + j

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

	for i := len(s) - 1; i >= 0; i-- {
		b := target[i] - 'a'
		left[b]++ // 撤销消耗

		if left[b] == 0 {
			neg--
		} else if left[b] == 1 {
			mx = max(mx, b)
		}

		// left 有负数 or 没有大于 target[i] 的字母
		if neg > 0 || b >= mx {
			continue
		}

		j := b + 1
		for left[j] == 0 {
			j++
		}

		// 把 target[i] 增大到 j
		left[j]--
		ans := []byte(target[:i+1])
		ans[i] = 'a' + byte(j)

		for k, c := range left {
			ch := string('a' + byte(k))
			ans = append(ans, strings.Repeat(ch, c)...)
		}
		return string(ans)
	}
	return ""
}
