证明过程请看 [视频讲解](https://www.bilibili.com/video/BV1PV411N76R/) 第二题。

假设 $x$ 出现次数最多，其出现次数为 $\textit{maxCnt}$。

分类讨论：

- 如果 $\textit{maxCnt}\cdot 2 > n$，其余所有 $n-\textit{maxCnt}$ 个数都要与 $x$ 消除，所以最后剩下 $\textit{maxCnt}\cdot 2 - n$ 个数。
- 如果 $\textit{maxCnt}\cdot 2 \le n$ 且 $n$ 是偶数，那么可以把其余数消除至剩下 $\textit{maxCnt}$ 个数，然后再和 $x$ 消除，最后剩下 $0$ 个数。
- 如果 $\textit{maxCnt}\cdot 2 \le n$ 且 $n$ 是奇数，同上，最后剩下 $1$ 个数。

所以本题核心是计算 $\textit{maxCnt}$，这可以遍历一遍 $\textit{nums}$ 算出来。

但我们还可以更快！

由于 $\textit{nums}$ 是有序的，如果 $\textit{maxCnt}$ 超过数组长度的一半，那么 $\textit{nums}[n/2]$ 一定是出现次数最多的那个数！

按照 [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/) 的做法，可以用二分查找在 $\mathcal{O}(\log n)$ 的时间计算 $\textit{nums}[n/2]$ 第一次和最后一次出现的位置，从而算出 $\textit{maxCnt}$。

关于二分的原理，请看视频[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

```py [sol-Python3]
class Solution:
    def minLengthAfterRemovals(self, nums: List[int]) -> int:
        n = len(nums)
        x = nums[n // 2]
        max_cnt = bisect_right(nums, x) - bisect_left(nums, x)
        return max(max_cnt * 2 - n, n % 2)
```

```java [sol-Java]
class Solution {
    public int minLengthAfterRemovals(List<Integer> nums) {
        int n = nums.size();
        int x = nums.get(n / 2);
        int maxCnt = lowerBound(nums, x + 1) - lowerBound(nums, x);
        return Math.max(maxCnt * 2 - n, n % 2);
    }

    // 开区间写法
    private int lowerBound(List<Integer> nums, int target) {
        int left = -1, right = nums.size(); // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] < target
            // nums[right] >= target
            int mid = left + (right - left) / 2;
            if (nums.get(mid) < target) {
                left = mid; // 范围缩小到 (mid, right)
            } else {
                right = mid; // 范围缩小到 (left, mid)
            }
        }
        return right; // 或者 left+1
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minLengthAfterRemovals(vector<int>& nums) {
        int n = nums.size();
        int x = nums[n / 2];
        int max_cnt = ranges::upper_bound(nums, x) - ranges::lower_bound(nums, x);
        return max(max_cnt * 2 - n, n % 2);
    }
};
```

```go [sol-Go]
func minLengthAfterRemovals(nums []int) int {
	n := len(nums)
	x := nums[n/2]
	maxCnt := sort.SearchInts(nums, x+1) - sort.SearchInts(nums, x)
	return max(maxCnt*2-n, n%2)
}
```

```js [sol-JavaScript]
var minLengthAfterRemovals = function(nums) {
    const n = nums.length;
    const x = nums[n >> 1];
    const maxCnt = lowerBound(nums, x + 1) - lowerBound(nums, x);
    return Math.max(maxCnt * 2 - n, n % 2);
};

var lowerBound = function(nums, target) {
    let left = -1, right = nums.length; // 开区间 (left, right)
    while (left + 1 < right) { // 区间不为空
        // 循环不变量：
        // nums[left] < target
        // nums[right] >= target
        const mid = left + ((right - left) >> 1);
        if (nums[mid] < target) {
            left = mid; // 范围缩小到 (mid, right)
        } else {
            right = mid; // 范围缩小到 (left, mid)
        }
    }
    return right; // 或者 left+1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
