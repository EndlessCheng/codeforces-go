[视频讲解](https://www.bilibili.com/video/BV1i24y1e7E7/) 已出炉，欢迎点赞三连~

---

#### 前置知识 1：二分

见 [【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

#### 前置知识 2&3：前缀和、差分数组

见 [视频讲解](https://www.bilibili.com/video/BV1i24y1e7E7/) 的第四题。

#### 提示 1

看到「最大化最小值」或者「最小化最大值」就要想到**二分答案**，这是一个固定的套路。

为什么？一般来说，二分的值越大，越能/不能满足要求；二分的值越小，越不能/能满足要求，有单调性，可以二分。

类似的题目在先前的周赛中出现过多次，例如：

- [2439. 最小化数组中的最大值](https://leetcode.cn/problems/minimize-maximum-of-array/)
- [2513. 最小化两个数组中的最大值](https://leetcode.cn/problems/minimize-the-maximum-of-two-arrays/)
- [2517. 礼盒的最大甜蜜度](https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/)

#### 提示 2

二分答案 $\textit{minPower}$，从左到右遍历 $\textit{stations}$，如果 $\textit{stations}[i]$ 电量不足 $\textit{minPower}$，那么需要建供电站来补足。

在哪建供电站最好呢？

#### 提示 3

由于 $i$ 左侧的不需要补足，所以贪心地在 $\min(i+r,n-1)$ 处建是最合适的，恰好让 $i$ 在覆盖范围的边界上。

#### 提示 4

设需要建 $m$ 个供电站，那么需要把 $[i,\min(i+2r,n-1)]$ 范围内的电量都增加 $m$。

方法很多，用**差分数组**来更新是最简单的。

最后判断修建的供电站是否超过 $k$，如果超过说明 $\textit{minPower}$ 偏大，否则说明偏小。

> 注：其实前缀和也不需要，可以改为长为 $2r+1$ 的滑动窗口，但这样写有点麻烦，感兴趣的读者可以实现下。

```py [sol1-Python3]
class Solution:
    def maxPower(self, stations: List[int], r: int, k: int) -> int:
        n = len(stations)
        sum = list(accumulate(stations, initial=0))  # 前缀和
        for i in range(n):
            stations[i] = sum[min(i + r + 1, n)] - sum[max(i - r, 0)]  # 电量

        def check(min_power: int) -> bool:
            diff = [0] * n  # 差分数组
            sum_d = need = 0
            for i, power in enumerate(stations):
                sum_d += diff[i]  # 累加差分值
                m = min_power - power - sum_d
                if m > 0:  # 需要 m 个供电站
                    need += m
                    if need > k: return False  # 提前退出这样快一些
                    sum_d += m  # 差分更新
                    if i + r * 2 + 1 < n: diff[i + r * 2 + 1] -= m  # 差分更新
            return True

        left = min(stations)
        right = left + k + 1  # 开区间写法
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid): left = mid
            else: right = mid
        return left
```

```java [sol1-Java]
class Solution {
    public long maxPower(int[] stations, int r, int k) {
        int n = stations.length;
        long[] sum = new long[n + 1]; // 前缀和
        for (int i = 0; i < n; ++i)
            sum[i + 1] = sum[i] + stations[i];
        long mn = Long.MAX_VALUE;
        long[] power = new long[n]; // 电量
        for (int i = 0; i < n; ++i) {
            power[i] = sum[Math.min(i + r + 1, n)] - sum[Math.max(i - r, 0)];
            mn = Math.min(mn, power[i]);
        }

        long left = mn, right = mn + k + 1; // 开区间写法
        while (left + 1 < right) {
            long mid = left + (right - left) / 2;
            if (check(mid, power, n, r, k)) left = mid;
            else right = mid;
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
                if (need > k) return false; // 提前退出这样快一些
                sumD += m; // 差分更新
                if (i + r * 2 + 1 < n) diff[i + r * 2 + 1] -= m; // 差分更新
            }
        }
        return true;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long maxPower(vector<int> &stations, int r, int k) {
        int n = stations.size();
        long sum[n + 1], power[n], diff[n];
        sum[0] = 0;
        for (int i = 0; i < n; ++i)
            sum[i + 1] = sum[i] + stations[i]; // 前缀和
        for (int i = 0; i < n; ++i)
            power[i] = sum[min(i + r + 1, n)] - sum[max(i - r, 0)]; // 电量

        auto check = [&](long min_power) -> bool {
            memset(diff, 0, sizeof(diff)); // 重置差分数组
            long sum_d = 0, need = 0;
            for (int i = 0; i < n; ++i) {
                sum_d += diff[i]; // 累加差分值
                long m = min_power - power[i] - sum_d;
                if (m > 0) { // 需要 m 个供电站
                    need += m;
                    if (need > k) return false; // 提前退出这样快一些
                    sum_d += m; // 差分更新
                    if (i + r * 2 + 1 < n) diff[i + r * 2 + 1] -= m; // 差分更新
                }
            }
            return true;
        };

        long left = *min_element(power, power + n), right = left + k + 1; // 开区间写法
        while (left + 1 < right) {
            long mid = left + (right - left) / 2;
            check(mid) ? left = mid : right = mid;
        }
        return left;
    }
};
```

```go [sol1-Go]
func maxPower(stations []int, r int, k int) int64 {
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

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n\log k)$，其中 $n$ 为 $\textit{stations}$ 的长度。二分需要循环 $O(\log k)$ 次。
- 空间复杂度：$O(n)$。
