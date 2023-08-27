## 视频讲解

见 [视频讲解](https://www.bilibili.com/video/BV1Em4y1T7Bq/) 第四题。

## 前置知识

[【模板讲解】树上倍增算法](https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/solution/mo-ban-jiang-jie-shu-shang-bei-zeng-suan-v3rw/)

按照这个模板改改就能过。

## 思路

利用倍增算法，预处理每个节点 $x$ 的第 $2^i$ 个祖先节点，以及从 $x$ 的父节点到 $x$ 的第 $2^i$ 个祖先节点的节点编号之和。

最后枚举起点 $x$，一边向上跳一边累加节点编号。

```py [sol-Python3]
class Solution:
    def getMaxFunctionValue(self, receiver: List[int], k: int) -> int:
        n = len(receiver)
        m = k.bit_length() - 1
        pa = [[(p, p)] + [None] * m for p in receiver]
        for i in range(m):
            for x in range(n):
                p, s = pa[x][i]
                pp, ss = pa[p][i]
                pa[x][i + 1] = (pp, s + ss)  # 合并节点值之和

        ans = 0
        for i in range(n):
            x = sum = i
            for j in range(m + 1):
                if (k >> j) & 1:  # k 的二进制从低到高第 j 位是 1
                    x, s = pa[x][j]
                    sum += s
            ans = max(ans, sum)
        return ans
```

```java [sol-Java]
class Solution {
    public long getMaxFunctionValue(List<Integer> receiver, long K) {
        int n = receiver.size();
        int m = 64 - Long.numberOfLeadingZeros(K); // K 的二进制长度
        var pa = new int[m][n];
        var sum = new long[m][n];
        for (int i = 0; i < n; i++) {
            pa[0][i] = receiver.get(i);
            sum[0][i] = receiver.get(i);
        }
        for (int i = 0; i < m - 1; i++) {
            for (int x = 0; x < n; x++) {
                int p = pa[i][x];
                pa[i + 1][x] = pa[i][p];
                sum[i + 1][x] = sum[i][x] + sum[i][p]; // 合并节点值之和
            }
        }

        long ans = 0;
        for (int i = 0; i < n; i++) {
            long s = i;
            int x = i;
            for (long k = K; k > 0; k &= k - 1) {
                int ctz = Long.numberOfTrailingZeros(k);
                s += sum[ctz][x];
                x = pa[ctz][x];
            }
            ans = Math.max(ans, s);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long getMaxFunctionValue(vector<int> &receiver, long long K) {
        int n = receiver.size();
        int m = 64 - __builtin_clzll(K); // K 的二进制长度
        vector<vector<pair<int, long long>>> pa(m, vector<pair<int, long long>>(n));
        for (int i = 0; i < n; i++)
            pa[0][i] = {receiver[i], receiver[i]};
        for (int i = 0; i < m - 1; i++) {
            for (int x = 0; x < n; x++) {
                auto [p, s] = pa[i][x];
                auto [pp, ss] = pa[i][p];
                pa[i + 1][x] = {pp, s + ss}; // 合并节点值之和
            }
        }

        long long ans = 0;
        for (int i = 0; i < n; i++) {
            long long sum = i;
            int x = i;
            for (long long k = K; k; k &= k - 1) {
                auto [p, s] = pa[__builtin_ctzll(k)][x];
                sum += s;
                x = p;
            }
            ans = max(ans, sum);
        }
        return ans;
    }
};
```

```go [sol-Go]
func getMaxFunctionValue(receiver []int, K int64) int64 {
	type pair struct{ pa, sum int }
	n, m := len(receiver), bits.Len(uint(K))
	pa := make([][]pair, n)
	for i, p := range receiver {
		pa[i] = make([]pair, m)
		pa[i][0] = pair{p, p}
	}
	for i := 0; i+1 < m; i++ {
		for x := range pa {
			p := pa[x][i]
			pp := pa[p.pa][i]
			pa[x][i+1] = pair{pp.pa, p.sum + pp.sum} // 合并节点值之和
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		x := i
		sum := i // 节点值之和，初始为节点 i
		for k := uint(K); k > 0; k &= k - 1 {
			p := pa[x][bits.TrailingZeros(k)]
			sum += p.sum
			x = p.pa
		}
		ans = max(ans, sum)
	}
	return int64(ans)
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log k)$，其中 $n$ 为 $\textit{receiver}$ 的长度。
- 空间复杂度：$\mathcal{O}(n\log k)$。

## 思考题

做到 $\mathcal{O}(n)$ 时间复杂度。

提示：内向基环树。之前写过一篇 [内向基环树的介绍](https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting/solution/nei-xiang-ji-huan-shu-tuo-bu-pai-xu-fen-c1i1b/)。
