package main

import "sort"

// 排序+遍历

// github.com/EndlessCheng/codeforces-go
func findOriginalArray(changed []int) (original []int) {
	sort.Ints(changed)
	cnt := map[int]int{}
	for _, v := range changed {
		if cnt[v] == 0 { // v 不是双倍后的元素
			cnt[v*2]++ // 标记一个双倍元素
			original = append(original, v)
		} else {
			cnt[v]-- // 清除一个标记
			if cnt[v] == 0 {
				delete(cnt, v)
			}
		}
	}
	// 只有当所有双倍标记都被清除掉时，才能说明 changed 是一个双倍数组
	if len(cnt) == 0 {
		return
	}
	return nil
}
