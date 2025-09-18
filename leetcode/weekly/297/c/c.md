本题 [视频讲解](https://www.bilibili.com/video/BV1aT41157bh) 已出炉，额外介绍了枚举子集的位运算原理，欢迎点赞三连~

---

**阅读提示**：请注意「前 $i$ 个」和「第 $i$ 个」的区别，前者用来表示状态，后者用于参与状态转移。

定义 $f[i][j]$ 表示前 $i$ 个孩子分配的饼干集合为 $j$ 时，前 $i$ 个孩子的不公平程度的最小值。

下文中 $j \setminus s$ 表示从集合 $j$ 中去掉集合 $s$ 的元素后，剩余元素组成的集合。

考虑给第 $i$ 个孩子分配的饼干集合为 $s$，设集合 $s$ 的元素和为 $\textit{sum}[s]$，分类讨论：

- 如果 $\textit{sum}[s] > f[i-1][j \setminus s]$，说明给第 $i$ 个孩子分配的饼干比前面的孩子多，不公平程度变为 $\textit{sum}[s]$；
- 如果 $\textit{sum}[s] \le f[i-1][j \setminus s]$，说明给第 $i$ 个孩子分配的饼干没有比前面的孩子多，不公平程度不变，仍为 $f[i-1][j \setminus s]$。

因此，给第 $i$ 个孩子分配饼干集合 $s$ 后，前 $i$ 个孩子的不公平程度为

$$
\max(f[i-1][j \setminus s], \textit{sum}[s])
$$

枚举 $j$ 的所有子集 $s$，则有

$$
f[i][j]=\min_{s\subseteq j} \max(f[i-1][j \setminus s], \textit{sum}[s])
$$

代码实现时，我们可以用一个二进制数来表示集合，其第 $i$ 位为 $1$ 表示分配了第 $i$ 块饼干，为 $0$ 表示未分配第 $i$ 块饼干。

此外通过倒序枚举 $j$，$f$ 的第一个维度可以省略。$\textit{sum}$ 也可以通过预处理得到。

#### 复杂度分析

- 时间复杂度：$O(k\cdot 3^n)$，其中 $n$ 为 $\textit{cookies}$ 的长度。由于元素个数为 $i$ 的集合有 $C(n,i)$ 个，其子集有 $2^i$ 个，根据二项式定理，$\sum C(n,i)2^i = (2+1)^n = 3^n$，所以枚举所有 $j$ 的所有子集 $s$ 的时间复杂度为 $O(3^n)$。
- 空间复杂度：$O(2^n)$。

注：本题和 [1723. 完成所有工作的最短时间](https://leetcode.cn/problems/find-minimum-time-to-finish-all-jobs/) 是相同的。

注 2：我另写了一篇 [题解](https://leetcode.cn/problems/find-minimum-time-to-finish-all-jobs/solution/by-endlesscheng-d2oa/)，讲述了如何用 Python 的状压 DP 写法通过 1723 题。

```py [sol-Python3]
class Solution:
    def distributeCookies(self, cookies: List[int], k: int) -> int:
        m = 1 << len(cookies)
        SUM = [0] * m
        for i, v in enumerate(cookies):
            bit = 1 << i
            for j in range(bit):
                SUM[bit | j] = SUM[j] + v

        f = SUM.copy()
        for _ in range(1, k):
            for j in range(m - 1, 0, -1):
                s = j
                while s:
                    v = f[j ^ s]
                    if SUM[s] > v: v = SUM[s]  # 不要用 max 和 min，那样会有额外的函数调用开销
                    if v < f[j]: f[j] = v
                    s = (s - 1) & j
        return f[-1]
```

```java [sol-Java]
class Solution {
    public int distributeCookies(int[] cookies, int k) {
        int n = cookies.length;
        int[] sum = new int[1 << n];
        for (int i = 0; i < n; i++)
            for (int j = 0, bit = 1 << i; j < bit; j++)
                sum[bit | j] = sum[j] + cookies[i];

        int[] f = sum.clone();
        for (int i = 1; i < k; i++) {
            for (int j = (1 << n) - 1; j > 0; j--) {
                for (int s = j; s > 0; s = (s - 1) & j) {
                    f[j] = Math.min(f[j], Math.max(f[j ^ s], sum[s]));
                }
            }
        }
        return f[(1 << n) - 1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int distributeCookies(vector<int>& cookies, int k) {
        int n = cookies.size();
        vector<int> sum(1 << n);
        for (int i = 0; i < n; i++)
            for (int j = 0, bit = 1 << i; j < bit; j++)
                sum[bit | j] = sum[j] + cookies[i];

        vector<int> f(sum);
        for (int i = 1; i < k; i++) {
            for (int j = (1 << n) - 1; j; j--) {
                for (int s = j; s; s = (s - 1) & j) {
                    f[j] = min(f[j], max(f[j ^ s], sum[s]));
                }
            }
        }
        return f.back();
    }
};
```

```go [sol-Go]
func distributeCookies(cookies []int, k int) int {
	u := 1 << len(cookies)
	sum := make([]int, u)
	for i, v := range cookies {
		highBit := 1 << i
		for j := range highBit {
			sum[highBit|j] = sum[j] + v
		}
	}

	f := slices.Clone(sum)
	for range k - 1 {
		for j := u - 1; j > 0; j-- {
			for s := j; s > 0; s = (s - 1) & j {
				f[j] = min(f[j], max(f[j^s], sum[s]))
			}
		}
	}
	return f[u-1]
}
```

```js [sol-JavaScript]
var distributeCookies = function (cookies, k) {
    const n = cookies.length;
    const sum = Array(1 << n).fill(0);
    for (let i = 0; i < n; i++) {
        for (let j = 0, bit = 1 << i; j < bit; j++) {
            sum[bit | j] = sum[j] + cookies[i];
        }
    }

    const f = sum.slice();
    for (let i = 1; i < k; i++) {
        for (let j = (1 << n) - 1; j; j--) {
            for (let s = j; s; s = (s - 1) & j) {
                f[j] = Math.min(f[j], Math.max(sum[j ^ s], f[s]));
            }
        }
    }
    return f[(1 << n) - 1];
};
```