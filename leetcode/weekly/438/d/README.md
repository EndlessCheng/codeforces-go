**前置知识**：[【图解】曼哈顿距离转切比雪夫距离](https://leetcode.cn/problems/minimize-manhattan-distances/solutions/2716755/tu-jie-man-ha-dun-ju-chi-heng-deng-shi-b-op84/)。

看到最大化最小值，考虑二分答案，即二分距离的下界 $\textit{low}$。

把正方形顺时针旋转 $45$ 度，转成切比雪夫距离

$$
\max(|x_1'-x_2'|,|y_1'-y_2'|)
$$

这样可以利用正方形的性质：对于上半圈来说，纵坐标之差不会超过横坐标之差，所以**只需比较横坐标之差**！对于下半圈同理。

这样可以暴力枚举所选的第一个点，然后二分横坐标找下一个点，如下图。

![w438d-c.png](https://pic.leetcode.cn/1740283369-FGocTC-w438d-c.png)

> 注：本题保证 $k\ge 4$，所以答案不会超过 $\textit{side}$。这也保证了「找下半圈最右点」是有单调性的，是可以二分的。而 $k\le 3$ 时，答案可能会超过 $\textit{side}$，此时找下半圈的最右点时，不一定有单调性，无法二分。

在下半圈的循环过程中，需要保证下半圈最左边的点和上半圈第一个点相距 $\ge \textit{low}$，如果不满足就退出循环。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

其他语言直播结束后补充。

```go [sol-Go]
func maxDistance(side int, points [][]int, k int) int {
	type pair struct{ x, y int }
	var a, b []pair
	for _, p := range points {
		x, y := p[0], p[1]
		q := pair{x + y, y - x}
		if x == 0 && y >= 0 || y == side {
			a = append(a, q)
		} else {
			b = append(b, q)
		}
	}

	cmp := func(a, b pair) int { return a.x - b.x }
	slices.SortFunc(a, cmp)
	slices.SortFunc(b, cmp)
	if len(a) == 0 {
		// 保证 a 至少有一个点
		a, b = b, a
	}

	// 本题保证 k >= 4，所以最远距离不会超过 side
	ans := sort.Search(side, func(low int) bool {
		low++
		for _, p := range a {
			// 绕一圈
			// 上半圈（从左往右找）
			firstX, firstY := p.x, p.y // 第一个点
			left := k - 1
			curX, lastY := p.x, p.y
			for left > 0 {
				j := sort.Search(len(a), func(i int) bool { return a[i].x >= curX+low })
				if j == len(a) {
					break
				}
				curX = a[j].x
				lastY = a[j].y
				left--
			}
			if left == 0 {
				return false
			}

			// 下半圈最右边的点
			j := sort.Search(len(b), func(i int) bool {
				return max(abs(b[i].x-curX), abs(b[i].y-lastY)) < low
			}) - 1
			// 不能和第一个点离得太近
			if j < 0 || max(abs(b[j].x-firstX), abs(b[j].y-firstY)) < low {
				continue
			}

			// 下半圈（从右往左找）
			left--
			curX = b[j].x
			for left > 0 {
				j := sort.Search(len(b), func(i int) bool { return b[i].x > curX-low }) - 1
				// 不能和第一个点离得太近
				if j < 0 || max(abs(b[j].x-firstX), abs(b[j].y-firstY)) < low {
					break
				}
				curX = b[j].x
				left--
			}
			if left == 0 {
				return false
			}
		}
		return true
	})
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk \log \textit{side}\log n)$，其中 $n$ 是 $\textit{points}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. 【本题相关】[二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
