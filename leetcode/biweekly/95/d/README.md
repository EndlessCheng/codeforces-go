### 前置知识

1. 二分：见 [【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。
2. 前缀和、差分数组：见本题 [视频讲解](https://www.bilibili.com/video/BV1i24y1e7E7/) 的第四题。

### 提示 1

看到「最大化最小值」或者「最小化最大值」就要想到**二分答案**，这是一个固定的套路。

为什么？一般来说，二分的值越大，越能/不能满足要求；二分的值越小，越不能/能满足要求，有单调性，可以二分。

### 提示 2

二分答案 $\textit{minPower}$，从左到右遍历 $\textit{stations}$，如果 $\textit{stations}[i]$ 电量不足 $\textit{minPower}$，那么需要建供电站来补足。

在哪建供电站最好呢？

### 提示 3

由于 $i$ 左侧的不需要补足，所以贪心地在 $\min(i+r,n-1)$ 处建是最合适的，恰好让 $i$ 在覆盖范围的边界上。

### 提示 4

设需要建 $m$ 个供电站，那么需要把 $[i,\min(i+2r,n-1)]$ 范围内的电量都增加 $m$。

方法很多，用**差分数组**来更新是最简单的。

最后判断修建的供电站是否超过 $k$，如果超过说明 $\textit{minPower}$ 偏大，否则说明偏小。

> 注：其实前缀和也不需要，可以改为长为 $2r+1$ 的滑动窗口，但这样写有点麻烦，感兴趣的读者可以实现下。

```py [sol-Python3]
class Solution:
    def maxPower(self, stations: List[int], r: int, k: int) -> int:
        n = len(stations)
        s = list(accumulate(stations, initial=0))  # 前缀和
        for i in range(n):
            stations[i] = s[min(i + r + 1, n)] - s[max(i - r, 0)]  # 电量

        def check(min_power: int) -> bool:
            diff = [0] * n  # 差分数组
            sum_d = need = 0
            for i, power in enumerate(stations):
                sum_d += diff[i]  # 累加差分值
                m = min_power - power - sum_d
                if m > 0:  # 需要 m 个供电站
                    need += m
                    if need > k:
                        return False  # 提前退出这样快一些
                    sum_d += m  # 差分更新
                    if i + r * 2 + 1 < n:
                        diff[i + r * 2 + 1] -= m  # 差分更新
            return True

        left = min(stations)
        right = left + k + 1  # 开区间写法
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid):
                left = mid
            else:
                right = mid
        return left
```

```java [sol-Java]
class Solution {
    public long maxPower(int[] stations, int r, int k) {
        int n = stations.length;
        long[] sum = new long[n + 1]; // 前缀和
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + stations[i];
        }
        long mn = Long.MAX_VALUE;
        long[] power = new long[n]; // 电量
        for (int i = 0; i < n; i++) {
            power[i] = sum[Math.min(i + r + 1, n)] - sum[Math.max(i - r, 0)];
            mn = Math.min(mn, power[i]);
        }

        long left = mn;
        long right = mn + k + 1; // 开区间写法
        while (left + 1 < right) {
            long mid = left + (right - left) / 2;
            if (check(mid, power, n, r, k)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }

    private boolean check(long minPower, long[] power, int n, int r, int k) {
        long[] diff = new long[n + 1]; // 差分数组
        long sumD = 0, need = 0;
        for (int i = 0; i < n; ++i) {
            sumD += diff[i]; // 累加差分值
            long m = minPower - power[i] - sumD;
            if (m > 0) { // 需要 m 个供电站
                need += m;
                if (need > k) {
                    return false; // 提前退出这样快一些
                }
                sumD += m; // 差分更新
                if (i + r * 2 + 1 < n) {
                    diff[i + r * 2 + 1] -= m; // 差分更新
                }
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxPower(vector<int>& stations, int r, int k) {
        int n = stations.size();
        vector<long long> sum(n + 1), power(n), diff(n);
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + stations[i]; // 前缀和
        }
        for (int i = 0; i < n; i++) {
            power[i] = sum[min(i + r + 1, n)] - sum[max(i - r, 0)]; // 电量
        }

        auto check = [&](long long min_power) -> bool {
            ranges::fill(diff, 0); // 重置差分数组
            long long sum_d = 0, need = 0;
            for (int i = 0; i < n; ++i) {
                sum_d += diff[i]; // 累加差分值
                long long m = min_power - power[i] - sum_d;
                if (m > 0) { // 需要 m 个供电站
                    need += m;
                    if (need > k) {
                        return false; // 提前退出这样快一些
                    }
                    sum_d += m; // 差分更新
                    if (i + r * 2 + 1 < n) {
                        diff[i + r * 2 + 1] -= m; // 差分更新
                    }
                }
            }
            return true;
        };

        long long left = ranges::min(power);
        long long right = left + k + 1; // 开区间写法
        while (left + 1 < right) {
            long long mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left;
    }
};
```

```go [sol-Go]
func maxPower(stations []int, r, k int) int64 {
    n := len(stations)
    sum := make([]int, n+1) // 前缀和
    for i, x := range stations {
        sum[i+1] = sum[i] + x
    }
    mn := math.MaxInt
    for i := range stations {
        stations[i] = sum[min(i+r+1, n)] - sum[max(i-r, 0)] // 电量
        mn = min(mn, stations[i])
    }
    return int64(mn + sort.Search(k, func(minPower int) bool {
        minPower += mn + 1 // 改为二分最小的不满足要求的值，这样 sort.Search 返回的就是最大的满足要求的值
        diff := make([]int, n) // 差分数组
        sumD, need := 0, 0
        for i, power := range stations {
            sumD += diff[i] // 累加差分值
            m := minPower - power - sumD
            if m > 0 { // 需要 m 个供电站
                need += m
                if need > k { // 提前退出这样快一些
                    return true // 不满足要求
                }
                sumD += m // 差分更新
                if i+r*2+1 < n {
                    diff[i+r*2+1] -= m // 差分更新
                }
            }
        }
        return false
    }))
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log k)$，其中 $n$ 为 $\textit{stations}$ 的长度。二分需要循环 $O(\log k)$ 次。
- 空间复杂度：$O(n)$。

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
