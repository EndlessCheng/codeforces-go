每次操作会把 $\textit{nums}$ 中的所有等于 $x$ 的数都改成对应的 $\textit{target}[i]$。请注意，$\textit{nums}$ 中的不等于 $x$ 的数保持不变。

例如 $\textit{nums}=[1,2,1,2]$，$\texttt{target}=[3,4,5,6]$。选择 $x=1$，操作后 $\textit{nums}=[3,2,5,2]$，其中不等于 $1$ 的数保持不变，需要继续操作。

所以答案的**理论最小值**，等于满足 $\textit{nums}[i]\ne \textit{target}[i]$ 的不同 $\textit{nums}[i]$ 的个数。这些不同元素都至少要操作一次。

可以达到理论最小值吗？

可以。把满足 $\textit{nums}[i]\ne \textit{target}[i]$ 的 $\textit{nums}[i]$ 找出来，相同元素只保留其中一个。这些元素可以按照**任意顺序**操作。一旦 $\textit{nums}[i]$ 改成了 $\textit{target}[i]$，那么后续操作即使要改 $\textit{nums}[i]$，那也只是把 $\textit{target}[i]$ 改成了 $\textit{target}[i]$，保持不变。

既然操作后的数不会再变，那么操作等价于：

- 选择一个 $x$，删除 $\textit{nums}$ 中的所有等于 $x$ 的数。这些数操作后都等于目标值，不用再管。

这样可以立刻看出，答案就是满足 $\textit{nums}[i]\ne \textit{target}[i]$ 的不同 $\textit{nums}[i]$ 的个数。

[本题视频讲解](https://www.bilibili.com/video/BV1MVkxBZE4D/?t=3m44s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: List[int], target: List[int]) -> int:
        return len({x for x, t in zip(nums, target) if x != t})
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums, int[] target) {
        HashSet<Integer> set = new HashSet<>(); // 更快的写法见【Java 数组】
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            if (x != target[i]) {
                set.add(x);
            }
        }
        return set.size();
    }
}
```

```java [sol-Java 数组]
class Solution {
    public int minOperations(int[] nums, int[] target) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        boolean[] vis = new boolean[mx + 1];
        int ans = 0;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            if (!vis[x] && x != target[i]) {
                vis[x] = true;
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int>& nums, vector<int>& target) {
        unordered_set<int> st;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            if (x != target[i]) {
                st.insert(x);
            }
        }
        return st.size();
    }
};
```

```go [sol-Go]
func minOperations(nums, target []int) int {
	set := map[int]struct{}{}
	for i, x := range nums {
		if x != target[i] {
			set[x] = struct{}{}
		}
	}
	return len(set)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面贪心与思维题单的「**§5.2 脑筋急转弯**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
