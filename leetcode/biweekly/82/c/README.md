下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

根据题意，在 $\textit{nums}_1[i]$ 上 $+1$，等价于在 $\textit{nums}_2[i]$ 上 $-1$，反之亦然。

定义 $a[i]=|\textit{nums}_1[i]-\textit{nums}_2[i]|$，$k=k_1+k_2$，则原问题可以转换成：

> 对数组 $a$ 执行至多 $k$ 次 $-1$ 操作，能得到的 $\sum a[i]^2$ 的最小值。

对于两个数，先把大的 $-1$ 会更优。我们可以将 $a$ 从大往小排序，然后从左到右遍历 $a$，同时更新剩余操作次数 $k$。

当遍历至 $a[i]$ 时，此时从 $a[0]$ 到 $a[i]$ 均减小至 $a[i]$。我们需要判断 $k$ 次操作能否让这些元素全部减小至 $a[i+1]$，即比较 $k$ 与所需次数 $c = (i + 1)  (a[i] - a[i+1])$ 的大小：

- 如果 $c<k$，那么从 $a[0]$ 到 $a[i]$ 均可以减小至 $a[i+1]$；
- 如果 $c\ge k$，那么从 $a[0]$ 到 $a[i]$ 中，有 $k \bmod (i+1)$ 个元素可以额外减小 $\left\lfloor\dfrac{k}{i+1}\right\rfloor+1$，有 $k-k \bmod (i+1)$ 个元素可以额外减小 $\left\lfloor\dfrac{k}{i+1}\right\rfloor$。后续无法继续减小，应退出循环。

代码实现时，可以在 $a$ 末尾加一个 $0$，减少边界判断。

```py [sol1-Python3]
class Solution:
    def minSumSquareDiff(self, a: List[int], nums2: List[int], k1: int, k2: int) -> int:
        ans, k = 0, k1 + k2
        for i in range(len(a)):
            a[i] = abs(a[i] - nums2[i])
            ans += a[i] * a[i]
        if sum(a) <= k:
            return 0  # 所有 a[i] 均可为 0
        a.sort(reverse=True)
        a.append(0)  # 哨兵
        for i, v in enumerate(a):
            ans -= v * v
            j = i + 1
            c = j * (v - a[j])
            if c < k:
                k -= c
                continue
            v -= k // j
            return ans + k % j * (v - 1) * (v - 1) + (j - k % j) * v * v
```

```java [sol1-Java]
class Solution {
    public long minSumSquareDiff(int[] a, int[] nums2, int k1, int k2) {
        int n = a.length, k = k1 + k2;
        long ans = 0L, sum = 0L;
        for (var i = 0; i < n; ++i) {
            a[i] = Math.abs(a[i] - nums2[i]);
            sum += a[i];
            ans += (long) a[i] * a[i];
        }
        if (sum <= k) return 0;
        Arrays.sort(a);
        for (var i = n - 1; ; --i) {
            var m = n - i;
            long v = a[i], c = m * (v - (i > 0 ? a[i - 1] : 0));
            ans -= v * v;
            if (c < k) {
                k -= c;
                continue;
            }
            v -= k / m;
            return ans + k % m * (v - 1) * (v - 1) + (m - k % m) * v * v;
        }
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long minSumSquareDiff(vector<int> &a, vector<int> &nums2, int k1, int k2) {
        int n = a.size(), k = k1 + k2;
        long ans = 0L, sum = 0L;
        for (int i = 0; i < n; ++i) {
            a[i] = abs(a[i] - nums2[i]);
            sum += a[i];
            ans += (long) a[i] * a[i];
        }
        if (sum <= k) return 0;
        sort(a.begin(), a.end(), greater<int>());
        a.push_back(0); // 哨兵
        for (int i = 0;; ++i) {
            long j = i + 1, v = a[i], c = j * (v - a[j]);
            ans -= v * v;
            if (c < k) {
                k -= c;
                continue;
            }
            v -= k / j;
            return ans + k % j * (v - 1) * (v - 1) + (j - k % j) * v * v;
        }
    }
};
```

```go [sol1-Go]
func minSumSquareDiff(a, nums2 []int, k1, k2 int) int64 {
	ans, sum := 0, 0
	for i, v := range a {
		a[i] = abs(v - nums2[i])
		sum += a[i]
		ans += a[i] * a[i]
	}
	k := k1 + k2
	if sum <= k {
		return 0 // 所有 a[i] 均可为 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	a = append(a, 0) // 哨兵
	for i, v := range a {
		i++
		ans -= v * v
		if c := i * (v - a[i]); c < k {
			k -= c
			continue
		}
		v -= k / i
		ans += k%i*(v-1)*(v-1) + (i-k%i)*v*v
		break
	}
	return int64(ans)
}

func abs(x int) int { if x < 0 { return -x }; return x }
```
