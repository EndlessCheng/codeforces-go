## 转化

假设数对差值 $\le \textit{mx}$。

$\textit{mx}$ 越大，要求就越**宽松**，越能找到 $p$ 个数对。

$\textit{mx}$ 越小，要求就越**苛刻**，越不能找到 $p$ 个数对。

据此，可以**二分猜答案**。关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

现在问题变成一个判定性问题：

- 给定 $\textit{mx}$，能否不重复地选择（至少）$p$ 个数对，每个数对的差值（绝对差）都 $\le \textit{mx}$？

## 思路

为方便计算，先把 $\textit{nums}$ 排序，这样选相邻元素（相差更小）比选不相邻元素更优。例如 $[1,2,3,4]$，如果选 $(1,3)$ 和 $(2,4)$，最大差值是 $2$；而如果选 $(1,2)$ 和 $(3,4)$，最大差值只有 $1$。

最多能选多少个数对？

定义 $f(n)$ 表示在 $\textit{nums}$ 的后 $n$ 个数中选出的最多数对个数。

讨论 $\textit{nums}[0]$ **选或不选**：

- 如果不选 $\textit{nums}[0]$，那么答案等于剩下的「后 $n-1$ 个数的最多数对个数」，即 $f(n-1)$。
- 如果选 $\textit{nums}[0]$ 和 $\textit{nums}[1]$（前提是差值 $\le \textit{mx}$），那么答案等于剩下的「后 $n-2$ 个数的最多数对个数」加一，即 $f(n-2)+1$。
- 两种情况取最大值，得 $f(n) = \max(f(n-1), f(n-2)+1)$。
- 如果无法选 $\textit{nums}[0]$，则 $f(n) = f(n-1)$。

这类似 [198. 打家劫舍](https://leetcode.cn/problems/house-robber/)，可以用动态规划实现。

但实际上，可以贪心。

注意到，$f(n-1)$ 至多为 $f(n-3)+1$。这里加一表示选 $\textit{nums}[1]$ 和 $\textit{nums}[2]$。

此外，由于元素个数越少，$f(i)$ 越小，所以 $f(n-2)\ge f(n-3)$，所以 

$$
f(n-2)+1\ge f(n-3)+1 \ge f(n-1)
$$

因此

$$
f(n) = \max(f(n-1), f(n-2)+1) = f(n-2)+1
$$

所以如果可以选 $\textit{nums}[0]$ 和 $\textit{nums}[1]$，那么直接选 $\textit{nums}[0]$ 和 $\textit{nums}[1]$ 就是最优的。

## 细节

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的，喜欢哪种写法就用哪种。

- 开区间左端点初始值：$-1$。绝对差不可能是负数，一定无法满足要求。
- 开区间右端点初始值：$\max(\textit{nums}) - \min(\textit{nums})$。可以随便选，所以一定满足要求。注意题目保证 $p\le n/2$。

对于开区间写法，简单来说 `check(mid) == true` 时更新的是谁，最后就返回谁。相比其他二分写法，开区间写法不需要思考加一减一等细节，更简单。推荐使用开区间写二分。

## 答疑

**问**：为什么二分结束后，答案 $\textit{ans}$ 一定来自 $\textit{nums}$ 中某两个数的差值？

**答**：反证法。假设 $\textit{ans}$ 不来自 $\textit{nums}$ 中某两个数的差值，这意味着最大差值 $\le \textit{ans}-1$，换句话说，$\text{check}(\textit{ans}-1)=\texttt{true}$。但根据循环不变量，二分结束后 $\text{check}(\textit{ans}-1)=\texttt{false}$，矛盾。故原命题成立。

[本题视频讲解](https://www.bilibili.com/video/BV1iN411w7my/)

```py [sol-Python3]
class Solution:
    def minimizeMax(self, nums: List[int], p: int) -> int:
        nums.sort()
        def check(mx: int) -> int:
            cnt = i = 0
            while i < len(nums) - 1:
                if nums[i + 1] - nums[i] <= mx:  # 选 nums[i] 和 nums[i+1]
                    cnt += 1
                    i += 2
                else:  # 不选 nums[i]
                    i += 1
            return cnt >= p

        left, right = -1, nums[-1] - nums[0]
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid):
                right = mid
            else:
                left = mid
        return right
```

```py [sol-Python3 库函数]
class Solution:
    def minimizeMax(self, nums: List[int], p: int) -> int:
        nums.sort()
        def check(mx: int) -> int:
            cnt = i = 0
            while i < len(nums) - 1:
                if nums[i + 1] - nums[i] <= mx:  # 选 nums[i] 和 nums[i+1]
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
        int right = nums[nums.length - 1] - nums[0];
        while (left + 1 < right) {
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
            if (nums[i + 1] - nums[i] <= mx) { // 选 nums[i] 和 nums[i+1]
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
            for (int i = 0; i + 1 < nums.size(); i++) {
                if (nums[i + 1] - nums[i] <= mx) { // 选 nums[i] 和 nums[i+1]
                    cnt++;
                    i++;
                }
            }
            return cnt >= p;
        };

        int left = -1, right = nums.back() - nums[0];
        while (left + 1 < right) {
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
            if nums[i+1]-nums[i] <= mx { // 选 nums[i] 和 nums[i+1]
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
- 空间复杂度：$\mathcal{O}(1)$。忽略排序时的栈空间。

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
