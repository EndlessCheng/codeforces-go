package main

import "slices"

// https://space.bilibili.com/206214
func canMakePalindromeQueries(s string, queries [][]int) []bool {
	// 分成左右两半，右半反转
	n := len(s) / 2
	t := []byte(s[n:])
	slices.Reverse(t)
	s = s[:n]

	// 预处理三种前缀和
	sumS := make([][26]int, n+1)
	for i, b := range s {
		sumS[i+1] = sumS[i]
		sumS[i+1][b-'a']++
	}

	sumT := make([][26]int, n+1)
	for i, b := range t {
		sumT[i+1] = sumT[i]
		sumT[i+1][b-'a']++
	}

	sumNe := make([]int, n+1)
	for i := range s {
		sumNe[i+1] = sumNe[i]
		if s[i] != t[i] {
			sumNe[i+1]++
		}
	}

	// 计算子串中各个字符的出现次数，闭区间 [l,r]
	count := func(sum [][26]int, l, r int) []int {
		res := sum[r+1]
		for i, s := range sum[l][:] {
			res[i] -= s
		}
		return res[:]
	}

	subtract := func(s1, s2 []int) []int {
		for i, s := range s2 {
			s1[i] -= s
			if s1[i] < 0 {
				return nil
			}
		}
		return s1
	}

	check := func(l1, r1, l2, r2 int, sumS, sumT [][26]int) bool {
		if sumNe[l1] > 0 || // [0,l1-1] 有 s[i] != t[i]
			sumNe[n]-sumNe[max(r1, r2)+1] > 0 { // [max(r1,r2)+1,n-1] 有 s[i] != t[i]
			return false
		}
		if r2 <= r1 { // 区间包含
			return slices.Equal(count(sumS, l1, r1), count(sumT, l1, r1))
		}
		if r1 < l2 { // 区间不相交
			return sumNe[l2]-sumNe[r1+1] == 0 && // [r1+1,l2-1] 没有 s[i] != t[i]
				slices.Equal(count(sumS, l1, r1), count(sumT, l1, r1)) &&
				slices.Equal(count(sumS, l2, r2), count(sumT, l2, r2))
		}
		// 区间相交但不包含
		s1 := subtract(count(sumS, l1, r1), count(sumT, l1, l2-1))
		s2 := subtract(count(sumT, l2, r2), count(sumS, r1+1, r2))
		return s1 != nil && s2 != nil && slices.Equal(s1, s2)
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		l1, r1, l2, r2 := q[0], q[1], n*2-1-q[3], n*2-1-q[2]
		if l1 <= l2 {
			ans[i] = check(l1, r1, l2, r2, sumS, sumT)
		} else {
			ans[i] = check(l2, r2, l1, r1, sumT, sumS)
		}
	}
	return ans
}
