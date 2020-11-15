package main

// github.com/EndlessCheng/codeforces-go
type CombinationIterator struct{}

var (
	s      string
	k      int
	idList [][]int
)

func combinations(n, r int, do func(ids []int)) {
	ids := make([]int, r)
	for i := range ids {
		ids[i] = i
	}
	do(ids)
	for {
		i := r - 1
		for ; i >= 0; i-- {
			if ids[i] != i+n-r {
				break
			}
		}
		if i == -1 {
			return
		}
		ids[i]++
		for j := i + 1; j < r; j++ {
			ids[j] = ids[j-1] + 1
		}
		do(ids)
		// 可以加一个提前退出的逻辑（虽然最多 6435 个）
		if len(idList) > 1e4 {
			return
		}
	}
}

func Constructor(S string, K int) (_ CombinationIterator) {
	s, k = S, K
	idList = nil
	combinations(len(s), k, func(ids []int) { idList = append(idList, append([]int(nil), ids...)) })
	return
}

func (CombinationIterator) Next() string {
	ids := idList[0]
	idList = idList[1:]
	ans := make([]byte, k)
	for i, id := range ids {
		ans[i] = s[id]
	}
	return string(ans)
}

func (CombinationIterator) HasNext() bool {
	return len(idList) > 0
}
