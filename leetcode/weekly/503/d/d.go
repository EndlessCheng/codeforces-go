package main

import "math"

// https://space.bilibili.com/206214
func numberOfPairs(nums1, nums2 []int, queries [][]int) (ans []int) {
	m, n := len(nums1), len(nums2)
	B := int(math.Sqrt(float64(m * n)))

	type block struct {
		l, r int         // 这一段对应 nums2 的子数组 [l, r)，注意是左闭右开区间
		cnt  map[int]int // 这一段每个元素的出现次数
		add  int         // 这一段整体要增加 add
	}
	blocks := make([]block, (n-1)/B+1)
	for i := 0; i < n; i += B {
		r := min(i+B, n)
		cnt := map[int]int{}
		for _, x := range nums2[i:r] {
			cnt[x]++
		}
		blocks[i/B] = block{i, r, cnt, 0}
	}

	for _, q := range queries {
		if q[0] == 1 {
			l, r, val := q[1], q[2]+1, q[3]
			for i := range blocks {
				b := &blocks[i]
				if b.r <= l {
					continue
				}
				if b.l >= r {
					break
				}
				// b 在 [l, r) 中
				if l <= b.l && b.r <= r {
					b.add += val
					continue
				}
				// b 的一部分在 [l, r) 中
				bl := max(b.l, l)
				br := min(b.r, r)
				// 暴力更新 nums2 的子数组 [bl, br) 的元素值及其出现次数
				for j := bl; j < br; j++ {
					b.cnt[nums2[j]]-- // 撤销旧的
					nums2[j] += val
					b.cnt[nums2[j]]++ // 添加新的
				}
			}
		} else {
			res := 0
			for _, b := range blocks {
				target := q[1] - b.add
				for _, x := range nums1 {
					res += b.cnt[target-x]
				}
			}
			ans = append(ans, res)
		}
	}
	return
}
