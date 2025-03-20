## 方法一：前后缀分解

### 提示 1

要让答案最大，首先应当最大化答案的二进制**长度**。

### 提示 2

把「乘 $2$」分配给多个数（雨露均沾），不如只分配给一个数，这样能得到更长（更大）的答案。

### 提示 3

**枚举**把 $\textit{nums}[i]$ 乘 $k$ 次 $2$（左移 $k$ 次）。修改后，如何计算所有元素的或值？

利用 [238. 除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/) 的 [做法](https://leetcode.cn/problems/product-of-array-except-self/solutions/2783788/qian-hou-zhui-fen-jie-fu-ti-dan-pythonja-86r1/)，预处理每个 $\textit{nums}[i]$ 左侧元素的或值 $\textit{pre}$，以及右侧元素的或值 $\textit{suf}$，从而 $\mathcal{O}(1)$ 得到把 $\textit{nums}[i]$ 乘 $k$ 次 $2$ 后，所有元素的或值。

代码实现时，只需预处理右侧元素的或值（后缀或），左侧元素的或值（前缀或）可以一边遍历 $\textit{nums}$ 一边计算。

```py [sol-Python3]
class Solution:
    def maximumOr(self, nums: List[int], k: int) -> int:
        n = len(nums)
        # suf[i] 表示 nums[i+1:] 的 OR
        suf = [0] * n
        for i in range(n - 2, -1, -1):
            suf[i] = suf[i + 1] | nums[i + 1]

        # pre 表示 nums[:i] 的 OR
        ans = pre = 0
        for x, suf_or in zip(nums, suf):
            ans = max(ans, pre | (x << k) | suf_or)
            pre |= x
        return ans
```

```java [sol-Java]
class Solution {
    public long maximumOr(int[] nums, int k) {
        int n = nums.length;
        // suf[i] 表示 nums[i+1] 到 nums[n-1] 的 OR
        int[] suf = new int[n];
        for (int i = n - 2; i >= 0; i--) {
            suf[i] = suf[i + 1] | nums[i + 1];
        }

        long ans = 0;
        // pre 表示 nums[0] 到 nums[i-1] 的 OR
        int pre = 0;
        for (int i = 0; i < n; i++) {
            ans = Math.max(ans, pre | ((long) nums[i] << k) | suf[i]);
            pre |= nums[i];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumOr(vector<int>& nums, int k) {
        int n = nums.size();
        // suf[i] 表示 nums[i+1] 到 nums[n-1] 的 OR
        vector<int> suf(n);
        for (int i = n - 2; i >= 0; i--) {
            suf[i] = suf[i + 1] | nums[i + 1];
        }

        long long ans = 0;
        // pre 表示 nums[0] 到 nums[i-1] 的 OR
        int pre = 0;
        for (int i = 0; i < n; i++) {
            ans = max(ans, pre | ((long long) nums[i] << k) | suf[i]);
            pre |= nums[i];
        }
        return ans;
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

long long maximumOr(int* nums, int numsSize, int k) {
    // suf[i] 表示 nums[i+1] 到 nums[n-1] 的 OR
    int* suf = malloc(numsSize * sizeof(int));
    suf[numsSize - 1] = 0;
    for (int i = numsSize - 2; i >= 0; i--) {
        suf[i] = suf[i + 1] | nums[i + 1];
    }

    long long ans = 0;
    // pre 表示 nums[0] 到 nums[i-1] 的 OR
    int pre = 0;
    for (int i = 0; i < numsSize; i++) {
        int x = nums[i];
        ans = MAX(ans, pre | ((long long) x << k) | suf[i]);
        pre |= x;
    }

    free(suf);
    return ans;
}
```

```go [sol-Go]
func maximumOr(nums []int, k int) int64 {
    n := len(nums)
    // suf[i] 表示 nums[i+1] 到 nums[n-1] 的 OR
    suf := make([]int, n)
    for i := n - 2; i >= 0; i-- {
        suf[i] = suf[i+1] | nums[i+1]
    }

    // pre 表示 nums[0] 到 nums[i-1] 的 OR
    ans, pre := 0, 0
    for i, x := range nums {
        ans = max(ans, pre|x<<k|suf[i])
        pre |= x
    }
    return int64(ans)
}
```

```js [sol-JavaScript]
var maximumOr = function(nums, k) {
    const n = nums.length;
    // suf[i] 表示 nums[i+1] 到 nums[n-1] 的 OR
    const suf = Array(n);
    suf[n - 1] = 0;
    for (let i = n - 2; i >= 0; i--) {
        suf[i] = suf[i + 1] | nums[i + 1];
    }

    // pre 表示 nums[0] 到 nums[i-1] 的 OR
    let ans = 0n, pre = 0n;
    for (let i = 0; i < n; i++) {
        const x = BigInt(nums[i]);
        const res = pre | (x << BigInt(k)) | BigInt(suf[i]);
        ans = res > ans ? res : ans;
        pre |= x;
    }
    return Number(ans);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_or(nums: Vec<i32>, k: i32) -> i64 {
        let n = nums.len();
        // suf[i] 表示 nums[i+1] 到 nums[n-1] 的 OR
        let mut suf = vec![0; n];
        for i in (0..n - 1).rev() {
            suf[i] = suf[i + 1] | nums[i + 1];
        }

        let mut ans = 0;
        // pre 表示 nums[0] 到 nums[i-1] 的 OR
        let mut pre = 0;
        for (x, suf_or) in nums.into_iter().zip(suf) {
            ans = ans.max(pre | ((x as i64) << k) | suf_or as i64);
            pre |= x as i64;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：空间优化

设 $\textit{nums}$ 所有数的 OR 为 $\textit{allOr}$。

能否直接算出从 $\textit{allOr}$ 中去掉 $x=\textit{nums}[i]$ 后的结果？把结果与 `x << k` 计算 OR，就是方法一计算的内容。

可以，做法如下：

1. 先通过异或运算，直接去掉 $x$，即 `allOr ^ x`。
2. 如果有多个 $\textit{nums}[i]$ 在同一个比特位上都是 $1$，那么去掉 $x$ 后的其余 $n-1$ 个数的 OR，在这个比特位上也是 $1$。换句话说，无论去掉哪个 $x$，这些比特位**恒为** $1$。用二进制数 $\textit{fixed}$ 记录这些恒为 $1$ 的比特位。
3. 于是，去掉 $x$ 后的其余 $n-1$ 个数的 OR 等于 `(allOr ^ x) | fixed`，表示先直接去掉 $\textit{allOr}$ 中的 $x$，再通过 $\textit{fixed}$ 修正。

如何计算 $\textit{fixed}$？

用「枚举右，维护左」的思想，在遍历 $\textit{nums}$ 的过程中，`allOr & x` 中的 $1$ 必然出现了两次，将其 OR 到 $\textit{fixed}$ 中。

最后，遍历 $\textit{nums}$，设 $x=\textit{nums}[i]$，首先计算 `(allOr ^ x) | fixed`，然后把结果与 `x << k` 计算 OR，更新答案的最大值。

```py [sol-Python3]
class Solution:
    def maximumOr(self, nums: List[int], k: int) -> int:
        all_or = fixed = 0
        for x in nums:
            # 如果在计算 all_or |= x 之前，all_or 和 x 有公共的 1
            # 那就意味着有多个 nums[i] 在这些比特位上都是 1
            fixed |= all_or & x  # 把公共的 1 记录到 fixed 中
            all_or |= x  # 所有数的 OR
        return max((all_or ^ x) | fixed | (x << k) for x in nums)
```

```java [sol-Java]
class Solution {
    public long maximumOr(int[] nums, int k) {
        int allOr = 0;
        int fixed = 0;
        for (int x : nums) {
            // 如果在计算 allOr |= x 之前，allOr 和 x 有公共的 1
            // 那就意味着有多个 nums[i] 在这些比特位上都是 1
            fixed |= allOr & x; // 把公共的 1 记录到 fixed 中
            allOr |= x; // 所有数的 OR
        }

        long ans = 0;
        for (int x : nums) {
            ans = Math.max(ans, (allOr ^ x) | fixed | ((long) x << k));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumOr(vector<int>& nums, int k) {
        int all_or = 0, fixed = 0;
        for (int x : nums) {
            // 如果在计算 all_or |= x 之前，all_or 和 x 有公共的 1
            // 那就意味着有多个 nums[i] 在这些比特位上都是 1
            fixed |= all_or & x; // 把公共的 1 记录到 fixed 中
            all_or |= x; // 所有数的 OR
        }

        long long ans = 0;
        for (int x : nums) {
            ans = max(ans, (all_or ^ x) | fixed | ((long long) x << k));
        }
        return ans;
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

long long maximumOr(int* nums, int numsSize, int k) {
    int all_or = 0, fixed = 0;
    for (int i = 0; i < numsSize; i++) {
        int x = nums[i];
        // 如果在计算 all_or |= x 之前，all_or 和 x 有公共的 1
        // 那就意味着有多个 nums[i] 在这些比特位上都是 1
        fixed |= all_or & x; // 把公共的 1 记录到 fixed 中
        all_or |= x; // 所有数的 OR
    }

    long long ans = 0;
    for (int i = 0; i < numsSize; i++) {
        int x = nums[i];
        ans = MAX(ans, (all_or ^ x) | fixed | ((long long) x << k));
    }
    return ans;
}
```

```go [sol-Go]
func maximumOr(nums []int, k int) int64 {
    allOr, fixed := 0, 0
    for _, x := range nums {
        // 如果在计算 allOr |= x 之前，allOr 和 x 有公共的 1
        // 那就意味着有多个 nums[i] 在这些比特位上都是 1
        fixed |= allOr & x // 把公共的 1 记录到 fixed 中
        allOr |= x // 所有数的 OR
    }

    ans := 0
    for _, x := range nums {
        ans = max(ans, (allOr^x)|fixed|x<<k)
    }
    return int64(ans)
}
```

```js [sol-JavaScript]
var maximumOr = function(nums, k) {
    let allOr = 0, fixed = 0;
    for (const x of nums) {
        // 如果在计算 allOr |= x 之前，allOr 和 x 有公共的 1
        // 那就意味着有多个 nums[i] 在这些比特位上都是 1
        fixed |= allOr & x; // 把公共的 1 记录到 fixed 中
        allOr |= x; // 所有数的 OR
    }

    let ans = 0n;
    for (const x of nums) {
        const res = (BigInt(allOr ^ x) | BigInt(fixed) | (BigInt(x) << BigInt(k)));
        ans = res > ans ? res : ans;
    }
    return Number(ans);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_or(nums: Vec<i32>, k: i32) -> i64 {
        let mut all_or = 0;
        let mut fixed = 0;
        for &x in &nums {
            // 如果在计算 all_or |= x 之前，all_or 和 x 有公共的 1
            // 那就意味着有多个 nums[i] 在这些比特位上都是 1
            fixed |= all_or & x; // 把公共的 1 记录到 fixed 中
            all_or |= x; // 所有数的 OR
        }
        nums.into_iter()
            .map(|x| (all_or ^ x) as i64 | fixed as i64 | ((x as i64) << k))
            .max()
            .unwrap()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面动态规划题单中的「**专题：前后缀分解**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
