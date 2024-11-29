本题 [视频讲解](https://www.bilibili.com/video/BV1pW4y1r7xs) 已出炉，欢迎点赞三连~

---

设 $s_1 = \sum\limits_{i}\textit{nums}_1[i]$。

交换 $[\textit{left},\textit{right}]$ 范围内的元素后，对于 $\textit{nums}'_1$ 有

$$
\sum\limits_{i}\textit{nums}'_1[i] = s_1 - (\textit{nums}_1[\textit{left}] + \cdots + \textit{nums}_1[\textit{right}]) + (\textit{nums}_2[\textit{left}] + \cdots + \textit{nums}_2[\textit{right}])
$$

合并相同下标，等号右侧变形为

$$
s_1 + (\textit{nums}_2[\textit{left}]-\textit{nums}_1[\textit{left}]) + \cdots + (\textit{nums}_2[\textit{right}]-\textit{nums}_1[\textit{right}])
$$

设 $\textit{diff}[i] = \textit{nums}_2[i]-\textit{nums}_1[i]$，上式变为

$$
s_1 + \textit{diff}[\textit{left}] + \cdots + \textit{diff}[\textit{right}]
$$

为了最大化上式，我们需要最大化 $\textit{diff}$ 数组的 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)，[我的题解](https://leetcode.cn/problems/maximum-subarray/solutions/2533977/qian-zhui-he-zuo-fa-ben-zhi-shi-mai-mai-abu71/)。注意子数组可以为空，所以初始化 $\textit{maxSum} = 0$。

对于 $\textit{nums}_2$ 也同理，求这两者的最大值，即为答案。

```py [sol-Python3]
class Solution:
    def solve(self, nums1: List[int], nums2: List[int]) -> int:
        max_sum = f = 0
        for x, y in zip(nums1, nums2):
            f = max(f, 0) + y - x
            max_sum = max(max_sum, f)
        return sum(nums1) + max_sum

    def maximumsSplicedArray(self, nums1: List[int], nums2: List[int]) -> int:
        return max(self.solve(nums1, nums2), self.solve(nums2, nums1))
```

```py [sol-Python3 手写 max]
class Solution:
    def solve(self, nums1: List[int], nums2: List[int]) -> int:
        max_sum = f = 0
        for x, y in zip(nums1, nums2):
            if f < 0: f = 0
            f += y - x
            if f > max_sum: max_sum = f
        return sum(nums1) + max_sum

    def maximumsSplicedArray(self, nums1: List[int], nums2: List[int]) -> int:
        return max(self.solve(nums1, nums2), self.solve(nums2, nums1))
```

```java [sol-Java]
class Solution {
    public int maximumsSplicedArray(int[] nums1, int[] nums2) {
        return Math.max(solve(nums1, nums2), solve(nums2, nums1));
    }

    private int solve(int[] nums1, int[] nums2) {
        int s1 = 0;
        int maxSum = 0;
        int f = 0;
        for (int i = 0; i < nums1.length; i++) {
            s1 += nums1[i];
            f = Math.max(f, 0) + nums2[i] - nums1[i];
            maxSum = Math.max(maxSum, f);
        }
        return s1 + maxSum;
    }
}
```

```cpp [sol-C++]
class Solution {
    int solve(vector<int>& nums1, vector<int>& nums2) {
        int s1 = 0, max_sum = 0, f = 0;
        for (int i = 0; i < nums1.size(); i++) {
            s1 += nums1[i];
            f = max(f, 0) + nums2[i] - nums1[i];
            max_sum = max(max_sum, f);
        }
        return s1 + max_sum;
    }

public:
    int maximumsSplicedArray(vector<int>& nums1, vector<int>& nums2) {
        return max(solve(nums1, nums2), solve(nums2, nums1));
    }
};
```

```go [sol-Go]
func solve(nums1, nums2 []int) int {
    var s1, maxSum, f int
    for i, x := range nums1 {
        s1 += x
        f = max(f, 0) + nums2[i] - x
        maxSum = max(maxSum, f)
    }
    return s1 + maxSum
}

func maximumsSplicedArray(nums1, nums2 []int) int {
    return max(solve(nums1, nums2), solve(nums2, nums1))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
