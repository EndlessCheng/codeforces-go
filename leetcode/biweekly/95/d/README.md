## 转化

如果存在一种方案，可以让所有城市的电量都 $\ge \textit{low}$，那么也可以 $\ge \textit{low}-1$ 或者更小的值（要求更宽松）。

如果不存在让所有城市的电量都 $\ge \textit{low}$ 的方案，那么也不存在 $\ge \textit{low}+1$ 或者更大的方案（要求更苛刻）。

据此，可以**二分猜答案**。关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

现在问题转化成一个判定性问题：

- 给定 $\textit{low}$，是否存在建造 $k$ 座额外电站的方案，使得所有城市的电量都 $\ge \textit{low}$？

如果存在，说明答案 $\ge \textit{low}$，否则答案 $<\textit{low}$。

## 思路

由于已经建造的电站是可以发电的，我们需要在二分之前，用 $\textit{stations}$ 计算每个城市的**初始电量** $\textit{power}$。这可以用**前缀和**或者**滑动窗口**做，具体后面解释。

然后从左到右遍历 $\textit{power}$，挨个处理。如果 $\textit{power}[i] < \textit{low}$，就需要建造电站，额外提供 $\textit{low} - \textit{power}[i]$ 的电力。

在哪建造电站最好呢？

由于我们是从左到右遍历的，在 $i$ 左侧的城市已经处理好了，所以建造的电站越靠右越好，**尽可能多地覆盖没遍历到的城市**。具体地，$i$ 应当恰好在电站供电范围的边缘上，也就是把电站建在 $i+r$ 的位置，使得电站覆盖范围为 $[i,i+2r]$。

我们要建 $m = \textit{low} - \textit{power}[i]$ 个供电站，也就是把下标在 $[i,i+2r]$ 中的城市的电量都增加 $m$。

这里有一个「区间增加 $m$」的需求，用**差分数组**实现，原理见 [差分数组原理讲解](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)。

我们要一边做差分更新，一边计算差分数组的前缀和，以得到当前城市的实际电量。

遍历的同时，累计额外建造的电站数量，如果超过 $k$，不满足要求，可以提前跳出循环。

## 细节

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的，喜欢哪种写法就用哪种。

- 开区间左端点初始值：$\min(\textit{power}) + \left\lfloor\dfrac{k}{n}\right\rfloor$。把 $k$ 均摊，即使 $r=0$，每个城市都能至少分到 $\left\lfloor\dfrac{k}{n}\right\rfloor$ 的额外电量，所以 $\textit{low} = \min(\textit{power}) + \left\lfloor\dfrac{k}{n}\right\rfloor$ 时一定满足要求。
- 开区间右端点初始值：$\min(\textit{power}) + k + 1$。即使把所有额外电站都建给电量最小的城市，也无法满足要求。

> 对于开区间写法，简单来说 `check(mid) == true` 时更新的是谁，最后就返回谁。相比其他二分写法，开区间写法不需要思考加一减一等细节，更简单。推荐使用开区间写二分。

## 写法一：前缀和

能覆盖城市 $i$ 的电站下标范围是 $[i-r,i+r]$。注意下标不能越界，所以实际范围是 $[\max(i-r,0),\min(i+r,n-1)]$。

我们需要计算 $\textit{stations}$ 的这个范围（子数组）的和。

计算 $\textit{stations}$ 的 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 数组后，可以 $\mathcal{O}(1)$ 计算 $\textit{stations}$ 的任意子数组的和。

```py [sol-Python3]
class Solution:
    def maxPower(self, stations: List[int], r: int, k: int) -> int:
        n = len(stations)
        # 前缀和
        s = list(accumulate(stations, initial=0))
        # 初始电量
        power = [s[min(i + r + 1, n)] - s[max(i - r, 0)] for i in range(n)]

        def check(low: int) -> bool:
            diff = [0] * n  # 差分数组
            sum_d = built = 0
            for i, p in enumerate(power):
                sum_d += diff[i]  # 累加差分值
                m = low - (p + sum_d)
                if m <= 0:
                    continue
                # 需要在 i+r 额外建造 m 个供电站
                built += m
                if built > k:  # 不满足要求
                    return False
                # 把区间 [i, i+2r] 加一
                sum_d += m  # 由于 diff[i] 后面不会再访问，我们直接加到 sum_d 中
                if (right := i + r * 2 + 1) < n:
                    diff[right] -= m
            return True

        # 开区间二分
        mn = min(power)
        left, right = mn + k // n, mn + k + 1
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
    def maxPower(self, stations: List[int], r: int, k: int) -> int:
        n = len(stations)
        # 前缀和
        s = list(accumulate(stations, initial=0))
        # 初始电量
        power = [s[min(i + r + 1, n)] - s[max(i - r, 0)] for i in range(n)]

        def check(low: int) -> bool:
            low += 1  # 二分最小的不满足要求的 low（符合库函数），这样最终返回的就是最大的满足要求的 low
            diff = [0] * n  # 差分数组
            sum_d = built = 0
            for i, p in enumerate(power):
                sum_d += diff[i]  # 累加差分值
                m = low - (p + sum_d)
                if m <= 0:
                    continue
                # 需要在 i+r 额外建造 m 个供电站
                built += m
                if built > k:  # 不满足要求
                    return True
                # 把区间 [i, i+2r] 加一
                sum_d += m  # 由于 diff[i] 后面不会再访问，我们直接加到 sum_d 中
                if (right := i + r * 2 + 1) < n:
                    diff[right] -= m
            return False

        mn = min(power)
        left, right = mn + k // n, mn + k
        return bisect_left(range(right), True, lo=left, key=check)
```

```java [sol-Java]
class Solution {
    public long maxPower(int[] stations, int r, int k) {
        int n = stations.length;
        // 前缀和
        long[] sum = new long[n + 1];
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + stations[i];
        }

        // 初始电量
        long[] power = new long[n];
        long mn = Long.MAX_VALUE;
        for (int i = 0; i < n; i++) {
            power[i] = sum[Math.min(i + r + 1, n)] - sum[Math.max(i - r, 0)];
            mn = Math.min(mn, power[i]);
        }

        // 开区间二分
        long left = mn + k / n;
        long right = mn + k + 1;
        while (left + 1 < right) {
            long mid = left + (right - left) / 2;
            if (check(mid, power, r, k)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }

    private boolean check(long low, long[] power, int r, int k) {
        int n = power.length;
        long[] diff = new long[n + 1];
        long sumD = 0;
        long built = 0;
        for (int i = 0; i < n; i++) {
            sumD += diff[i]; // 累加差分值
            long m = low - (power[i] + sumD);
            if (m <= 0) {
                continue;
            }
            // 需要在 i+r 额外建造 m 个供电站
            built += m;
            if (built > k) { // 不满足要求
                return false;
            }
            // 把区间 [i, i+2r] 加一
            sumD += m; // 由于 diff[i] 后面不会再访问，我们直接加到 sumD 中
            diff[Math.min(i + r * 2 + 1, n)] -= m;
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
        // 前缀和
        vector<long long> sum(n + 1);
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + stations[i];
        }

        // 初始电量
        vector<long long> power(n);
        long long mn = LLONG_MAX;
        for (int i = 0; i < n; i++) {
            power[i] = sum[min(i + r + 1, n)] - sum[max(i - r, 0)];
            mn = min(mn, power[i]);
        }

        auto check = [&](long long low) -> bool {
            vector<long long> diff(n + 1);
            long long sum_d = 0, built = 0;
            for (int i = 0; i < n; i++) {
                sum_d += diff[i]; // 累加差分值
                long long m = low - (power[i] + sum_d);
                if (m <= 0) {
                    continue;
                }
                // 需要在 i+r 额外建造 m 个供电站
                built += m;
                if (built > k) { // 不满足要求
                    return false;
                }
                // 把区间 [i, i+2r] 加一
                sum_d += m; // 由于 diff[i] 后面不会再访问，我们直接加到 sum_d 中
                diff[min(i + r * 2 + 1, n)] -= m;
            }
            return true;
        };

        // 开区间二分
        long long left = mn + k / n, right = mn + k + 1;
        while (left + 1 < right) {
            long long mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left;
    }
};
```

```go [sol-Go]
func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	// 前缀和
	sum := make([]int, n+1)
	for i, x := range stations {
		sum[i+1] = sum[i] + x
	}

	// 初始电量
	power := make([]int, n)
	mn := math.MaxInt
	for i := range power {
		power[i] = sum[min(i+r+1, n)] - sum[max(i-r, 0)]
		mn = min(mn, power[i])
	}

	// 二分答案
	left := mn + k/n
	right := mn + k
	ans := left + sort.Search(right-left, func(low int) bool {
		// 这里 +1 是为了二分最小的不满足要求的 low（符合库函数），这样最终返回的就是最大的满足要求的 low
		low += left + 1
		diff := make([]int, n+1) // 差分数组
		sumD, built := 0, 0
		for i, p := range power {
			sumD += diff[i] // 累加差分值
			m := low - (p + sumD)
			if m <= 0 {
				continue
			}
			// 需要在 i+r 额外建造 m 个供电站
			built += m
			if built > k { // 不满足要求
				return true
			}
			// 把区间 [i, i+2r] 加一
			sumD += m // 由于 diff[i] 后面不会再访问，我们直接加到 sumD 中
			diff[min(i+r*2+1, n)] -= m
		}
		return false
	})
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log k)$，其中 $n$ 是 $\textit{stations}$ 的长度。二分 $\mathcal{O}(\log k)$ 次，每次 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法二：滑动窗口

用 [滑动窗口](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/2809359/tao-lu-jiao-ni-jie-jue-ding-chang-hua-ch-fzfo/) 计算 $\textit{power}$。

```py [sol-Python3]
class Solution:
    def maxPower(self, stations: List[int], r: int, k: int) -> int:
        n = len(stations)
        # 滑动窗口
        s = sum(stations[:r])  # 先计算 [0, r-1] 的发电量，为第一个窗口做准备
        power = [0] * n
        for i in range(n):
            # 右边进
            if (right := i + r) < n:
                s += stations[right]
            # 左边出
            if (left := i - r - 1) >= 0:
                s -= stations[left]
            power[i] = s

        def check(low: int) -> bool:
            diff = [0] * n  # 差分数组
            sum_d = built = 0
            for i, p in enumerate(power):
                sum_d += diff[i]  # 累加差分值
                m = low - (p + sum_d)
                if m <= 0:
                    continue
                # 需要在 i+r 额外建造 m 个供电站
                built += m
                if built > k:  # 不满足要求
                    return False
                # 把区间 [i, i+2r] 加一
                sum_d += m  # 由于 diff[i] 后面不会再访问，我们直接加到 sum_d 中
                if (right := i + r * 2 + 1) < n:
                    diff[right] -= m
            return True

        # 开区间二分
        mn = min(power)
        left, right = mn + k // n, mn + k + 1
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
    def maxPower(self, stations: List[int], r: int, k: int) -> int:
        n = len(stations)
        # 滑动窗口
        s = sum(stations[:r])  # 先计算 [0, r-1] 的发电量，为第一个窗口做准备
        power = [0] * n
        for i in range(n):
            # 右边进
            if (right := i + r) < n:
                s += stations[right]
            # 左边出
            if (left := i - r - 1) >= 0:
                s -= stations[left]
            power[i] = s

        def check(low: int) -> bool:
            low += 1  # 二分最小的不满足要求的 low（符合库函数），这样最终返回的就是最大的满足要求的 low
            diff = [0] * n  # 差分数组
            sum_d = built = 0
            for i, p in enumerate(power):
                sum_d += diff[i]  # 累加差分值
                m = low - (p + sum_d)
                if m <= 0:
                    continue
                # 需要在 i+r 额外建造 m 个供电站
                built += m
                if built > k:  # 不满足要求
                    return True
                # 把区间 [i, i+2r] 加一
                sum_d += m  # 由于 diff[i] 后面不会再访问，我们直接加到 sum_d 中
                if (right := i + r * 2 + 1) < n:
                    diff[right] -= m
            return False

        mn = min(power)
        left, right = mn + k // n, mn + k
        return bisect_left(range(right), True, lo=left, key=check)
```

```java [sol-Java]
class Solution {
    public long maxPower(int[] stations, int r, int k) {
        int n = stations.length;
        // 滑动窗口
        // 先计算 [0, r-1] 的发电量，为第一个窗口做准备
        long sum = 0;
        for (int i = 0; i < r; i++) {
            sum += stations[i];
        }
        long[] power = new long[n];
        long mn = Long.MAX_VALUE;
        for (int i = 0; i < n; i++) {
            // 右边进
            if (i + r < n) {
                sum += stations[i + r];
            }
            // 左边出
            if (i - r - 1 >= 0) {
                sum -= stations[i - r - 1];
            }
            power[i] = sum;
            mn = Math.min(mn, sum);
        }

        // 开区间二分
        long left = mn + k / n;
        long right = mn + k + 1;
        while (left + 1 < right) {
            long mid = left + (right - left) / 2;
            if (check(mid, power, r, k)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }

    private boolean check(long low, long[] power, int r, int k) {
        int n = power.length;
        long[] diff = new long[n + 1];
        long sumD = 0;
        long built = 0;
        for (int i = 0; i < n; i++) {
            sumD += diff[i]; // 累加差分值
            long m = low - (power[i] + sumD);
            if (m <= 0) {
                continue;
            }
            // 需要在 i+r 额外建造 m 个供电站
            built += m;
            if (built > k) { // 不满足要求
                return false;
            }
            // 把区间 [i, i+2r] 加一
            sumD += m; // 由于 diff[i] 后面不会再访问，我们直接加到 sumD 中
            diff[Math.min(i + r * 2 + 1, n)] -= m;
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
        // 滑动窗口
        // 先计算 [0, r-1] 的发电量，为第一个窗口做准备
        long long sum = reduce(stations.begin(), stations.begin() + r, 0LL);
        vector<long long> power(n);
        long long mn = LLONG_MAX;
        for (int i = 0; i < n; i++) {
            // 右边进
            if (i + r < n) {
                sum += stations[i + r];
            }
            // 左边出
            if (i - r - 1 >= 0) {
                sum -= stations[i - r - 1];
            }
            power[i] = sum;
            mn = min(mn, sum);
        }

        auto check = [&](long long low) -> bool {
            vector<long long> diff(n + 1);
            long long sum_d = 0, built = 0;
            for (int i = 0; i < n; i++) {
                sum_d += diff[i]; // 累加差分值
                long long m = low - (power[i] + sum_d);
                if (m <= 0) {
                    continue;
                }
                // 需要在 i+r 额外建造 m 个供电站
                built += m;
                if (built > k) { // 不满足要求
                    return false;
                }
                // 把区间 [i, i+2r] 加一
                sum_d += m; // 由于 diff[i] 后面不会再访问，我们直接加到 sum_d 中
                diff[min(i + r * 2 + 1, n)] -= m;
            }
            return true;
        };

        // 开区间二分
        long long left = mn + k / n, right = mn + k + 1;
        while (left + 1 < right) {
            long long mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left;
    }
};
```

```go [sol-Go]
func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	// 滑动窗口
	// 先计算 [0, r-1] 的发电量，为第一个窗口做准备
	sum := 0
	for _, x := range stations[:r] {
		sum += x
	}
	power := make([]int, n)
	mn := math.MaxInt
	for i := range power {
		// 右边进
		if i+r < n {
			sum += stations[i+r]
		}
		// 左边出
		if i-r-1 >= 0 {
			sum -= stations[i-r-1]
		}
		power[i] = sum
		mn = min(mn, sum)
	}

	// 二分答案
	left := mn + k/n
	right := mn + k
	ans := left + sort.Search(right-left, func(low int) bool {
		// 这里 +1 是为了二分最小的不满足要求的 low（符合库函数），这样最终返回的就是最大的满足要求的 low
		low += left + 1
		diff := make([]int, n+1) // 差分数组
		sumD, built := 0, 0
		for i, p := range power {
			sumD += diff[i] // 累加差分值
			m := low - (p + sumD)
			if m <= 0 {
				continue
			}
			// 需要在 i+r 额外建造 m 个供电站
			built += m
			if built > k { // 不满足要求
				return true
			}
			// 把区间 [i, i+2r] 加一
			sumD += m // 由于 diff[i] 后面不会再访问，我们直接加到 sumD 中
			diff[min(i+r*2+1, n)] -= m
		}
		return false
	})
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log k)$，其中 $n$ 是 $\textit{stations}$ 的长度。二分 $\mathcal{O}(\log k)$ 次，每次 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
