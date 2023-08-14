请看 [视频讲解](https://www.bilibili.com/video/BV1bV4y1e72v/) 第四题。

### 提示 1

每个下标 $i$ 至多操作一次。

因为操作多次的话，可以只保留最后一次，前面的操作是完全多余的（反而会让其它数字变得更大）。

所以至多操作 $n$ 次。

试试枚举答案。

### 提示 2

考虑第 $t$ 秒元素之和最小是多少。

如果从一开始到第 $t$ 秒都不做任何操作的话，元素之和等于

$$
s_1 + s_2\cdot t
$$

其中 $s_1$ 是 $\textit{nums}_1$ 的元素和，$s_2$ 是 $\textit{nums}_2$ 的元素和。

考虑如何**分配**这 $t$ 次操作，让数组元素分别**减少**多少。用 $s_1 + s_2\cdot t$ 减去这些元素减少量之和的最大值，就是第 $t$ 秒元素之和的最小值。

### 提示 3

假设已经选好了要操作的元素，那么 $\textit{num}_2[i]$ 越大，操作的时间就应该越靠后。

例如 $t=3$，假设选的三个数的下标分别是 $0,1,2$，且 $\textit{nums}_2[0]\le\textit{nums}_2[1]\le\textit{nums}_2[2]$。我们可以让这些数分别减少

$$
\textit{nums}_1[0] + \textit{nums}_2[0] \cdot x \\
\textit{nums}_1[1] + \textit{nums}_2[1] \cdot y \\
\textit{nums}_1[2] + \textit{nums}_2[2] \cdot z
$$

根据 [排序不等式](https://baike.baidu.com/item/%E6%8E%92%E5%BA%8F%E4%B8%8D%E7%AD%89%E5%BC%8F/7775728)，上式中的 $x,y,z$ 应分别取 $1,2,3$，分别对应在第 $1,2,3$ 秒操作，能让第 $3$ 秒的 $s_1 + s_2\cdot t$ 减少多少。比如 $i=1$ 操作前是 $\textit{nums}_1[1] + \textit{nums}_2[1] \cdot 3$，操作后是 $\textit{nums}_2[1]$（因为在第 $2$ 秒操作的，现在是第 $3$ 秒），所以减少了 $\textit{nums}_1[1] + \textit{nums}_2[1] \cdot 2$。

### 提示 4

在第 $t$ 秒，$s_1 + s_2\cdot t$ 的**减少量的最大值**相当于求解如下问题：

按照 $\textit{nums}_2[i]$ 从小到大排序后，从 $\textit{nums}_1$ 中选择一个长为 $t$ 的子序列，子序列第 $j$ 个数的 $\textit{nums}_2[i]$ 的系数为 $j$，计算减少量的最大值。

设子序列第 $j$ 个数（$j$ 从 $1$ 开始）的下标为 $i$，那么它对减少量的贡献是

$$
\textit{nums}_1[i] + \textit{nums}_2[i] \cdot j
$$

由于上式中 $\textit{nums}_1$ 并不是有序的，无法贪心，考虑用动态规划求解。

定义 $f[i+1][j]$ 表示从前 $i$ 个数中选出 $j$ 个数，减少量最大是多少。

考虑第 $i$ 个数「选或不选」：

- 不选：$f[i+1][j] = f[i][j]$。
- 选：$f[i+1][j] = f[i][j-1] + \textit{nums}_1[i] + \textit{nums}_2[i] \cdot j$。
- 这两种情况取最大值，即

$$
f[i+1][j] = max(f[i][j], f[i][j-1] + \textit{nums}_1[i] + \textit{nums}_2[i] \cdot j)
$$

初始值 $f[0][j]=0$。

答案就是第一个满足

$$
s_1 + s_2\cdot t - f[n][t] \le x
$$

的 $t$。

如果不存在，则返回 $-1$。

代码实现时，第一个维度可以优化掉，然后像 [0-1 背包](https://www.bilibili.com/video/BV16Y411v7Y6/) 那样倒序循环 $j$。

```py [sol-Python3]
class Solution:
    def minimumTime(self, nums1: List[int], nums2: List[int], x: int) -> int:
        n = len(nums1)
        f = [0] * (n + 1)
        for a, b in sorted(zip(nums1, nums2), key=lambda z: z[1]):
            for j in range(n, 0, -1):
                f[j] = max(f[j], f[j - 1] + a + b * j)

        s1 = sum(nums1)
        s2 = sum(nums2)
        for t, v in enumerate(f):
            if s1 + s2 * t - v <= x:
                return t
        return -1
```

```java [sol-Java]
class Solution {
    public int minimumTime(List<Integer> nums1, List<Integer> nums2, int x) {
        int n = nums1.size(), s1 = 0, s2 = 0;
        var ps = new int[n][2];
        for (int i = 0; i < n; i++) {
            int a = nums1.get(i), b = nums2.get(i);
            s1 += a;
            s2 += b;
            ps[i] = new int[]{a, b};
        }
        Arrays.sort(ps, (a, b) -> a[1] - b[1]);

        var f = new int[n + 1];
        for (var p : ps)
            for (int j = n; j > 0; j--)
                f[j] = Math.max(f[j], f[j - 1] + p[0] + p[1] * j);

        for (int t = 0; t <= n; t++)
            if (s1 + s2 * t - f[t] <= x)
                return t;
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumTime(vector<int> &nums1, vector<int> &nums2, int x) {
        int n = nums1.size();
        // 对下标数组排序，避免破坏 nums1 和 nums2 的对应关系
        vector<int> ids(n);
        iota(ids.begin(), ids.end(), 0);
        sort(ids.begin(), ids.end(), [&](const int i, const int j) {
            return nums2[i] < nums2[j];
        });

        vector<int> f(n + 1);
        for (int i: ids)
            for (int j = n; j; j--)
                f[j] = max(f[j], f[j - 1] + nums1[i] + nums2[i] * j);

        int s1 = accumulate(nums1.begin(), nums1.end(), 0);
        int s2 = accumulate(nums2.begin(), nums2.end(), 0);
        for (int t = 0; t <= n; t++)
            if (s1 + s2 * t - f[t] <= x)
                return t;
        return -1;
    }
};
```

```go [sol-Go]
func minimumTime(nums1, nums2 []int, x int) int {
	s1, s2, n := 0, 0, len(nums1)
	id := make([]int, n)
	for i := range id {
		id[i] = i
		s1 += nums1[i]
		s2 += nums2[i]
	}
	sort.Slice(id, func(i, j int) bool { return nums2[id[i]] < nums2[id[j]] })

	f := make([]int, n+1)
	for _, i := range id {
		for j := n; j > 0; j-- {
			f[j] = max(f[j], f[j-1]+nums1[i]+nums2[i]*j)
		}
	}

	for t, v := range f {
		if s1+s2*t-v <= x {
			return t
		}
	}
	return -1
}

func max(a, b int) int { if b > a { return b }; return a }
```

```js [sol-JavaScript]
var minimumTime = function(nums1, nums2, x) {
    const n = nums1.length;
    let f = Array(n + 1).fill(0);
    for (const [a, b] of _.zip(nums1, nums2).sort((a, b) => a[1] - b[1]))
        for (let j = n; j; j--)
            f[j] = Math.max(f[j], f[j - 1] + a + b * j);

    const s1 = _.sum(nums1);
    const s2 = _.sum(nums2);
    for (let t = 0; t <= n; t++)
        if (s1 + s2 * t - f[t] <= x)
            return t;
    return -1;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}_1$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
