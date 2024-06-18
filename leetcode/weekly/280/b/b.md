正难则反，考虑最多保留多少个数**不变**。

根据题意，我们可以按奇偶下标分组，记作 $a_0$ 和 $a_1$，每组内的元素需要相等。

考虑每组内出现次数最多的元素，分别记作 $x$ 和 $y$。

- 如果 $x\ne y$，那么 $a_0$ 保留 $x$，$a_1$ 保留 $y$，其余元素需要修改。
- 如果 $x=y$，那么可以考虑每组内出现次数第二多的元素，分别记作 $x'$ 和 $y'$，那么可以 $a_0$ 保留 $x$，$a_1$ 保留 $y'$，或者 $a_0$ 保留 $x'$，$a_1$ 保留 $y$，取这两种情况可以保留的最大值。

代码实现时，若某个组内元素不足两个，可以用不在 $\textit{nums}$ 中的元素（比如 $0$）来填充。

```go
type pair struct{ num, cnt int }

// 计算出现次数最多的两个元素及其出现次数
func getMaxCnt2(cnt map[int]int) []pair {
	a := make([]pair, 0, max(len(cnt), 2))
	for num, c := range cnt {
		a = append(a, pair{num, c})
	}
	sort.Slice(a, func(i, j int) bool { return a[i].cnt > a[j].cnt })
	return a[:2] // 不足两个时，用 pair{0, 0} 填充
}

func minimumOperations(nums []int) int {
	cnt := [2]map[int]int{{}, {}}
	for i, num := range nums {
		cnt[i&1][num]++
	}
	a0 := getMaxCnt2(cnt[0])
	a1 := getMaxCnt2(cnt[1])
	if a0[0].num != a1[0].num {
		return len(nums) - a0[0].cnt - a1[0].cnt // 不相等时，保留出现次数最多的两个
	}
	return len(nums) - max(a0[0].cnt+a1[1].cnt, a0[1].cnt+a1[0].cnt) // 相等时，保留出现次数最多的和另一个出现次数次多的
}
```

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
