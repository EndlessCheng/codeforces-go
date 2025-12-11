package main

// https://space.bilibili.com/206214
func minDeletionSize1(strs []string) (ans int) {
	n, m := len(strs), len(strs[0])
	a := make([]string, n) // 最终得到的字符串数组
next:
	for j := range m {
		for i := range n - 1 {
			if a[i]+string(strs[i][j]) > a[i+1]+string(strs[i+1][j]) {
				// j 列不是升序，必须删
				ans++
				continue next
			}
		}
		// j 列是升序，不删更好
		for i, s := range strs {
			a[i] += string(s[j])
		}
	}
	return
}

func minDeletionSize(strs []string) (ans int) {
	n, m := len(strs), len(strs[0])
	checkList := make([]int, n-1)
	for i := range checkList {
		checkList[i] = i
	}

next:
	for j := range m {
		for _, i := range checkList {
			if strs[i][j] > strs[i+1][j] {
				// j 列不是升序，必须删
				ans++
				continue next
			}
		}
		// j 列是升序，不删更好
		newCheckList := checkList[:0]
		for _, i := range checkList {
			if strs[i][j] == strs[i+1][j] {
				// 相邻字母相等，下一列 i 和 i+1 需要继续比大小
				newCheckList = append(newCheckList, i)
			}
		}
		checkList = newCheckList
	}
	return
}
