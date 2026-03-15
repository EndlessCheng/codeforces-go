由于在同一个数组内交换是免费的，我们可以把数组视作**无序集合**，里面的元素是什么顺序都可以。

于是问题转化成：

- 给你两个大小都为 $n$ 的集合 $A$ 和 $B$。每次操作，从两个集合中各选一个数，交换。最少交换多少次，可以让 $A=B$？

首先，每种元素的**总**出现次数必须是**偶数**，不然无法均分。

否则可以均分。比如元素 $x$ 在集合 $A$ 中出现 $8$ 次，在集合 $B$ 中出现 $2$ 次，一共有 $10$ 个 $x$。交换后，需要满足两个集合各有 $5$ 个 $x$，也就是把 $A$ 中的 $3$ 个 $x$ 与 $B$ 中的另外 $3$ 个数交换（交换的数是集合 $A$ 需要的数），交换 $3$ 次。

定义 $\textit{diff}[x]$ 表示元素 $x$ 在集合 $A$ 中的出现次数，减去 $x$ 在集合 $B$ 中的出现次数。

我们需要把 $\textit{diff}[x]$ 变成 $0$，这样两个集合中的元素 $x$ 的个数就相等了。

由于两个集合的大小都是 $n$，根据 $\textit{diff}$ 的定义，$\textit{diff}$ 的总和（出现次数之差的总和）是 $n-n=0$。换句话说，$\textit{diff}$ 中的正数之和等于负数之和的绝对值。**集合 $A$ 多出的数，恰好也是集合 $B$ 多出的数**。把这些数交换，即可让 $A=B$。由于交换一次，可以让 $\textit{diff}[x]$ 中的正数之和减少 $1-(-1) = 2$，所以交换次数等于 $\textit{diff}$ 中的正数之和除以 $2$。在上面的例子中，$\textit{diff}[x]=6$，需要交换 $\dfrac{6}{2} = 3$ 次。

[本题视频讲解](https://www.bilibili.com/video/BV111wTzQEbp/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minCost(self, nums1: List[int], nums2: List[int]) -> int:
        diff = Counter(nums1)
        diff.subtract(nums2)

        ans = 0
        for d in diff.values():
            if d % 2:
                return -1
            if d > 0:
                ans += d
        return ans // 2
```

```java [sol-Java]
class Solution {
    public int minCost(int[] nums1, int[] nums2) {
        Map<Integer, Integer> diff = new HashMap<>(); // 更快的写法见【Java 数组】
        for (int x : nums1) {
            diff.merge(x, 1, Integer::sum); // diff[x]++
        }
        for (int x : nums2) {
            diff.merge(x, -1, Integer::sum); // diff[x]--
        }

        int ans = 0;
        for (int d : diff.values()) {
            if (d % 2 != 0) {
                return -1;
            }
            if (d > 0) {
                ans += d;
            }
        }
        return ans / 2;
    }
}
```

```java [sol-Java 数组]
class Solution {
    public int minCost(int[] nums1, int[] nums2) {
        int mx = 0;
        for (int x : nums1) {
            mx = Math.max(mx, x);
        }
        for (int x : nums2) {
            mx = Math.max(mx, x);
        }

        int[] diff = new int[mx + 1];
        for (int x : nums1) {
            diff[x]++;
        }
        for (int x : nums2) {
            diff[x]--;
        }

        int ans = 0;
        for (int d : diff) {
            if (d % 2 != 0) {
                return -1;
            }
            if (d > 0) {
                ans += d;
            }
        }
        return ans / 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minCost(vector<int>& nums1, vector<int>& nums2) {
        unordered_map<int, int> diff;
        for (int x : nums1) {
            diff[x]++;
        }
        for (int x : nums2) {
            diff[x]--;
        }

        int ans = 0;
        for (auto& [_, d] : diff) {
            if (d % 2) {
                return -1;
            }
            if (d > 0) {
                ans += d;
            }
        }
        return ans / 2;
    }
};
```

```go [sol-Go]
func minCost(nums1, nums2 []int) (ans int) {
	diff := map[int]int{}
	for _, x := range nums1 {
		diff[x]++
	}
	for _, x := range nums2 {
		diff[x]--
	}

	for _, d := range diff {
		if d%2 != 0 {
			return -1
		}
		if d > 0 {
			ans += d
		}
	}
	return ans / 2
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}_1$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
