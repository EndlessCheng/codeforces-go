package main

import "sort"

// https://space.bilibili.com/206214
func equalFrequency(word string) bool {
	mCnt := map[rune]int{}
	for _, c := range word {
		mCnt[c]++
	}
	cnt := make([]int, 0, len(mCnt))
	for _, c := range mCnt {
		cnt = append(cnt, c)
	}
	sort.Ints(cnt)
	m := len(cnt)
	return m == 1 ||
		cnt[0] == 1 && isSame(cnt[1:]) ||
		cnt[m-1] == cnt[m-2]+1 && isSame(cnt[:m-1])
}

func isSame(a []int) bool {
	for _, x := range a[1:] {
		if x != a[0] {
			return false
		}
	}
	return true
}
