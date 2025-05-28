package main

// https://space.bilibili.com/206214
func robotWithString(s string) string {
	n := len(s)
	sufMin := make([]byte, n+1)
	sufMin[n] = 'z'
	for i := n - 1; i >= 0; i-- {
		sufMin[i] = min(sufMin[i+1], s[i])
	}

	ans := make([]byte, 0, n) // 预分配空间
	st := sufMin[:0]
	for i, ch := range s {
		st = append(st, byte(ch))
		for len(st) > 0 && st[len(st)-1] <= sufMin[i+1] {
			ans = append(ans, st[len(st)-1])
			st = st[:len(st)-1]
		}
	}
	return string(ans)
}

func robotWithString2(s string) string {
	ans := make([]byte, 0, len(s))
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	min := byte(0) // 剩余字母的最小值
	st := []byte{}
	for _, c := range s {
		cnt[c-'a']--
		for min < 25 && cnt[min] == 0 {
			min++
		}
		st = append(st, byte(c))
		for len(st) > 0 && st[len(st)-1]-'a' <= min {
			ans = append(ans, st[len(st)-1])
			st = st[:len(st)-1]
		}
	}
	return string(ans)
}
