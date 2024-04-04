## 方法一：遍历

遍历数组，统计负数数目 $\textit{neg}$ 和正数数目 $\textit{pos}$。

最后返回 $\max(\textit{neg}, \textit{pos})$。

```py [sol-Python3]
class Solution:
    def maximumCount(self, nums: List[int]) -> int:
        neg = pos = 0
        for x in nums:
            if x < 0:
                neg += 1
            elif x > 0:
                pos += 1
        return max(neg, pos)
```

```java [sol-Java]
public class Solution {
    public int maximumCount(int[] nums) {
        int neg = 0;
        int pos = 0;
        for (int x : nums) {
            if (x < 0) {
                neg++;
            } else if (x > 0) {
                pos++;
            }
        }
        return Math.max(neg, pos);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumCount(vector<int> &nums) {
        int neg = 0, pos = 0;
        for (int x : nums) {
            if (x < 0) {
                neg++;
            } else if (x > 0) {
                pos++;
            }
        }
        return max(neg, pos);
    }
};
```

```go [sol-Go]
func maximumCount(nums []int) int {
    neg, pos := 0, 0
    for _, x := range nums {
        if x < 0 {
            neg++
        } else if x > 0 {
            pos++
        }
    }
    return max(neg, pos)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干额外变量。

## 方法二：二分查找

由于数组是有序的，我们可以二分找到第一个 $\ge 0$ 的数的下标 $i$，那么 $[0,i-1]$ 中的数都小于 $0$，这恰好有 $i$ 个。

同样地，二分找到第一个 $> 0$ 的数的下标 $j$，那么 $[j,n-1]$ 中的数都大于 $0$，这有 $n-j$ 个。

所以通过二分查找第一个 $\ge 0$ 和第一个 $> 0$ 的位置，就可以用 $\mathcal{O}(\log n)$ 的时间解决本题，原理请看 [【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

注：第一个 $>0$ 的位置，等价于第一个 $\ge 1$ 的位置。

```py [sol-Python3]
class Solution:
    def maximumCount(self, nums: List[int]) -> int:
        neg = bisect_left(nums, 0)
        pos = len(nums) - bisect_right(nums, 0)
        return max(neg, pos)
```

```java [sol-Java]
public class Solution {
    public int maximumCount(int[] nums) {
        int neg = lowerBound(nums, 0);
        int pos = nums.length - lowerBound(nums, 1);
        return Math.max(neg, pos);
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int target) {
        int left = -1, right = nums.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] < target
            // nums[right] >= target
            int mid = left + (right - left) / 2;
            if (nums[mid] < target) {
                left = mid; // 范围缩小到 (mid, right)
            } else {
                right = mid; // 范围缩小到 (left, mid)
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumCount(vector<int> &nums) {
        int neg = ranges::lower_bound(nums, 0) - nums.begin();
        int pos = nums.end() - ranges::upper_bound(nums, 0);
        return max(neg, pos);
    }
};
```

```go [sol-Go]
func maximumCount(nums []int) int {
    neg := sort.SearchInts(nums, 0)
    pos := len(nums) - sort.SearchInts(nums, 1)
    return max(neg, pos)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干额外变量。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndnegCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
