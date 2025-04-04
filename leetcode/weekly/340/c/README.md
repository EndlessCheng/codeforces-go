### 前置知识：二分

见 [二分查找【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

### 提示 1

看到「最大化最小值」或者「最小化最大值」就要想到**二分答案**，这是一个固定的套路。

为什么？一般来说，二分的值越大，越能/不能满足要求；二分的值越小，越不能/能满足要求，有单调性，可以二分。

### 提示 2

二分数对中的最大差值 $mx$。

由于下标和答案无关，可以先排序。为了让匹配的数对尽量多，应尽量选相邻的元素，这样更能满足要求。例如 $[1,2,3,4]$，如果 $1,3$ 匹配，$2,4$ 匹配，最大差值是 $2$；而如果 $1,2$ 相邻匹配，$3,4$ 相邻匹配，最大差值只有 $1$。

我们来算一算最多能匹配多少个数对：

- 如果可以选 $\textit{nums}[0]$ 和 $\textit{nums}[1]$，那么答案等于「$n-2$ 个数的最多数对个数」$+1$。
- 如果不选 $\textit{nums}[0]$，那么答案等于「$n-1$ 个数的最多数对个数」。
- 这两种情况取最大值。

这看上去很像 [198. 打家劫舍](https://leetcode.cn/problems/house-robber/)，可以用动态规划实现。

也可以用贪心做：

- 注意到，「$n-1$ 个数的最多数对个数」不会超过「$n-3$ 个数的最多数对个数」$+1$。这里 $+1$ 表示选 $\textit{nums}[1]$ 和 $\textit{nums}[2]$。
- 由于「$n-2$ 个数的最多数对个数」$\ge$「$n-3$ 个数的最多数对个数」，所以如果可以选 $\textit{nums}[0]$ 和 $\textit{nums}[1]$，那么直接选就行。
- 依此类推，不断缩小问题规模。所以遍历一遍数组就能求出最多数对个数，具体见代码。

[本题视频讲解](https://www.bilibili.com/video/BV1iN411w7my/)

```py [sol-Python3]
class Solution:
    def minimizeMax(self, nums: List[int], p: int) -> int:
        nums.sort()
        def check(mx: int) -> int:
            cnt = i = 0
            while i < len(nums) - 1:
                if nums[i + 1] - nums[i] <= mx:  # 都选
                    cnt += 1
                    i += 2
                else:  # 不选 nums[i]
                    i += 1
            return cnt >= p
        return bisect_left(range(nums[-1] - nums[0]), True, key=check)
```

```java [sol-Java]
class Solution {
    public int minimizeMax(int[] nums, int p) {
        Arrays.sort(nums);
        int left = -1;
        int right = nums[nums.length - 1] - nums[0]; // 开区间
        while (left + 1 < right) { // 开区间
            int mid = (left + right) >>> 1;
            if (check(mid, nums, p)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(int mx, int[] nums, int p) {
        int cnt = 0;
        for (int i = 0; i < nums.length - 1; i++) {
            if (nums[i + 1] - nums[i] <= mx) { // 都选
                cnt++;
                i++;
            }
        }
        return cnt >= p;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimizeMax(vector<int>& nums, int p) {
        ranges::sort(nums);
        auto check = [&](int mx) -> bool {
            int cnt = 0;
            for (int i = 0; i < nums.size() - 1; i++) {
                if (nums[i + 1] - nums[i] <= mx) { // 都选
                    cnt++;
                    i++;
                }
            }
            return cnt >= p;
        };
        int left = -1, right = nums.back() - nums[0]; // 开区间
        while (left + 1 < right) { // 开区间
            int mid = left + (right - left) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func minimizeMax(nums []int, p int) int {
    slices.Sort(nums)
    n := len(nums)
    return sort.Search(nums[n-1]-nums[0], func(mx int) bool {
        cnt := 0
        for i := 0; i < n-1; i++ {
            if nums[i+1]-nums[i] <= mx { // 都选
                cnt++
                i++
            }
        }
        return cnt >= p
    })
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})-\min(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序时的栈空间，仅用到若干额外变量。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
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
