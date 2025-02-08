「任意两种糖果价格绝对差的**最小值**」等价于「排序后，任意两种**相邻**糖果价格绝对差的最小值」。

如果题目要求「最大化最小值」或者「最小化最大值」，一般是二分答案。为什么？对于本题来说，甜蜜度越大，能选择的糖果越少，有**单调性**，所以可以二分答案。关于二分的原理，请看[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

定义 $f(d)$ 表示甜蜜度**至少**为 $d$ 时，**最多**能选多少类糖果。

二分答案 $d$：

- 如果 $f(d)\ge k$，说明答案至少为 $d$。
- 如果 $f(d)< k$，说明答案至多为 $d-1$。
- 二分结束后，设答案为 $\textit{d}_0$，那么 $f(d_0)\ge k$ 且 $f(d_0+1)< k$。

如何计算 $f(d)$？对 $\textit{price}$ 从小到大排序，贪心地计算 $f(d)$：

- 第一个数 $\textit{price}[0]$ 一定可以选。
- 假设上一个选的数是 $\textit{pre}$，那么当 $\textit{price}[i] \ge \textit{pre}+d$ 时，才可以选 $\textit{price}[i]$。

## 细节

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的。

- 开区间左端点初始值：$0$。所有糖果都可以选，一定可以满足要求。注意题目保证 $k\le n$。
- 开区间右端点初始值：$\left\lfloor\dfrac{\textit{price}[n-1]-\textit{price}[0]}{k-1}\right\rfloor+1$。假设可以每隔 $d$ 就选一类糖果，那么必须满足 $\textit{price}[0] + (k-1)\cdot d \le \textit{price}[n-1]$，解得 $d\le \left\lfloor\dfrac{\textit{price}[n-1]-\textit{price}[0]}{k-1}\right\rfloor$。所以当 $d=\left\lfloor\dfrac{\textit{price}[n-1]-\textit{price}[0]}{k-1}\right\rfloor+1$ 时，一定无法选 $k$ 类糖果。

## 答疑

**问**：为什么二分出来的答案，一定来自数组中价格的差？有没有可能，二分出来的答案不是任何价格的差？

**答**：反证法。如果答案 $d$ 不是任何价格的差，也就是说，礼盒中任意两种糖果的价格的绝对差都大于 $d$，也就是大于等于 $d+1$。那么对于 $d+1$ 来说，它也满足 $f(d + 1) \ge k$，这与循环不变量相矛盾。所以原命题成立。

```py [sol-Python3]
class Solution:
    def maximumTastiness(self, price: List[int], k: int) -> int:
        def f(d: int) -> int:
            cnt = 1
            pre = price[0]  # 先选一个最小的甜蜜度
            for p in price:
                if p >= pre + d:  # 可以选
                    cnt += 1
                    pre = p  # 上一个选的甜蜜度
            return cnt

        price.sort()
        left = 0
        right = (price[-1] - price[0]) // (k - 1) + 1
        while left + 1 < right:  # 开区间不为空
            # 循环不变量：
            # f(left) >= k
            # f(right) < k
            mid = (left + right) // 2
            if f(mid) >= k:
                left = mid  # 下一轮二分 (mid, right)
            else:
                right = mid  # 下一轮二分 (left, mid)
        return left  # 最大的满足 f(left) >= k 的数
```

```py [sol-Python3 库函数]
class Solution:
    def maximumTastiness(self, price: List[int], k: int) -> int:
        def check(d: int) -> bool:
            # 二分最小的 f(d+1) < k，从而知道最大的 f(d) >= k
            d += 1
            cnt = 1
            pre = price[0]  # 先选一个最小的甜蜜度
            for p in price:
                if p >= pre + d:  # 可以选
                    cnt += 1
                    pre = p  # 上一个选的甜蜜度
            return cnt < k

        price.sort()
        right = (price[-1] - price[0]) // (k - 1)
        return bisect_left(range(right), True, key=check)
```

```java [sol-Java]
class Solution {
    public int maximumTastiness(int[] price, int k) {
        Arrays.sort(price);
        int left = 0;
        int right = (price[price.length - 1] - price[0]) / (k - 1) + 1;
        while (left + 1 < right) { // 开区间不为空
            // 循环不变量：
            // f(left) >= k
            // f(right) < k
            int mid = left + (right - left) / 2;
            if (f(price, mid) >= k) {
                left = mid; // 下一轮二分 (mid, right)
            } else {
                right = mid; // 下一轮二分 (left, mid)
            }
        }
        return left; // 最大的满足 f(left) >= k 的数
    }

    private int f(int[] price, int d) {
        int cnt = 1;
        int pre = price[0]; // 先选一个最小的甜蜜度
        for (int p : price) {
            if (p >= pre + d) { // 可以选
                cnt++;
                pre = p; // 上一个选的甜蜜度
            }
        }
        return cnt;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumTastiness(vector<int>& price, int k) {
        auto f = [&](int d) -> int {
            int cnt = 1, pre = price[0]; // 先选一个最小的甜蜜度
            for (int p : price) {
                if (p >= pre + d) { // 可以选
                    cnt++;
                    pre = p; // 上一个选的甜蜜度
                }
            }
            return cnt;
        };

        ranges::sort(price);
        int left = 0;
        int right = (price.back() - price[0]) / (k - 1) + 1;
        while (left + 1 < right) { // 开区间不为空
            // 循环不变量：
            // f(left) >= k
            // f(right) < k
            int mid = left + (right - left) / 2;
            (f(mid) >= k ? left : right) = mid;
        }
        return left; // 最大的满足 f(left) >= k 的数
    }
};
```

```c [sol-C]
int cmp(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}

int maximumTastiness(int* price, int priceSize, int k) {
    int f(int d) {
        int cnt = 1, pre = price[0]; // 先选一个最小的甜蜜度
        for (int i = 1; i < priceSize; i++) {
            if (price[i] >= pre + d) { // 可以选
                cnt++;
                pre = price[i]; // 上一个选的甜蜜度
            }
        }
        return cnt;
    }

    qsort(price, priceSize, sizeof(int), cmp);
    int left = 0;
    int right = (price[priceSize - 1] - price[0]) / (k - 1) + 1;
    while (left + 1 < right) { // 开区间不为空
        // 循环不变量：
        // f(left) >= k
        // f(right) < k
        int mid = left + (right - left) / 2;
        if (f(mid) >= k) {
            left = mid;
        } else {
            right = mid;
        }
    }
    return left; // 最大的满足 f(left) >= k 的数
}
```

```go [sol-Go]
func maximumTastiness(price []int, k int) int {
    slices.Sort(price)
    return sort.Search((price[len(price)-1]-price[0])/(k-1), func(d int) bool {
        d++ // 二分最小的 f(d+1) < k，从而知道最大的 f(d) >= k
        cnt, pre := 1, price[0]
        for _, p := range price[1:] {
            if p >= pre+d {
                cnt++
                pre = p
            }
        }
        return cnt < k
    })
}
```

```js [sol-JavaScript]
var maximumTastiness = function(price, k) {
    function f(d) {
        let cnt = 1, pre = price[0]; // 先选一个最小的甜蜜度
        for (const p of price) {
            if (p >= pre + d) { // 可以选
                cnt++;
                pre = p; // 上一个选的甜蜜度
            }
        }
        return cnt;
    }

    price.sort((a, b) => a - b);
    let left = 0;
    let right = Math.floor((price[price.length - 1] - price[0]) / (k - 1)) + 1;
    while (left + 1 < right) { // 开区间不为空
        // 循环不变量：
        // f(left) >= k
        // f(right) < k
        const mid = Math.floor((left + right) / 2);
        if (f(mid) >= k) {
            left = mid;
        } else {
            right = mid;
        }
    }
    return left; // 最大的满足 f(left) >= k 的数
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_tastiness(mut price: Vec<i32>, k: i32) -> i32 {
        price.sort_unstable();

        let f = |d: i32| -> i32 {
            let mut cnt = 1;
            let mut pre = price[0]; // 先选一个最小的甜蜜度
            for &p in &price {
                if p >= pre + d { // 可以选
                    cnt += 1;
                    pre = p; // 上一个选的甜蜜度
                }
            }
            cnt
        };

        let mut left = 0;
        let mut right = (price.last().unwrap() - price[0]) / (k - 1) + 1;
        while left + 1 < right { // 开区间不为空
            // 循环不变量：
            // f(left) >= k
            // f(right) < k
            let mid = left + (right - left) / 2;
            if f(mid) >= k {
                left = mid;
            } else {
                right = mid;
            }
        }
        left // 最大的满足 f(left) >= k 的数
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + n\log U)$，其中 $n$ 为 $\textit{price}$ 的长度，$U=\dfrac{\max(\textit{price})-\min(\textit{price})}{k-1}$。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. 【本题相关】[二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
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
