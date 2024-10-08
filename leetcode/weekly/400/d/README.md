怎么计算连续子数组的 OR？

首先，我们有如下 $\mathcal{O}(n^2)$ 的暴力算法：

外层循环，从 $i=0$ 开始，**从左到右**遍历 $\textit{nums}$。内层循环，从 $j=i-1$ 开始，**从右到左**遍历 $\textit{nums}$，更新 $\textit{nums}[j]=\textit{nums}[j]\ \vert\ \textit{nums}[i]$。

- $i=1$ 时，我们会把 $\textit{nums}[0]$ 到 $\textit{nums}[1]$ 的 OR 记录在 $\textit{nums}[0]$ 中。 
- $i=2$ 时，我们会把 $\textit{nums}[1]$ 到 $\textit{nums}[2]$ 的 OR 记录在 $\textit{nums}[1]$ 中，$\textit{nums}[0]$ 到 $\textit{nums}[2]$ 的 OR 记录在 $\textit{nums}[0]$ 中。
- $i=3$ 时，我们会把 $\textit{nums}[2]$ 到 $\textit{nums}[3]$ 的 OR 记录在 $\textit{nums}[2]$ 中；$\textit{nums}[1]$ 到 $\textit{nums}[3]$ 的 OR 记录在 $\textit{nums}[1]$ 中；$\textit{nums}[0]$ 到 $\textit{nums}[3]$ 的 OR 记录在 $\textit{nums}[0]$ 中。
- 依此类推。

按照该算法，可以计算出所有子数组的 OR。注意单个元素也算子数组。

为方便大家理解后续优化，先写出暴力代码：

```py [sol-Python3]
# 暴力算法，会超时
class Solution:
    def minimumDifference(self, nums: List[int], k: int) -> int:
        ans = inf
        for i, x in enumerate(nums):
            ans = min(ans, abs(x - k))  # 单个元素也算子数组
            for j in range(i - 1, -1, -1):
                nums[j] |= x  # 现在 nums[j] = 原数组 nums[j] 到 nums[i] 的 OR
                ans = min(ans, abs(nums[j] - k))
        return ans
```

```java [sol-Java]
// 暴力算法，会超时
class Solution {
    public int minimumDifference(int[] nums, int k) {
        int ans = Integer.MAX_VALUE;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            ans = Math.min(ans, Math.abs(x - k)); // 单个元素也算子数组
            for (int j = i - 1; j >= 0; j--) {
                nums[j] |= x; // 现在 nums[j] = 原数组 nums[j] 到 nums[i] 的 OR
                ans = Math.min(ans, Math.abs(nums[j] - k));
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
// 暴力算法，会超时
class Solution {
public:
    int minimumDifference(vector<int>& nums, int k) {
        int ans = INT_MAX;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            ans = min(ans, abs(x - k)); // 单个元素也算子数组
            for (int j = i - 1; j >= 0; j--) {
                nums[j] |= x; // 现在 nums[j] = 原数组 nums[j] 到 nums[i] 的 OR
                ans = min(ans, abs(nums[j] - k));
            }
        }
        return ans;
    }
};
```

```c [sol-C]
// 暴力算法，会超时
#define MIN(a, b) ((a) < (b) ? (a) : (b))

int minimumDifference(int* nums, int numsSize, int k) {
    int ans = INT_MAX;
    for (int i = 0; i < numsSize; i++) {
        int x = nums[i];
        ans = MIN(ans, abs(x - k)); // 单个元素也算子数组
        for (int j = i - 1; j >= 0; j--) {
            nums[j] |= x; // 现在 nums[j] = 原数组 nums[j] 到 nums[i] 的 OR
            ans = MIN(ans, abs(nums[j] - k));
        }
    }
    return ans;
}
```

```go [sol-Go]
// 暴力算法，会超时
func minimumDifference(nums []int, k int) int {
    ans := math.MaxInt
    for i, x := range nums {
        ans = min(ans, abs(x-k)) // 单个元素也算子数组
        for j := i - 1; j >= 0; j-- {
            nums[j] |= x // 现在 nums[j] = 原数组 nums[j] 到 nums[i] 的 OR
            ans = min(ans, abs(nums[j]-k))
        }
    }
    return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

```js [sol-JavaScript]
// 暴力算法，会超时
var minimumDifference = function(nums, k) {
    let ans = Infinity;
    for (let i = 0; i < nums.length; i++) {
        const x = nums[i];
        ans = Math.min(ans, Math.abs(x - k)); // 单个元素也算子数组
        for (let j = i - 1; j >= 0; j--) {
            nums[j] |= x; // 现在 nums[j] = 原数组 nums[j] 到 nums[i] 的 OR
            ans = Math.min(ans, Math.abs(nums[j] - k));
        }
    }
    return ans;
};
```

```rust [sol-Rust]
// 暴力算法，会超时
impl Solution {
    pub fn minimum_difference(mut nums: Vec<i32>, k: i32) -> i32 {
        let mut ans = i32::MAX;
        for i in 0..nums.len() {
            let x = nums[i];
            ans = ans.min((x - k).abs()); // 单个元素也算子数组
            for j in (0..i).rev() {
                nums[j] |= x; // 现在 nums[j] = 原数组 nums[j] 到 nums[i] 的 OR
                ans = ans.min((nums[j] - k).abs());
            }
        }
        ans
    }
}
```

暴力算法没有充分利用 OR 运算的性质。为了优化，我们需要考察上述过程中，这些元素之间有什么关系。

为方便理解，我们从**集合**的角度来看上述暴力过程。

前置知识：[从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

把二进制数看成集合，两个数的 OR 就是两个集合的**并集**。

- 把 $\textit{nums}[i]$ 对应的集合记作 $A_i$。
- $i=1$ 时，我们会把 $A_0$ 到 $A_1$ 的并集记录在 $A_0$ 中，也就是把 $A_1$ 并入 $A_0$。所以 $A_1$ 必然是 $A_0$ 的子集，即 $A_0 \supseteq A_1$。
- $i=2$ 时，我们会把 $A_2$ 并入 $A_1$ 和 $A_0$，所以有 $A_0 \supseteq A_1 \supseteq A_2$。
- $i=3$ 时，我们会把 $A_3$ 并入 $A_2$、$A_1$ 和 $A_0$，所以有 $A_0 \supseteq A_1 \supseteq A_2 \supseteq A_3$。
- 一般地，上述代码的内层循环结束时，有 $A_0 \supseteq A_1 \supseteq A_2 \supseteq \cdots \supseteq A_i$。

想一想，如果 $A_i$ 是 $A_j$ 的子集，那么内层循环还需要继续跑吗？

不需要。如果 $A_i$ 已经是 $A_j$ 的子集，那么 $A_i$ 必然也是更左边的 $A_0,A_1,A_2,\cdots,A_{j-1}$ 的子集。既然 $A_i$ 都已经是这些集合的子集了，那么并入操作不会改变这些集合。

所以当我们发现 $A_i$ 是 $A_j$ 的子集时，就可以退出内层循环了。

具体到代码，对于两个二进制数 $a$ 和 $b$，如果 $a\ \vert\ b = a$，那么 $b$ 对应的集合是 $a$ 对应的集合的子集。

具体例子可以看 [视频讲解](https://www.bilibili.com/video/BV1Qx4y1E7zj/) 第四题（计算的是子数组 AND）。

```py [sol-Python3]
class Solution:
    def minimumDifference(self, nums: List[int], k: int) -> int:
        ans = inf
        for i, x in enumerate(nums):
            ans = min(ans, abs(x - k))
            j = i - 1
            # 如果 x 是 nums[j] 的子集，就退出循环
            while j >= 0 and nums[j] | x != nums[j]:
                nums[j] |= x
                ans = min(ans, abs(nums[j] - k))
                j -= 1
        return ans
```

```java [sol-Java]
class Solution {
    public int minimumDifference(int[] nums, int k) {
        int ans = Integer.MAX_VALUE;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            ans = Math.min(ans, Math.abs(x - k));
            // 如果 x 是 nums[j] 的子集，就退出循环
            for (int j = i - 1; j >= 0 && (nums[j] | x) != nums[j]; j--) {
                nums[j] |= x;
                ans = Math.min(ans, Math.abs(nums[j] - k));
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumDifference(vector<int>& nums, int k) {
        int ans = INT_MAX;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            ans = min(ans, abs(x - k));
            // 如果 x 是 nums[j] 的子集，就退出循环
            for (int j = i - 1; j >= 0 && (nums[j] | x) != nums[j]; j--) {
                nums[j] |= x;
                ans = min(ans, abs(nums[j] - k));
            }
        }
        return ans;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

int minimumDifference(int* nums, int numsSize, int k) {
    int ans = INT_MAX;
    for (int i = 0; i < numsSize; i++) {
        int x = nums[i];
        ans = MIN(ans, abs(x - k));
        // 如果 x 是 nums[j] 的子集，就退出循环
        for (int j = i - 1; j >= 0 && (nums[j] | x) != nums[j]; j--) {
            nums[j] |= x;
            ans = MIN(ans, abs(nums[j] - k));
        }
    }
    return ans;
}

```

```go [sol-Go]
func minimumDifference(nums []int, k int) int {
    ans := math.MaxInt
    for i, x := range nums {
        ans = min(ans, abs(x-k))
        // 如果 x 是 nums[j] 的子集，就退出循环
        for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
            nums[j] |= x
            ans = min(ans, abs(nums[j]-k))
        }
    }
    return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

```js [sol-JavaScript]
var minimumDifference = function(nums, k) {
    let ans = Infinity;
    for (let i = 0; i < nums.length; i++) {
        const x = nums[i];
        ans = Math.min(ans, Math.abs(x - k));
        // 如果 x 是 nums[j] 的子集，就退出循环
        for (let j = i - 1; j >= 0 && (nums[j] | x) !== nums[j]; j--) {
            nums[j] |= x;
            ans = Math.min(ans, Math.abs(nums[j] - k));
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_difference(mut nums: Vec<i32>, k: i32) -> i32 {
        let mut ans = i32::MAX;
        for i in 0..nums.len() {
            let x = nums[i];
            ans = ans.min((x - k).abs());
            let mut j = i - 1;
            // 如果 x 是 nums[j] 的子集，就退出循环
            while j < nums.len() && (nums[j] | x) != nums[j] {
                nums[j] |= x;
                ans = ans.min((nums[j] - k).abs());
                j -= 1;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。由于 $2^{29}-1<10^9<2^{30}-1$，二进制数对应集合的大小不会超过 $29$，因此在 OR 运算下，每个数字至多可以增大 $29$ 次。**总体上看**，二重循环的总循环次数等于每个数字可以增大的次数之和，即 $O(n\log U)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

1. 把 OR 换成 AND 怎么做？[1521. 找到最接近目标值的函数值](https://leetcode.cn/problems/find-a-value-of-a-mysterious-function-closest-to-target/)
2. 把 OR 换成 GCD 怎么做？
3. 把 OR 换成 LCM 怎么做？

欢迎在评论区发表你的思路/代码。

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
