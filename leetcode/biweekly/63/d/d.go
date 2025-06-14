package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func kthSmallestProduct(a, b []int, K int64) int64 {
	n, m, k := len(a), len(b), int(K)
	i0 := sort.SearchInts(a, 0)
	j0 := sort.SearchInts(b, 0)

	corners := []int{a[0] * b[0], a[0] * b[m-1], a[n-1] * b[0], a[n-1] * b[m-1]}
	left := slices.Min(corners)
	right := slices.Max(corners)
	ans := left + sort.Search(right-left, func(mx int) bool {
		mx += left
		cnt := 0

		if mx < 0 {
			// 右上区域
			i, j := 0, j0
			for i < i0 && j < m {
				if a[i]*b[j] > mx {
					j++
				} else {
					cnt += m - j
					i++
				}
			}

			// 左下区域
			i, j = i0, 0
			for i < n && j < j0 {
				if a[i]*b[j] > mx {
					i++
				} else {
					cnt += n - i
					j++
				}
			}
		} else {
			// 右上区域和左下区域的所有数都 <= 0 <= mx
			cnt = i0*(m-j0) + (n-i0)*j0

			// 左上区域
			i, j := 0, j0-1
			for i < i0 && j >= 0 {
				if a[i]*b[j] > mx {
					i++
				} else {
					cnt += i0 - i
					j--
				}
			}

			// 右下区域
			i, j = i0, m-1
			for i < n && j >= j0 {
				if a[i]*b[j] > mx {
					j--
				} else {
					cnt += j - j0 + 1
					i++
				}
			}
		}

		return cnt >= k
	})
	return int64(ans)
}
