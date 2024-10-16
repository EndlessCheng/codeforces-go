为方便计算，把 $\textit{nums}$ 从小到大排序。

排序后，对于 $i=0,1,2,\cdots,n/2$，计算 $\textit{nums}[i]+\textit{nums}[n-1-i]$ 的最小值。最后，返回最小值除以 $2$ 的结果。

最后除以 $2$，这样可以避免在循环中做浮点运算，只在最后返回时做一次浮点运算。

注意题目保证 $n$ 是偶数。

```py [sol-Python3]
class Solution:
    def minimumAverage(self, nums: List[int]) -> float:
        nums.sort()
        return min(nums[i] + nums[-1 - i] for i in range(len(nums) // 2)) / 2
```

```java [sol-Java]
class Solution {
    public double minimumAverage(int[] nums) {
        Arrays.sort(nums);
        int ans = Integer.MAX_VALUE;
        int n = nums.length;
        for (int i = 0; i < n / 2; i++) {
            ans = Math.min(ans, nums[i] + nums[n - 1 - i]);
        }
        return ans / 2.0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    double minimumAverage(vector<int>& nums) {
        ranges::sort(nums);
        int n = nums.size();
        int ans = INT_MAX;
        for (int i = 0; i < n / 2; i++) {
            ans = min(ans, nums[i] + nums[n - 1 - i]);
        }
        return ans / 2.0;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

int cmp(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}

double minimumAverage(int* nums, int numsSize) {
    qsort(nums, numsSize, sizeof(int), cmp);
    int ans = INT_MAX;
    for (int i = 0; i < numsSize / 2; i++) {
        ans = MIN(ans, nums[i] + nums[numsSize - 1 - i]);
    }
    return ans / 2.0;
}
```

```go [sol-Go]
func minimumAverage(nums []int) float64 {
    slices.Sort(nums)
    ans := math.MaxInt
    for i, n := 0, len(nums); i < n/2; i++ {
        ans = min(ans, nums[i]+nums[n-1-i])
    }
    return float64(ans) / 2
}
```

```js [sol-JavaScript]
var minimumAverage = function(nums) {
    nums.sort((a, b) => a - b);
    let n = nums.length;
    let ans = Infinity;
    for (let i = 0; i < n / 2; i++) {
        ans = Math.min(ans, nums[i] + nums[n - 1 - i]);
    }
    return ans / 2;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_average(mut nums: Vec<i32>) -> f64 {
        nums.sort_unstable();
        (0..nums.len() / 2)
            .map(|i| nums[i] + nums[nums.len() - 1 - i])
            .min()
            .unwrap() as f64 / 2.0
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

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
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
