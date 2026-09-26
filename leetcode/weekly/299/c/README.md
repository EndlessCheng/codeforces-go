设 $S_1 = \sum\limits_{i}\textit{nums}_1[i]$。

交换下标在 $[\textit{left},\textit{right}]$ 内的元素后，新的 $\textit{nums}'_1$ 的元素和为

$$
S_1 - (\textit{nums}_1[\textit{left}] + \cdots + \textit{nums}_1[\textit{right}]) + (\textit{nums}_2[\textit{left}] + \cdots + \textit{nums}_2[\textit{right}])
$$

合并相同下标，上式变形为

$$
S_1 + (\textit{nums}_2[\textit{left}]-\textit{nums}_1[\textit{left}]) + \cdots + (\textit{nums}_2[\textit{right}]-\textit{nums}_1[\textit{right}])
$$

定义 $\textit{diff}[i] = \textit{nums}_2[i]-\textit{nums}_1[i]$，上式为

$$
S_1 + \textit{diff}[\textit{left}] + \cdots + \textit{diff}[\textit{right}]
$$

为了最大化上式，我们需要最大化 $\textit{diff}$ 数组的 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)，[我的题解](https://leetcode.cn/problems/maximum-subarray/solutions/2533977/qian-zhui-he-zuo-fa-ben-zhi-shi-mai-mai-abu71/)。注意子数组可以为空，所以初始化 $\textit{maxSum} = 0$。

对于 $\textit{nums}_2$ 交换后的最大和，计算方式同理。

[本题视频讲解](https://www.bilibili.com/video/BV1pW4y1r7xs)，欢迎点赞关注~

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

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}_i$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
