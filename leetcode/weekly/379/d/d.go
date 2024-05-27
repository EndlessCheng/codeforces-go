package main

import (
	"math/bits"
	"sort"
)

// https://space.bilibili.com/206214
func maxPartitionsAfterOperations(s string, k int) int {
	if k == 26 {
		return 1
	}

	seg, mask, size := 1, 0, 0
	update := func(i int) {
		bit := 1 << (s[i] - 'a')
		if mask&bit > 0 {
			return
		}
		size++
		if size > k {
			seg++ // s[i] 在新的一段中
			mask = bit
			size = 1
		} else {
			mask |= bit
		}
	}

	n := len(s)
	type pair struct{ seg, mask int }
	suf := make([]pair, n+1)
	for i := n - 1; i >= 0; i-- {
		update(i)
		suf[i] = pair{seg, mask}
	}

	ans := seg // 不修改任何字母
	seg, mask, size = 1, 0, 0
	for i := range s {
		p := suf[i+1]
		res := seg + p.seg // 情况 3
		unionSize := bits.OnesCount(uint(mask | p.mask))
		if unionSize < k {
			res-- // 情况 1
		} else if unionSize < 26 && size == k && bits.OnesCount(uint(p.mask)) == k {
			res++ // 情况 2
		}
		ans = max(ans, res)
		update(i)
	}
	return ans
}

func maxPartitionsAfterOperations2(s string, k int) int {
	n := len(s)
	type args struct {
		i, mask int
		changed bool
	}
	memo := map[args]int{}
	var dfs func(int, int, bool) int
	dfs = func(i, mask int, changed bool) (res int) {
		if i == n {
			return 1
		}

		a := args{i, mask, changed}
		if v, ok := memo[a]; ok { // 之前计算过
			return v
		}

		// 不改 s[i]
		bit := 1 << (s[i] - 'a')
		newMask := mask | bit
		if bits.OnesCount(uint(newMask)) > k {
			// 分割出一个子串，这个子串的最后一个字母在 i-1
			// s[i] 作为下一段的第一个字母，也就是 bit 作为下一段的 mask 的初始值
			res = dfs(i+1, bit, changed) + 1
		} else { // 不分割
			res = dfs(i+1, newMask, changed)
		}

		if !changed {
			// 枚举把 s[i] 改成 a,b,c,...,z
			for j := 0; j < 26; j++ {
				newMask := mask | 1<<j
				if bits.OnesCount(uint(newMask)) > k {
					// 分割出一个子串，这个子串的最后一个字母在 i-1
					// j 作为下一段的第一个字母，也就是 1<<j 作为下一段的 mask 的初始值
					res = max(res, dfs(i+1, 1<<j, true)+1)
				} else { // 不分割
					res = max(res, dfs(i+1, newMask, true))
				}
			}
		}

		memo[a] = res // 记忆化
		return res
	}
	return dfs(0, 0, false)
}

func maxPartitionsAfterOperationsWA(s string, k int) int {
	if k == 26 {
		return 1
	}

	n := len(s)
	sum := make([][26]int, n+1)
	for i, b := range s {
		sum[i+1] = sum[i]
		sum[i+1][b-'a']++
	}

	// 左闭右开 [l, r)
	count := func(l, r int) []int {
		if l > r {
			panic(-1)
		}
		res := sum[r]
		for i, sl := range sum[l][:] {
			res[i] -= sl
		}
		return res[:]
	}

	type pair struct{ seg, nxt int }
	suf := make([]pair, n+1)
	sufChange := make([][26]pair, n+1)
	for i := n - 1; i >= 0; i-- {
		r := sort.Search(n, func(r int) bool {
			r++
			if r <= i { // todo
				return false
			}
			cnt := 0
			for _, c := range count(i, r) {
				if c > 0 {
					cnt++
				}
			}
			return cnt > k
		})
		suf[i] = pair{suf[r].seg + 1, r}

		// 改 s[i] 为 ch
		for ch := 0; ch < 26; ch++ {
			r = sort.Search(n, func(r int) bool {
				r++
				if r <= i { // todo
					return false
				}
				vis := 1 << ch
				for p, c := range count(i+1, r) {
					if c > 0 {
						vis |= 1 << p
					}
				}
				return bits.OnesCount(uint(vis)) > k
			})
			sufChange[i][ch] = pair{suf[r].seg + 1, r}
		}
	}

	ans := suf[0].seg
	preSeg := 0
	for i := 0; i < n; {
		preSeg++
		maxR := suf[i].nxt
		for ch := 0; ch < 26; ch++ {
			for j := i; j < maxR; j++ {
				// 改 s[j] 为 ch

				r := sort.Search(n, func(r int) bool {
					r++
					if r <= j {
						return false
					}
					vis := 1 << ch
					for p, c := range count(i, j) {
						if c > 0 {
							vis |= 1 << p
						}
					}
					if j+1 < r {
						for p, c := range count(j+1, r) {
							if c > 0 {
								vis |= 1 << p
							}
						}
					}
					return bits.OnesCount(uint(vis)) > k
				})
				if r == j {
					ans = max(ans, preSeg+sufChange[r][ch].seg)
					break
				} else {
					ans = max(ans, preSeg+suf[r].seg)
				}
			}
		}
		i = maxR
	}
	return ans
}
