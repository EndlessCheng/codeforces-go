本题 [视频讲解](https://www.bilibili.com/video/BV1aT41157bh) 已出炉，欢迎点赞三连~

---

定义 $f[i][j]$ 表示前 $i$ 个孩子分配的饼干集合为 $j$ 时，前 $i$ 个孩子的不公平程度的最小值。

考虑给第 $i$ 个孩子分配的饼干集合为 $s$，设集合 $s$ 的元素和为 $\textit{sum}[s]$，那么此时前 $i$ 个孩子的不公平程度为

$$
\max(f[i-1][j \setminus s], \textit{sum}[s])
$$

其中 $j \setminus s$ 表示从集合 $j$ 中去掉集合 $s$ 的元素后，剩余元素组成的集合。

枚举 $j$ 的所有子集 $s$，取上式的最小值即为 $f[i][j]$。

代码实现时，我们可以用一个二进制数来表示集合，其第 $i$ 位为 $1$ 表示分配了第 $i$ 块饼干，为 $0$ 表示未分配第 $i$ 块饼干。

此外通过倒序枚举 $j$，$f$ 的第一个维度可以省略。$\textit{sum}$ 也可以通过预处理得到。

时间复杂度：$O(k\cdot 3^n)$，其中 $n$ 为 $\textit{cookies}$ 的长度。由于元素个数为 $k$ 的集合有 $C(n,k)$ 个，其子集有 $2^k$ 个，根据二项式定理，$\sum C(n,k)2^k = (2+1)^n = 3^n$，所以枚举所有 $j$ 的所有子集 $s$ 的时间复杂度为 $O(3^n)$。

注：本题和 [1723. 完成所有工作的最短时间](https://leetcode.cn/problems/find-minimum-time-to-finish-all-jobs/) 是相同的。

```java [sol1-Java]
class Solution {
    public int distributeCookies(int[] cookies, int k) {
        var n = cookies.length;
        var sum = new int[1 << n];
        for (var i = 0; i < n; i++)
            for (int mask = 0, bit = 1 << i; mask < bit; ++mask)
                sum[bit | mask] = sum[mask] + cookies[i];

        var f = sum.clone();
        for (var i = 1; i < k; i++) {
            for (var j = (1 << n) - 1; j > 0; j--) {
                for (var s = j; s > 0; s = (s - 1) & j) {
                    f[j] = Math.min(f[j], Math.max(f[j ^ s], sum[s]));
                }
            }
        }
        return f[(1 << n) - 1];
    }
}
```

```C++ [sol1-C++]
class Solution {
public:
    int distributeCookies(vector<int> &cookies, int k) {
        int n = cookies.size();
        vector<int> sum(1 << n);
        for (int i = 0; i < n; i++)
            for (int mask = 0, bit = 1 << i; mask < bit; ++mask)
                sum[bit | mask] = sum[mask] + cookies[i];

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

```go [sol1-Go]
func distributeCookies(a []int, k int) int {
	n := 1 << len(a)
	sum := make([]int, n)
	for i, v := range a {
		for mask, bit := 0, 1<<i; mask < bit; mask++ {
			sum[bit|mask] = sum[mask] + v
		}
	}

	f := append([]int{}, sum...)
	for i := 1; i < k; i++ {
		for j := n - 1; j > 0; j-- {
			for s := j; s > 0; s = (s - 1) & j {
				f[j] = min(f[j], max(f[j^s], sum[s]))
			}
		}
	}
	return f[n-1]
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```
