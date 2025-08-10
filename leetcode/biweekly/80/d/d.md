**前置知识**：[滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)

滑动窗口**使用前提**：

1. 连续子数组/子串。
2. 有单调性。本题元素均为正数，所以子数组越长，分数越高；子数组越短，分数越低。这意味着只要某个子数组的分数小于 $k$，在该子数组内的更短的子数组，分数也小于 $k$。

**做法**：

1. 枚举子数组的右端点 $\textit{right}$，同时维护窗口内的元素和 $\textit{sum}$ 以及窗口左端点 $\textit{left}$。窗口的分数为 $\textit{sum}\cdot (\textit{right}-\textit{left}+1)$。
2. 元素 $x=\textit{nums}[i]$ 进入窗口，把 $\textit{sum}$ 增加 $x$。
3. 如果窗口的分数 $\ge k$，那么把左端点元素 $\textit{nums}[\textit{left}]$ 移出窗口，同时减少 $\textit{sum}$，把 $\textit{left}$ 加一。
4. 内层循环结束后，$[\textit{left},\textit{right}]$ 这个子数组是合法的。根据上面的使用前提 2，在这个子数组内部的子数组也是合法的，其中右端点小于 $\textit{right}$ 的子数组，我们在之前的循环中已经统计过了，这里只需要统计右端点**恰好等于** $\textit{right}$ 的合法子数组，即
   $$
   [\textit{left},\textit{right}],[\textit{left}+1,\textit{right}],\ldots,[\textit{right},\textit{right}]
   $$
   这一共有 $\textit{right}-\textit{left}+1$ 个，加入答案。

```py [sol-Python3]
class Solution:
    def countSubarrays(self, nums: List[int], k: int) -> int:
        ans = s = left = 0
        for right, x in enumerate(nums):
            s += x
            while s * (right - left + 1) >= k:
                s -= nums[left]
                left += 1
            ans += right - left + 1
        return ans
```

```java [sol-Java]
class Solution {
    public long countSubarrays(int[] nums, long k) {
        long ans = 0;
        long sum = 0;
        int left = 0;
        for (int right = 0; right < nums.length; right++) {
            sum += nums[right];
            while (sum * (right - left + 1) >= k) {
                sum -= nums[left];
                left++;
            }
            ans += right - left + 1;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countSubarrays(vector<int>& nums, long long k) {
        long long ans = 0, sum = 0;
        int left = 0;
        for (int right = 0; right < nums.size(); right++) {
            sum += nums[right];
            while (sum * (right - left + 1) >= k) {
                sum -= nums[left];
                left++;
            }
            ans += right - left + 1;
        }
        return ans;
    }
};
```

```c [sol-C]
long long countSubarrays(int* nums, int numsSize, long long k) {
    long long ans = 0, sum = 0;
    int left = 0;
    for (int right = 0; right < numsSize; right++) {
        sum += nums[right];
        while (sum * (right - left + 1) >= k) {
            sum -= nums[left];
            left++;
        }
        ans += right - left + 1;
    }
    return ans;
}
```

```go [sol-Go]
func countSubarrays(nums []int, k int64) (ans int64) {
    sum, left := int64(0), 0
    for right, x := range nums {
        sum += int64(x)
        for sum*int64(right-left+1) >= k {
            sum -= int64(nums[left])
            left++
        }
        ans += int64(right - left + 1)
    }
    return
}
```

```js [sol-JavaScript]
var countSubarrays = function(nums, k) {
    let ans = 0, sum = 0, left = 0;
    for (let right = 0; right < nums.length; right++) {
        sum += nums[right];
        while (sum * (right - left + 1) >= k) {
            sum -= nums[left];
            left++;
        }
        ans += right - left + 1;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_subarrays(nums: Vec<i32>, k: i64) -> i64 {
        let mut ans = 0;
        let mut sum = 0;
        let mut left = 0;
        for (right, &x) in nums.iter().enumerate() {
            sum += x as i64;
            while sum * ((right - left + 1) as i64) >= k {
                sum -= nums[left] as i64;
                left += 1;
            }
            ans += (right - left + 1) as i64;
        }
        ans
    }
}
```

## 思考题

本题的分数是子数组的元素和乘以子数组的长度，如果把这里的乘法改成除法，即 $\dfrac{子数组的元素和}{子数组的长度} < k$，要怎么做？

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。虽然写了个二重循环，但是内层循环中对 $\textit{left}$ 加一的**总**执行次数不会超过 $n$ 次，所以总的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面滑动窗口题单的「**§2.3.2 越短越合法**」。

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
