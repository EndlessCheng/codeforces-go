**题意**：把 $0$ 视作 $-1$，找一个最短前缀，其元素和大于剩余元素和。

设 $\textit{possible}$ 的元素和为 $s$（把 $0$ 视作 $-1$）。

枚举 $x=\textit{possible}[i]$，同时计算前缀和 $\textit{pre}$，那么剩余元素和为 $s - \textit{pre}$

如果

$$
\textit{pre} > s - \textit{pre}
$$

即

$$
\textit{pre}\cdot 2 > s
$$

就返回 $i+1$，即前缀长度。

代码实现时，计算 $\textit{pre}$ 可以把 $1$ 视作 $2$，把 $0$ 视作 $-2$，这样无需计算乘 $2$。

附：[视频讲解](https://www.bilibili.com/video/BV19t421g7Pd/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minimumLevels(self, possible: List[int]) -> int:
        # cnt1 - cnt0 = cnt1 - (n - cnt1) = cnt1 * 2 - n
        s = sum(possible) * 2 - len(possible)
        pre = 0
        for i, x in enumerate(possible[:-1]):
            pre += 2 if x else -2
            if pre > s:
                return i + 1
        return -1
```

```java [sol-Java]
class Solution {
    public int minimumLevels(int[] possible) {
        // cnt1 - cnt0 = cnt1 - (n - cnt1) = cnt1 * 2 - n
        int n = possible.length;
        int s = 0;
        for (int x : possible) {
            s += x;
        }
        s = s * 2 - n;
        int pre = 0;
        for (int i = 0; i < n - 1; i++) {
            pre += possible[i] == 1 ? 2 : -2;
            if (pre > s) {
                return i + 1;
            }
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumLevels(vector<int> &possible) {
        // cnt1 - cnt0 = cnt1 - (n - cnt1) = cnt1 * 2 - n
        int n = possible.size();
        int s = accumulate(possible.begin(), possible.end(), 0) * 2 - n;
        int pre = 0;
        for (int i = 0; i < n - 1; i++) {
            pre += possible[i] ? 2 : -2;
            if (pre > s) {
                return i + 1;
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
func minimumLevels(possible []int) int {
	// cnt1 - cnt0 = cnt1 - (n - cnt1) = cnt1 * 2 - n
	n := len(possible)
	s := 0
	for _, x := range possible {
		s += x
	}
	s = s*2 - n
	pre := 0
	for i, x := range possible[:n-1] {
		pre += x*4 - 2
		if pre > s {
			return i + 1
		}
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{possible}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。Python 忽略切片空间。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
