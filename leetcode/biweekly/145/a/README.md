## 思考框架

1. 理解操作在做什么。
2. 把 $\textit{nums}$ 中的数都变成一样的，能变成哪些数？
3. 如何最小化操作次数？

## 操作在做什么

题目说「选择一个整数 $h$」，哪些 $h$ 是可以选的？

根据「合法」的定义，$h$ 不能低于 $\textit{nums}$ 的次大值。比如 $\textit{nums}=[5,2,5,4,5]$，$h$ 不能小于次大值 $4$，否则大于 $h$ 的数不都相等。

所以操作只能改变大于次大值的数，也就是**最大值**。

## 能变成哪些数

要让所有数都变成 $k$，前提条件是所有数都变成一样的。那么，能变成哪些数呢？

仍然以 $\textit{nums}=[5,2,5,4,5]$ 为例。选择 $h=4$，可以把最大值 $5$ 改成次大值 $4$。修改后 $\textit{nums}=[4,2,4,4,4]$，有更多的数都相同。并且修改后，原来的次大值 $4$ 就变成最大值了。下一次操作，我们就可以选择更小的 $h$，把更多的数都变成一样的。

选择 $h=2$，可以把最大值 $4$ 改成次大值 $2$。修改后 $\textit{nums}=[2,2,2,2,2]$，所有数都相同。

如果想继续改成比 $2$ 小的数，比如 $0$，选择 $h=0$ 即可。

所以，$\textit{nums}$ 中的数可以都变成任意 $\le \min(\textit{nums})$ 的数。

## 最小化操作次数

为了最小化操作次数，每次选 $h$ 为次大值是最优的。贪心地想，能一步到达次大值，没必要分好几步。

分类讨论：

- 如果 $k > \min(nums)$，无法满足要求，返回 $-1$。
- 如果 $k = \min(nums)$，操作次数为 $\textit{nums}$ 中的不同元素个数减一。比如 $[5,2,5,4,5]\to [4,2,4,4,4]\to [2,2,2,2,2]$，最大值 $5\to 4\to 2$，用了 $2$ 次操作。
- 如果 $k < \min(nums)$，操作次数为 $\textit{nums}$ 中的不同元素个数。因为都变成 $\min(nums)$ 后，还需要再操作一次，才能都变成 $k$。

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: List[int], k: int) -> int:
        mn = min(nums)
        if k > mn:
            return -1
        return len(set(nums)) - (k == mn)
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums, int k) {
        int min = Arrays.stream(nums).min().getAsInt();
        if (k > min) {
            return -1;
        }
        int distinctCount = (int) Arrays.stream(nums).distinct().count();
        return distinctCount - (k == min ? 1 : 0);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int>& nums, int k) {
        int mn = ranges::min(nums);
        if (k > mn) {
            return -1;
        }
        unordered_set<int> st(nums.begin(), nums.end());
        return st.size() - (k == mn);
    }
};
```

```go [sol-Go]
func minOperations(nums []int, k int) int {
    mn := slices.Min(nums)
    if k > mn {
        return -1
    }

    set := map[int]struct{}{}
    for _, x := range nums {
        set[x] = struct{}{}
    }
    if k == mn {
        return len(set) - 1
    }
    return len(set)
}
```

```js [sol-JavaScript]
var minOperations = function(nums, k) {
    const min = Math.min(...nums);
    if (k > min) {
        return -1;
    }
    return new Set(nums).size - (k === min ? 1 : 0);
};
```

```rust [sol-Rust]
use std::collections::HashSet;

impl Solution {
    pub fn min_operations(nums: Vec<i32>, k: i32) -> i32 {
        let min = *nums.iter().min().unwrap();
        if k > min {
            return -1;
        }
        let set = nums.into_iter().collect::<HashSet<_>>();
        set.len() as i32 - if k == min { 1 } else { 0 }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
