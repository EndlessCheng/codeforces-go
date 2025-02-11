假设 $\textit{gameScore}$ 中的每个数都**至少**为 $\textit{low}$，那么 $\textit{low}$ 越大，操作次数就越多，有单调性，可以**二分答案**。关于二分的原理，请看视频 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

在计算前，还需要说明一个转换关系：我们可以把任何一种或长或短的、来来回回的移动方式，转换成若干组「左右横跳」，也就是先在 $0$ 和 $1$ 之间左右横跳，然后在 $1$ 和 $2$ 之间左右横跳，在 $2$ 和 $3$ 之间左右横跳……直到最终位置为 $n-1$ 或者 $n-2$。如下图：

![lc3449-3-c.png](https://pic.leetcode.cn/1739098814-MZuCoO-lc3449-3-c.png)

从第一个数开始计算。设 $p=\textit{points}[0]$，至少要增加 $k=\left\lceil\dfrac{\textit{low}}{p}\right\rceil$ 次。

第一次操作需要从 $-1$ 走到 $0$，后面的 $k-1$ 次增加可以在 $0$ 和 $1$ 之间左右横跳。

所以一共需要

$$
2(k-1)+1 = 2k-1
$$

次操作。

注意这会导致下一个数已经操作了 $k-1$ 次。

如此循环，直到最后一个数。如果循环中发现操作次数已经超过 $m$，退出循环。

注意，如果最后一个数还需要操作的次数 $\le 0$，那么是不需要继续操作的，退出循环。

## 细节

### 1)

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的。

- 开区间左端点初始值：$0$。无需操作，一定可以满足要求。
- 开区间右端点初始值：$\left\lceil\dfrac{m}{2}\right\rceil\cdot \min(\textit{points})+1$。假设第一个数是最小值，那么它可以通过左右横跳操作 $\left\lceil\dfrac{m}{2}\right\rceil$ 次。结果 $+1$ 之后一定无法满足要求。

### 2)

关于上取整的计算，当 $a$ 和 $b$ 均为正整数时，我们有

$$
\left\lceil\dfrac{a}{b}\right\rceil = \left\lfloor\dfrac{a-1}{b}\right\rfloor + 1
$$

讨论 $a$ 被 $b$ 整除，和不被 $b$ 整除两种情况，可以证明上式的正确性。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1ekN2ebEHx/?t=50m34s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxScore(self, points: List[int], m: int) -> int:
        def check(low: int) -> bool:
            n = len(points)
            rem = m
            pre = 0
            for i, p in enumerate(points):
                k = (low - 1) // p + 1 - pre  # 还需要操作的次数
                if i == n - 1 and k <= 0:  # 最后一个数已经满足要求
                    break
                if k < 1:
                    k = 1  # 至少要走 1 步
                rem -= k * 2 - 1  # 左右横跳
                if rem < 0:
                    return False
                pre = k - 1  # 右边那个数顺带操作了 k-1 次
            return True

        left = 0
        right = (m + 1) // 2 * min(points) + 1
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid):
                left = mid
            else:
                right = mid
        return left
```

```py [sol-Python3 库函数]
class Solution:
    def maxScore(self, points: List[int], m: int) -> int:
        def check(low: int) -> bool:
            # 二分最小的不满足要求的 low+1，即可得到最大的满足要求的 low
            n = len(points)
            rem = m
            pre = 0
            for i, p in enumerate(points):
                k = low // p + 1 - pre  # 还需要操作的次数
                if i == n - 1 and k <= 0:  # 最后一个数已经满足要求
                    break
                if k < 1:
                    k = 1  # 至少要走 1 步
                rem -= k * 2 - 1  # 左右横跳
                if rem < 0:
                    return True  # 取反
                pre = k - 1  # 右边那个数顺带操作了 k-1 次
            return False  # 取反

        right = (m + 1) // 2 * min(points)
        return bisect_left(range(right), True, key=check)
```

```java [sol-Java]
class Solution {
    public long maxScore(int[] points, int m) {
        int mn = Integer.MAX_VALUE;
        for (int p : points) {
            mn = Math.min(mn, p);
        }
        long left = 0;
        long right = (long) (m + 1) / 2 * mn + 1;
        while (left + 1 < right) {
            long mid = (left + right) >>> 1;
            if (check(mid, points, m)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }

    private boolean check(long low, int[] points, int m) {
        int n = points.length;
        int left = m; // 剩余操作次数
        int pre = 0;
        for (int i = 0; i < n; i++) {
            int k = (int) ((low - 1) / points[i]) + 1 - pre; // 还需要操作的次数
            if (i == n - 1 && k <= 0) { // 最后一个数已经满足要求
                break;
            }
            k = Math.max(k, 1); // 至少要走 1 步
            left -= k * 2 - 1; // 左右横跳
            if (left < 0) {
                return false;
            }
            pre = k - 1; // 右边那个数顺带操作了 k-1 次
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxScore(vector<int>& points, int m) {
        auto check = [&](long long low) -> bool {
            int n = points.size(), rem = m, pre = 0;
            for (int i = 0; i < n; i++) {
                int k = (low - 1) / points[i] + 1 - pre; // 还需要操作的次数
                if (i == n - 1 && k <= 0) { // 最后一个数已经满足要求
                    break;
                }
                k = max(k, 1); // 至少要走 1 步
                rem -= k * 2 - 1; // 左右横跳
                if (rem < 0) {
                    return false;
                }
                pre = k - 1; // 右边那个数顺带操作了 k-1 次
            }
            return true;
        };

        long long left = 0;
        long long right = 1LL * (m + 1) / 2 * ranges::min(points) + 1;
        while (left + 1 < right) {
            long long mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left;
    }
};
```

```go [sol-Go]
func maxScore(points []int, m int) int64 {
	right := (m + 1) / 2 * slices.Min(points)
	ans := sort.Search(right, func(low int) bool {
		// 二分最小的不满足要求的 low+1，即可得到最大的满足要求的 low
		low++
		left := m
		pre := 0
		for i, p := range points {
			k := (low-1)/p + 1 - pre // 还需要操作的次数
			if i == len(points)-1 && k <= 0 { // 最后一个数已经满足要求
				break
			}
			k = max(k, 1) // 至少要走 1 步
			left -= k*2 - 1 // 左右横跳
			if left < 0 {
				return true
			}
			pre = k - 1 // 右边那个数顺带操作了 k-1 次
		}
		return false
	})
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{points}$ 的长度，$U=\min(points)\cdot m$。
- 空间复杂度：$\mathcal{O}(1)$。

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
10. 【本题相关】[贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
