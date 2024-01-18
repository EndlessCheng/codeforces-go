[视频讲解](https://www.bilibili.com/video/BV1bV4y1e72v/) 第四题。

## 提示 1

每个下标 $i$ 至多操作一次。

因为对下标 $i$ 操作多次的话，可以只保留最后一次，前面对下标 $i$ 的操作是完全多余的（反而会让其它数字变得更大）。

所以至多操作 $n$ 次。

试试枚举答案：

- 第 $0$ 秒元素之和最小是多少？
- 第 $1$ 秒元素之和最小是多少？
- 第 $2$ 秒元素之和最小是多少？
- ……
- 第 $n$ 秒元素之和最小是多少？

## 提示 2

考虑第 $t$ 秒（$t\le n$）元素之和最小是多少。

设 $s_1$ 是 $\textit{nums}_1$ 的元素和，$s_2$ 是 $\textit{nums}_2$ 的元素和。如果从一开始到第 $t$ 秒都不做任何操作，由于每一秒数组元素和都会增加 $s_2$，所以第 $t$ 秒的元素和等于

$$
s_1 + s_2\cdot t
$$

如何**分配**这 $t$ 次操作，可以让数组元素和减少的尽量多？

用 $s_1 + s_2\cdot t$ 减去**减少量的最大值**，就是第 $t$ 秒元素之和的最小值。

现在问题变成：

- 第 $0$ 秒减少量的最大值是多少？（显然这是 $0$）
- 第 $1$ 秒减少量的最大值是多少？
- 第 $2$ 秒减少量的最大值是多少？
- ……
- 第 $n$ 秒减少量的最大值是多少？

## 提示 3

例如 $t=3$，如果不做任何操作，数组元素和是 $s_1 + s_2\cdot 3$。

如果下标为 $i$ 的数不操作的话，它现在的值是 $\textit{nums}_1[i] + \textit{nums}_2[i] \cdot 3$，

考虑在第 $2$ 秒把它变成 $0$，那么在第 $3$ 秒它增加到 $\textit{nums}_2[i]$，所以前后一对比，这个操作让数组元素和减少了 

$$
\begin{aligned}
&\textit{nums}_1[i] + \textit{nums}_2[i] \cdot 3 - \textit{nums}_2[i]\\
=\ &\textit{nums}_1[i] + \textit{nums}_2[i] \cdot 2
\end{aligned}
$$

一般地，下标为 $i$ 的数如果不操作，它在第 $t$ 秒的值等于 $\textit{nums}_1[i] + \textit{nums}_2[i] \cdot t$。如果在第 $k$ 秒把它变成 $0$，那么在第 $t$ 秒它增加到 $\textit{nums}_2[i] \cdot (t-k)$，所以前后一对比，这个操作让数组元素和减少了

$$
\begin{aligned}
&\textit{nums}_1[i] + \textit{nums}_2[i] \cdot t - \textit{nums}_2[i] \cdot (t-k)\\
=\ &\textit{nums}_1[i] + \textit{nums}_2[i] \cdot k
\end{aligned}
$$

注意这里的**系数** $k$ 刚好等于操作的时间。由于每秒恰好操作一次，所以被操作的数字对应的系数刚好就是 $1,2,\cdots, t$。

假设我们操作的 $3$ 个数的下标分别是 $4,5,7$，且 $\textit{nums}_2[4]\le\textit{nums}_2[5]\le\textit{nums}_2[7]$。

通过分配这 $3$ 次操作，我们可以让这些数分别减少

$$
\textit{nums}_1[4] + \textit{nums}_2[4] \cdot k_1 \\
\textit{nums}_1[5] + \textit{nums}_2[5] \cdot k_2 \\
\textit{nums}_1[7] + \textit{nums}_2[7] \cdot k_3
$$

根据 [排序不等式](https://baike.baidu.com/item/%E6%8E%92%E5%BA%8F%E4%B8%8D%E7%AD%89%E5%BC%8F/7775728)，上式中的 $k_1,k_2,k_3$ 应分别取 $1,2,3$，分别对应在第 $1,2,3$ 秒操作。这样可以使元素和的减少量尽量大。

现在的问题是，选哪 $3$ 个数操作，可以使元素和的减少量尽量大？

按照 $\textit{nums}_2[i]$ 贪心？不行，减少量不仅与 $\textit{nums}_2[i]$ 有关，还与 $\textit{nums}_1[i]$ 有关。

## 提示 4

把 $\textit{nums}_1$ 和 $\textit{nums}_2$ 绑在一起，按照 $\textit{nums}_2[i]$ 从小到大排序。

> 注：下面讨论的「下标」指排序后的数组元素下标。

在第 $t$ 秒，$s_1 + s_2\cdot t$ 减少量的最大值相当于求解如下问题：

从 $0,1,2,\cdots,n-1$ 中选一个长为 $t$ 的下标子序列，依次操作子序列中的下标。选择哪 $t$ 个下标，可以使减少量尽量大？

设子序列中的第 $j$ 个下标为 $i$，那么它对减少量的贡献是

$$
\textit{nums}_1[i] + \textit{nums}_2[i] \cdot j
$$

> 注：由于已经对数组排序，根据排序不等式，子序列第 $j$ 个数对应的系数恰好是 $j$。

类似 [0-1 背包](https://www.bilibili.com/video/BV16Y411v7Y6/)，定义 $f[i+1][j]$ 表示从 $0,1,2,\cdots,i$ 中选 $j$ 个下标（$j\le i+1$），减少量最大是多少。

考虑下标 $i$「选或不选」：

- 不选，问题变成从 $0,1,2,\cdots,i-1$ 中选 $j$ 个下标，减少量最大是多少，即 $f[i+1][j] = f[i][j]$。
- 选，问题变成从 $0,1,2,\cdots,i-1$ 中选 $j-1$ 个下标，减少量最大是多少，即 $f[i+1][j] = f[i][j-1] + \textit{nums}_1[i] + \textit{nums}_2[i] \cdot j$。
- 这两种情况取最大值，即

$$
f[i+1][j] = \max(f[i][j], f[i][j-1] + \textit{nums}_1[i] + \textit{nums}_2[i] \cdot j)
$$

初始值 $f[0][0]=0$。没有数，选 $0$ 个，减少量是 $0$。

答案就是第一个满足

$$
s_1 + s_2\cdot t - f[n][t] \le x
$$

的 $t$。

如果不存在，则返回 $-1$。

代码实现时，也可以像 0-1 背包那样，去掉 $f$ 的第一个维度，并倒序循环 $j$。

```py [sol-Python3]
class Solution:
    def minimumTime(self, nums1: List[int], nums2: List[int], x: int) -> int:
        pairs = sorted(zip(nums1, nums2), key=lambda p: p[1])
        n = len(pairs)
        f = [0] * (n + 1)
        for i, (a, b) in enumerate(pairs):
            for j in range(i + 1, 0, -1):
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
        int[][] pairs = new int[n][2];
        for (int i = 0; i < n; i++) {
            int a = nums1.get(i);
            int b = nums2.get(i);
            pairs[i][0] = a;
            pairs[i][1] = b;
            s1 += a;
            s2 += b;
        }
        Arrays.sort(pairs, (a, b) -> a[1] - b[1]);

        int[] f = new int[n + 1];
        for (int i = 0; i < n; i++) {
            int a = pairs[i][0];
            int b = pairs[i][1];
            for (int j = i + 1; j > 0; j--) {
                f[j] = Math.max(f[j], f[j - 1] + a + b * j);
            }
        }

        for (int t = 0; t <= n; t++) {
            if (s1 + s2 * t - f[t] <= x) {
                return t;
            }
        }
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
        for (int i = 0; i < n; i++) {
            int a = nums1[ids[i]], b = nums2[ids[i]];
            for (int j = i + 1; j; j--) {
                f[j] = max(f[j], f[j - 1] + a + b * j);
            }
        }

        int s1 = accumulate(nums1.begin(), nums1.end(), 0);
        int s2 = accumulate(nums2.begin(), nums2.end(), 0);
        for (int t = 0; t <= n; t++) {
            if (s1 + s2 * t - f[t] <= x) {
                return t;
            }
        }
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
	// 对下标数组排序，避免破坏 nums1 和 nums2 的对应关系
	slices.SortFunc(id, func(i, j int) int { return nums2[i] - nums2[j] })

	f := make([]int, n+1)
	for i, p := range id {
		a, b := nums1[p], nums2[p]
		for j := i + 1; j > 0; j-- {
			f[j] = max(f[j], f[j-1]+a+b*j)
		}
	}

	for t, v := range f {
		if s1+s2*t-v <= x {
			return t
		}
	}
	return -1
}
```

```js [sol-JavaScript]
var minimumTime = function(nums1, nums2, x) {
    const pairs = _.zip(nums1, nums2).sort((a, b) => a[1] - b[1]);
    const n = pairs.length;
    const f = Array(n + 1).fill(0);
    for (let i = 0; i < n; i++) {
        const [a, b] = pairs[i];
        for (let j = i + 1; j; j--) {
            f[j] = Math.max(f[j], f[j - 1] + a + b * j);
        }
    }

    const s1 = _.sum(nums1);
    const s2 = _.sum(nums2);
    for (let t = 0; t <= n; t++) {
        if (s1 + s2 * t - f[t] <= x) {
            return t;
        }
    }
    return -1;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_time(nums1: Vec<i32>, nums2: Vec<i32>, x: i32) -> i32 {
        let mut pairs = nums1.iter().zip(nums2.iter()).collect::<Vec<_>>();
        pairs.sort_unstable_by(|&a, &b| a.1.cmp(&b.1));

        let n = pairs.len();
        let mut f = vec![0; n + 1];
        for (i, &(a, b)) in pairs.iter().enumerate() {
            for j in (1..=i + 1).rev() {
                f[j] = f[j].max(f[j - 1] + a + b * j as i32);
            }
        }

        let s1 = nums1.iter().sum::<i32>();
        let s2 = nums2.iter().sum::<i32>();
        for (t, &v) in f.iter().enumerate() {
            if s1 + s2 * t as i32 - v <= x {
                return t as i32;
            }
        }
        -1
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}_1$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
