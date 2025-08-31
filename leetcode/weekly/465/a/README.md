**题意**：对于 $\textit{order}$ 中的数字 $x$，如果 $x$ 在 $\textit{friends}$ 中，那么把 $x$ 添加到答案中。

把 $\textit{friends}$ 转成哈希集合（或者布尔数组），这样我们可以快速判断一个数是否在 $\textit{friends}$ 中。

然后遍历 $\textit{order}$ 中的数字 $x$，如果 $x$ 是朋友，那么把 $x$ 添加到答案中。

```py [sol-Python3]
class Solution:
    def recoverOrder(self, order: List[int], friends: List[int]) -> List[int]:
        st = set(friends)
        return [x for x in order if x in st]
```

```py [sol-Python3 数组]
class Solution:
    def recoverOrder(self, order: List[int], friends: List[int]) -> List[int]:
        n = len(order)
        is_friend = [False] * (n + 1)
        for x in friends:
            is_friend[x] = True

        return [x for x in order if is_friend[x]]
```

```java [sol-Java]
class Solution {
    public int[] recoverOrder(int[] order, int[] friends) {
        int n = order.length;
        boolean[] isFriend = new boolean[n + 1];
        for (int x : friends) {
            isFriend[x] = true;
        }

        int[] ans = new int[friends.length];
        int idx = 0;
        for (int x : order) {
            if (isFriend[x]) {
                ans[idx++] = x;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> recoverOrder(vector<int>& order, vector<int>& friends) {
        int n = order.size();
        vector<int8_t> is_friend(n + 1);
        for (int x : friends) {
            is_friend[x] = true;
        }

        vector<int> ans;
        ans.reserve(friends.size()); // 预分配空间
        for (int x : order) {
            if (is_friend[x]) {
                ans.push_back(x);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func recoverOrder(order, friends []int) []int {
	n := len(order)
	isFriend := make([]bool, n+1)
	for _, x := range friends {
		isFriend[x] = true
	}

	ans := make([]int, 0, len(friends)) // 预分配空间
	for _, x := range order {
		if isFriend[x] {
			ans = append(ans, x)
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{order}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(m)$，其中 $m$ 是 $\textit{friends}$ 的长度。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
