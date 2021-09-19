package main

/*

根据题意，答案长度不会超过 $\lceil\dfrac{n}{k}\rceil$，而 $n<8k$，故答案长度不超过 $7$。

我们可以统计出 $s$ 中个数不低于 $k$ 的字符，答案只能由这些字符注册，而这些字符的个数不会超过 $\lceil\dfrac{n}{k}\rceil$，由于 $n<8k$，故字符个数不超过 $7$。

这些然后暴力枚举这些字符的某个排列，然后贪心地匹配子序列，从中找到符合题目要求的子序列。

*/

// github.com/EndlessCheng/codeforces-go
func longestSubsequenceRepeatedK(s string, k int) (ans string) {
	n := len(s)
	right := [26]int{}
	for i := range right {
		right[i] = n
	}
	nxt := make([][26]int, n)
	cnt := [26]int{}
	for i := n - 1; i >= 0; i-- {
		nxt[i] = right
		right[s[i]-'a'] = i
		cnt[s[i]-'a']++
	}

	// 计算所有可能出现在 ans 中的字符，包括重复的
	// 倒着统计，这样下面计算排列时的第一个合法方案就是答案，从而提前退出
	a := []byte{}
	for i := 25; i >= 0; i-- {
		for c := cnt[i]; c >= k; c -= k {
			a = append(a, 'a'+byte(i))
		}
	}

	for m := len(a); m > 0 && ans == ""; m-- {
		permutations(len(a), m, func(ids []int) bool {
			t := make([]byte, m)
			for i, id := range ids {
				t[i] = a[id]
			}
			i, j := 0, 0
			if t[0] == s[0] {
				j = 1
			}
			for {
				i = nxt[i][t[j%m]-'a']
				if i == n {
					break
				}
				j++
			}
			if j >= k*len(t) {
				ans = string(t)
				return true
			}
			return false
		})
	}
	return
}

// 生成 n 选 r 的排列
func permutations(n, r int, do func(ids []int) bool) {
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	if do(ids[:r]) {
		return
	}
	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}
	for {
		i := r - 1
		for ; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				tmp := ids[i]
				copy(ids[i:], ids[i+1:])
				ids[n-1] = tmp
				cycles[i] = n - i
			} else {
				j := cycles[i]
				ids[i], ids[n-j] = ids[n-j], ids[i]
				if do(ids[:r]) {
					return
				}
				break
			}
		}
		if i == -1 {
			return
		}
	}
}
