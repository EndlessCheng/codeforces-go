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
class Solution {
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

```js [sol-JavaScript]
var maximumCount = function(nums) {
    let neg = 0;
    let pos = 0;
    for (const x of nums) {
        if (x < 0) {
            neg++;
        } else if (x > 0) {
            pos++;
        }
    }
    return Math.max(neg, pos);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_count(nums: Vec<i32>) -> i32 {
        let mut neg = 0;
        let mut pos = 0;
        for &x in &nums {
            if x < 0 {
                neg += 1;
            } else if x > 0 {
                pos += 1;
            }
        }
        neg.max(pos)
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干额外变量。

## 方法二：二分查找

由于数组是有序的，我们可以二分找到第一个 $\ge 0$ 的数的下标 $i$，那么下标在 $[0,i-1]$ 中的数都小于 $0$，这恰好有 $i$ 个。

同样地，二分找到第一个 $> 0$ 的数的下标 $j$，那么下标在 $[j,n-1]$ 中的数都大于 $0$，这有 $n-j$ 个。

所以通过二分查找第一个 $\ge 0$ 和第一个 $> 0$ 的位置，就可以用 $\mathcal{O}(\log n)$ 的时间解决本题，原理请看 [【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

```py [sol-Python3]
class Solution:
    def maximumCount(self, nums: List[int]) -> int:
        neg = bisect_left(nums, 0)
        pos = len(nums) - bisect_right(nums, 0)
        return max(neg, pos)
```

```java [sol-Java]
class Solution {
    public int maximumCount(int[] nums) {
        int neg = lowerBound(nums, 0);
        // 第一个 > 0 的位置，等价于第一个 >= 1 的位置
        int pos = nums.length - lowerBound(nums, 1);
        return Math.max(neg, pos);
    }

    // 返回 nums 中第一个 >= target 的数的下标
    // 如果不存在这样的数，返回 nums.length
    // 详见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int target) {
        // 二分范围：开区间 (left, right)
        int left = -1;
        int right = nums.length;
        // 开区间不为空
        while (left + 1 < right) {
            // 循环不变量：
            // nums[left] < target
            // nums[right] >= target
            int mid = left + (right - left) / 2;
            if (nums[mid] >= target) {
                // 二分范围缩小至 (left, mid)
                right = mid;
            } else {
                // 二分范围缩小至 (mid, right)
                left = mid;
            }
        }
        // 此时 left 等于 right - 1
        // 因为 nums[right - 1] < target 且 nums[right] >= target，所以答案是 right
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

```cpp [sol-C++ 写法二]
class Solution {
public:
    int maximumCount(vector<int> &nums) {
        auto [left, right] = ranges::equal_range(nums, 0);
        int neg = left - nums.begin();
        int pos = nums.end() - right;
        return max(neg, pos);
    }
};
```

```go [sol-Go]
func maximumCount(nums []int) int {
    neg := sort.SearchInts(nums, 0)
    // 第一个 > 0 的位置，等价于第一个 >= 1 的位置
    pos := len(nums) - sort.SearchInts(nums, 1)
    return max(neg, pos)
}
```

```js [sol-JS]
var maximumCount = function(nums) {
    const neg = lowerBound(nums, 0);
    // 第一个 > 0 的位置，等价于第一个 >= 1 的位置
    const pos = nums.length - lowerBound(nums, 1);
    return Math.max(neg, pos);
};

// 返回 nums 中第一个 >= target 的数的下标
// 如果不存在这样的数，返回 nums.length
// 详见 https://www.bilibili.com/video/BV1AP41137w7/
function lowerBound(nums, target) {
    // 二分范围：开区间 (left, right)
    let left = -1;
    let right = nums.length;
    // 开区间不为空
    while (left + 1 < right) {
        // 循环不变量：
        // nums[left] < target
        // nums[right] >= target
        const mid = Math.floor((left + right) / 2);
        if (nums[mid] >= target) {
            // 二分范围缩小至 (left, mid)
            right = mid;
        } else {
            // 二分范围缩小至 (mid, right)
            left = mid;
        }
    }
    // 此时 left 等于 right - 1
    // 因为 nums[right - 1] < target 且 nums[right] >= target，所以答案是 right
    return right;
}
```

```js [sol-JS lodash]
var maximumCount = function(nums) {
    const neg = _.sortedIndex(nums, 0);
    const pos = nums.length - _.sortedLastIndex(nums, 0);
    return Math.max(neg, pos)
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_count(nums: Vec<i32>) -> i32 {
        let neg = nums.partition_point(|&x| x < 0);
        let pos = nums.len() - nums.partition_point(|&x| x <= 0);
        neg.max(pos) as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干额外变量。

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
