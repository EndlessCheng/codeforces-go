## 方法一：二分答案

「最小化最大值」就是二分答案的代名词。我们猜测一个上界 $\textit{limit}$，即要求操作后所有元素均不超过 $\textit{limit}$。由于 $\textit{limit}$ 越大越能够满足，越小越无法满足，有**单调性**，可以二分答案。

从后往前模拟：如果 $\textit{nums}[i]>\textit{limit}$，那么应当去掉多余的 $\textit{extra}=\textit{nums}[i]-\textit{limit}$ 加到 $\textit{nums}[i-1]$ 上，最后如果 $\textit{nums}[0]\le\textit{limit}$，则二分判定成功。

代码实现时可以不用修改 $\textit{nums}$，而是维护 $\textit{extra}$ 变量。

### 细节

开区间二分下界：$\min(\textit{nums})-1$，无法操作。也可以简单地写成 $-1$。

开区间二分上界：$\max(\textit{nums})$，一定可以操作。

附：[视频讲解](https://www.bilibili.com/video/BV1cV4y157BY) 第三题。

```py [sol-Python3]
class Solution:
    def minimizeArrayValue(self, nums: List[int]) -> int:
        def check(limit: int) -> bool:
            extra = 0
            for i in range(len(nums) - 1, 0, -1):
                extra = max(nums[i] + extra - limit, 0)
            return nums[0] + extra <= limit
        return bisect_left(range(max(nums)), True, lo=min(nums), key=check)
```

```java [sol-Java]
class Solution {
    public int minimizeArrayValue(int[] nums) {
        int left = -1;
        int right = 0;
        for (int x : nums) {
            right = Math.max(right, x);
        }
        // 开区间二分，原理见 https://www.bilibili.com/video/BV1AP41137w7/
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            if (check(nums, mid)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(int[] nums, int limit) {
        long extra = 0;
        for (int i = nums.length - 1; i > 0; i--) {
            extra = Math.max(nums[i] + extra - limit, 0);
        }
        return nums[0] + extra <= limit;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimizeArrayValue(vector<int> &nums) {
        auto check = [&](int limit) -> bool {
            long long extra = 0;
            for (int i = nums.size() - 1; i > 0; i--) {
                extra = max(nums[i] + extra - limit, 0LL);
            }
            return nums[0] + extra <= limit;
        };
        // 开区间二分，原理见 https://www.bilibili.com/video/BV1AP41137w7/
        int left = -1, right = ranges::max(nums);
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func minimizeArrayValue(nums []int) int {
	return sort.Search(slices.Max(nums), func(limit int) bool {
		extra := 0
		for i := len(nums) - 1; i > 0; i-- {
			extra = max(nums[i]+extra-limit, 0)
		}
		return nums[0]+extra <= limit
	})
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干变量。

## 方法二：分类讨论

从 $\textit{nums}[0]$ 开始讨论：

- 如果数组中只有 $\textit{nums}[0]$，那么最大值为 $\textit{nums}[0]$。
- 再考虑 $\textit{nums}[1]$，如果 $\textit{nums}[1]\le\textit{nums}[0]$，最大值还是 $\textit{nums}[0]$；否则可以平均这两个数，平均后的最大值为平均值的上取整，即 $\left\lceil\dfrac{\textit{nums}[0]+\textit{nums}[1]}{2}\right\rceil$。
- 再考虑 $\textit{nums}[2]$，如果 $\textit{nums}[2]\le$ 前面算出的最大值，或者这三个数的平均值不超过前面算出的最大值，那么最大值不变；否则可以平均这三个数，做法同上。
- 以此类推直到最后一个数。
- 过程中的最大值为答案。

关于上取整的计算，我们有

$$
\left\lceil\dfrac{a}{b}\right\rceil = \left\lfloor\dfrac{a+b-1}{b}\right\rfloor
$$

讨论 $a$ 被 $b$ 整除，和不被 $b$ 整除两种情况，可以证明上式的正确性。

所以平均值的计算公式为

$$
\left\lceil\dfrac{s}{i+1}\right\rceil = \left\lfloor\dfrac{s+i}{i+1}\right\rfloor
$$

```py [sol-Python3]
class Solution:
    def minimizeArrayValue(self, nums: List[int]) -> int:
        return max((s + i) // (i + 1) for i, s in enumerate(accumulate(nums)))
```

```java [sol-Java]
class Solution {
    public int minimizeArrayValue(int[] nums) {
        long ans = 0;
        long s = 0;
        for (int i = 0; i < nums.length; i++) {
            s += nums[i];
            ans = Math.max(ans, (s + i) / (i + 1));
        }
        return (int) ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimizeArrayValue(vector<int> &nums) {
        long long ans = 0, s = 0;
        for (int i = 0; i < nums.size(); i++) {
            s += nums[i];
            ans = max(ans, (s + i) / (i + 1));
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimizeArrayValue(nums []int) (ans int) {
	s := 0
	for i, x := range nums {
		s += x
		ans = max(ans, (s+i)/(i+1))
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干变量。

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
