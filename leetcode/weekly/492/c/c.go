package main

import "slices"

// https://space.bilibili.com/206214
func minOperations(s string) int {
	t := []byte(s)
	// s 已经是升序
	if slices.IsSorted(t) {
		return 0
	}

	n := len(t)
	// 长为 2，无法排序
	if n == 2 {
		return -1
	}

	mn := slices.Min(t)
	mx := slices.Max(t)
	// 如果 s[0] 是最小值，排序 [1,n-1] 即可
	// 如果 s[n-1] 是最大值，排序 [0,n-2] 即可
	if t[0] == mn || t[n-1] == mx {
		return 1
	}

	// 如果 [1,n-2] 中有最小值，那么先排序 [0,n-2]，把最小值排在最前面，然后排序 [1,n-1] 即可
	// 如果 [1,n-2] 中有最大值，那么先排序 [1,n-1]，把最大值排在最后面，然后排序 [0,n-2] 即可
	for _, ch := range t[1 : n-1] {
		if ch == mn || ch == mx {
			return 2
		}
	}

	// 现在只剩下一种情况：s[0] 是最大值，s[n-1] 是最小值，且 [1,n-2] 不含最小值和最大值
	// 先排序 [0,n-2]，把最大值排到 n-2
	// 然后排序 [1,n-1]，把最大值排在最后面，且最小值排在 1
	// 最后排序 [0,n-2]，把最小值排在最前面
	return 3
}
