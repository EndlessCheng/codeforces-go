package main

import "math/bits"

// https://space.bilibili.com/206214
func clearStars(S string) string {
	s := []byte(S)
	stacks := [26][]int{}
	mask := 0
	for i, c := range s {
		if c != '*' {
			c -= 'a'
			stacks[c] = append(stacks[c], i)
			mask |= 1 << c
		} else {
			k := bits.TrailingZeros(uint(mask))
			st := stacks[k]
			s[st[len(st)-1]] = '*'
			stacks[k] = st[:len(st)-1]
			if len(stacks[k]) == 0 {
				mask ^= 1 << k
			}
		}
	}

	t := s[:0]
	for _, c := range s {
		if c != '*' {
			t = append(t, c)
		}
	}
	return string(t)
}

func clearStars2(S string) string {
	s := []byte(S)
	stacks := make([][]int, 26)
	for i, c := range s {
		if c != '*' {
			stacks[c-'a'] = append(stacks[c-'a'], i)
			continue
		}
		for j, st := range stacks {
			if m := len(st); m > 0 {
				s[st[m-1]] = '*'
				stacks[j] = st[:m-1]
				break
			}
		}
	}

	ans := s[:0]
	for _, c := range s {
		if c != '*' {
			ans = append(ans, c)
		}
	}
	return string(ans)
}
