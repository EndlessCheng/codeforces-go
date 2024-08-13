题目说，同一个包裹中的苹果可以分装到不同的箱子中。

那么按照容量从大到小选择箱子装苹果，直到所有苹果均装入箱子为止。

注意题目保证可以将包裹中的苹果重新分装到箱子中。

```py [sol-Python3]
class Solution:
    def minimumBoxes(self, apple: List[int], capacity: List[int]) -> int:
        s = sum(apple)
        capacity.sort(reverse=True)
        for i, x in enumerate(capacity, 1):
            s -= x
            if s <= 0:  # 所有苹果都装入了箱子
                return i
```

```java [sol-Java]
class Solution {
    public int minimumBoxes(int[] apple, int[] capacity) {
        int s = 0;
        for (int x : apple) {
            s += x;
        }
        Arrays.sort(capacity);
        int m = capacity.length;
        int i = m - 1;
        while (s > 0) {
            s -= capacity[i--];
        }
        return m - 1 - i;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumBoxes(vector<int>& apple, vector<int>& capacity) {
        int s = reduce(apple.begin(), apple.end());
        ranges::sort(capacity, greater());
        int i = 0;
        while (s > 0) {
            s -= capacity[i++];
        }
        return i;
    }
};
```

```go [sol-Go]
func minimumBoxes(apple, capacity []int) int {
	s := 0
	for _, x := range apple {
		s += x
	}
	slices.SortFunc(capacity, func(a, b int) int { return b - a })
	for i, c := range capacity {
		s -= c
		if s <= 0 { // 所有苹果都装入了箱子
			return i + 1 // 0 到 i 有 i+1 个箱子
		}
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m\log m)$，其中 $n$ 为 $\textit{apple}$ 的长度，$m$ 为 $\textit{capacity}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
