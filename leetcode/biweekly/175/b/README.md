## 转化

对于不等式

$$
\text{nonPositive}(\textit{nums}, k) \le k^2
$$

当 $k$ 逐渐增大时，操作次数会变小（或者不变），所以 $\text{nonPositive}(\textit{nums}, k)$ 会变小（或者不变）；另一方面，$k^2$ 会随着 $k$ 的增大而增大。

所以当 $k$ 较小时，不等式不成立；当 $k$ 较大时，不等式成立。

据此，可以**二分猜答案**。关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

现在问题转化成一个判定性问题：

- 给定 $k$，计算每个数的操作次数，判断不等式是否成立。

如果成立，说明答案 $\le k$，否则答案 $> k$。

## 思路

对于 $x = \textit{nums}[i]$，设需要操作 $t$ 次，那么有

$$
k\cdot t \ge x
$$

解得

$$
t\ge \dfrac{x}{k}
$$

由于 $t$ 是整数，所以最小操作次数为

$$
\left\lceil\dfrac{x}{k}\right\rceil
$$

所以有

$$
\text{nonPositive}(\textit{nums}, k) = \sum_{i=0}^{n-1}\left\lceil\dfrac{\textit{nums}[i]}{k}\right\rceil
$$

## 细节

### 1)

下面代码采用开区间二分。使用闭区间或者半闭半开区间也是可以的，喜欢哪种写法就用哪种。

- 开区间左端点初始值：$0$。无法满足题目要求。
- 开区间左端点初始值（优化）：$\left\lceil\sqrt n\right\rceil - 1$。由于 $\textit{nums}$ 中的元素都是正数，每个数都至少要操作一次，所以 $\text{nonPositive}(\textit{nums}, k) \ge n$，所以 $k$ 必须满足 $k^2\ge n$，即 $k\ge \left\lceil\sqrt n\right\rceil$。减一后，一定无法满足题目要求。
- 开区间右端点初始值：$M$，其中 $M = \max(\textit{nums})$。此时 $\text{nonPositive}(\textit{nums}, M)=n$。如果 $n \le M^2$，那么满足题目要求。这引出了一个**特殊情况**：如果 $M\le \left\lceil\sqrt n\right\rceil$，那么答案就是理论最小值 $\left\lceil\sqrt n\right\rceil$，此时 $\text{nonPositive}(\textit{nums}, k) \le k^2$ 为 $n\le \left\lceil\sqrt n\right\rceil^2$，一定成立，可以提前返回 $\left\lceil\sqrt n\right\rceil$，无需二分。
- 开区间右端点初始值（优化）：$\left\lceil\sqrt[3] {2nM}\right\rceil$。最坏情况下，$\textit{nums}$ 中的元素都是 $M$，一共需要操作 $\left\lceil\dfrac{M}{k}\right\rceil n$ 次。当 $k\le M$ 时，有 $\left\lceil\dfrac{M}{k}\right\rceil n\le \dfrac{2M}{k}n$，所以当 $\dfrac{2M}{k}n\le k^2$，即 $k \ge \left\lceil\sqrt[3] {2nM}\right\rceil$ 时，一定满足题目要求。

> **注**：对于开区间写法，简单来说 `check(mid) == true` 时更新的是谁，最后就返回谁。相比其他二分写法，开区间写法不需要思考加一减一等细节，更简单。推荐使用开区间写二分。

### 2)

关于上取整的计算，当 $a$ 为整数，$b$ 为正整数时，有恒等式

$$
\left\lceil\dfrac{a}{b}\right\rceil = \left\lfloor\dfrac{a+b-1}{b}\right\rfloor = \left\lfloor\dfrac{a-1}{b}\right\rfloor + 1
$$

见 [上取整下取整转换公式的证明](https://zhuanlan.zhihu.com/p/1890356682149838951)。

所以

$$
\text{nonPositive}(\textit{nums}, k) = n + \sum_{i=0}^{n-1}\left\lfloor\dfrac{\textit{nums}[i]-1}{k}\right\rfloor
$$

这样做可以避免浮点运算，避免浮点数的舍入误差导致计算错误。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def minimumK(self, nums: List[int]) -> int:
        n = len(nums)

        def check(k: int) -> bool:
            return n + sum((x - 1) // k for x in nums) <= k * k

        left = ceil(sqrt(n))  # 答案的下界
        right = ceil(cbrt(n * max(nums) * 2))  # 答案的上界
        return bisect_left(range(right), True, left, key=check)
```

```py [sol-Python3 手写二分]
class Solution:
    def minimumK(self, nums: List[int]) -> int:
        n = len(nums)
        mx = max(nums)
        rt = ceil(sqrt(n))  # 答案的下界
        if mx <= rt:
            return rt

        def check(k: int) -> bool:
            return n + sum((x - 1) // k for x in nums) <= k * k

        left = rt - 1
        right = ceil(cbrt(n * mx * 2))  # 答案的上界
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid):
                right = mid
            else:
                left = mid
        return right
```

```java [sol-Java]
class Solution {
    public int minimumK(int[] nums) {
        int n = nums.length;
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int rt = (int) Math.ceil(Math.sqrt(n)); // 答案的下界
        if (mx <= rt) {
            return rt;
        }

        int left = rt - 1;
        int right = (int) Math.ceil(Math.cbrt((long) n * mx * 2)); // 答案的上界
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            if (check(mid, nums)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(int k, int[] nums) {
        long sum = nums.length;
        for (int x : nums) {
            sum += (x - 1) / k;
        }
        return sum <= (long) k * k;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumK(vector<int>& nums) {
        int n = nums.size();
        int mx = ranges::max(nums);
        int rt = ceil(sqrt(n)); // 答案的下界
        if (mx <= rt) {
            return rt;
        }

        auto check = [&](int k) -> bool {
            long long sum = n;
            for (int x : nums) {
                sum += (x - 1) / k;
            }
            return sum <= 1LL * k * k;
        };

        int left = rt - 1;
        int right = ceil(cbrt(1LL * n * mx * 2)); // 答案的上界
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func minimumK(nums []int) int {
	n := len(nums)
	mx := slices.Max(nums)
	left := int(math.Ceil(math.Sqrt(float64(n))))           // 答案的下界
	right := int(math.Ceil(math.Cbrt(float64(n * mx * 2)))) // 答案的上界
	ans := left + sort.Search(right-left, func(k int) bool {
		k += left
		sum := n
		for _, x := range nums {
			sum += (x - 1) / k
		}
		return sum <= k*k
	})
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log (\sqrt[3] {nU} - \sqrt n))$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面二分题单的「**§2.1 求最小**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
