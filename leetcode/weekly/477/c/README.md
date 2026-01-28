对于数字和，用**前缀和**计算即可。关于前缀和数组的定义，请看 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

本题的难点是计算子串中的非零字符组成的数字。

先解决没有 $0$ 的简单情况。

比如 $s=12345$，计算子串 $34$ 对应的数字。

我们可以先计算 $s$ 的每个前缀（包括空前缀）对应的数字，即 $\textit{preNum} = [0,1,12,123,1234,12345]$。

想一想，如何从这些数中得到 $34$？

我们可以计算 $1234 - 12\cdot 10^2 = 1234-1200 = 34$。其中 $2$ 是子串 $34$ 的长度。

一般地，子串 $[l,r]$ 对应的数字为

$$
\textit{preNum}[r+1] - \textit{preNum}[l] \cdot 10^{r-l+1}
$$

然后解决包含 $0$ 的情况。

这意味着我们还需要计算子串中的非零字符个数，代替上式中的 $r-l+1$。

把 $s$ 中的非零字符视作 $1$，计算其前缀和。

代码实现时，注意取模。为什么可以在中途取模？原理见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

[本题视频讲解](https://www.bilibili.com/video/BV1arUKBbEks/?t=52m5s)，欢迎点赞关注~

```py [sol-Python3]
MOD = 1_000_000_007
MAX_N = 100_001

# 预处理 10 的幂
pow10 = [1] * MAX_N
for i in range(1, MAX_N):
    pow10[i] = (pow10[i - 1] * 10) % MOD

class Solution:
    def sumAndMultiply(self, s: str, queries: List[List[int]]) -> List[int]:
        n = len(s)
        sum_d = [0] * (n + 1)         # s 的前缀和
        pre_num = [0] * (n + 1)       # s 的前缀对应的数字（模 MOD）
        sum_non_zero = [0] * (n + 1)  # s 的前缀中的非零数字个数
        for i, d in enumerate(map(int, s)):
            sum_d[i + 1] = sum_d[i] + d
            pre_num[i + 1] = (pre_num[i] * 10 + d) % MOD if d else pre_num[i]
            sum_non_zero[i + 1] = sum_non_zero[i] + (d > 0)

        ans = []
        for l, r in queries:
            r += 1  # 避免下面多次计算 r+1
            length = sum_non_zero[r] - sum_non_zero[l]
            x = pre_num[r] - pre_num[l] * pow10[length]
            ans.append(x * (sum_d[r] - sum_d[l]) % MOD)
        return ans
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MAX_N = 100_001;
    private static final int[] pow10 = new int[MAX_N];
    private static boolean initialized = false;

    // 这样写比 static block 快
    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        // 预处理 10 的幂
        pow10[0] = 1;
        for (int i = 1; i < MAX_N; i++) {
            pow10[i] = (int) (pow10[i - 1] * 10L % MOD);
        }
    }

    public int[] sumAndMultiply(String s, int[][] queries) {
        int n = s.length();
        int[] sumD = new int[n + 1];       // s 的前缀和
        int[] preNum = new int[n + 1];     // s 的前缀对应的数字（模 MOD）
        int[] sumNonZero = new int[n + 1]; // s 的前缀中的非零数字个数
        for (int i = 0; i < n; i++) {
            int d = s.charAt(i) - '0';
            sumD[i + 1] = sumD[i] + d;
            preNum[i + 1] = d > 0 ? (int) ((preNum[i] * 10L + d) % MOD) : preNum[i];
            sumNonZero[i + 1] = sumNonZero[i] + (d > 0 ? 1 : 0);
        }

        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int l = queries[i][0];
            int r = queries[i][1] + 1; // 注意这里已经把 r 加一了
            int length = sumNonZero[r] - sumNonZero[l];
            long x = preNum[r] - (long) preNum[l] * pow10[length] % MOD + MOD; // +MOD 保证结果非负
            ans[i] = (int) (x * (sumD[r] - sumD[l]) % MOD);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
static constexpr int MOD = 1'000'000'007;
static constexpr int MAX_N = 100'001;
int pow10[MAX_N];

int init = [] {
    // 预处理 10 的幂
    pow10[0] = 1;
    for (int i = 1; i < MAX_N; i++) {
        pow10[i] = pow10[i - 1] * 10LL % MOD;
    }
    return 0;
}();

class Solution {
public:
    vector<int> sumAndMultiply(const string& s, const vector<vector<int>>& queries) {
        int n = s.size();
        vector<int> sum_d(n + 1), pre_num(n + 1), sum_non_zero(n + 1);
        for (int i = 0; i < n; i++) {
            int d = s[i] - '0';
            sum_d[i + 1] = sum_d[i] + d; // s 的前缀和
            pre_num[i + 1] = d > 0 ? (pre_num[i] * 10LL + d) % MOD : pre_num[i]; // s 的前缀对应的数字（模 MOD）
            sum_non_zero[i + 1] = sum_non_zero[i] + (d > 0); // s 的前缀中的非零数字个数
        }

        vector<int> ans;
        ans.reserve(queries.size()); // 预分配空间
        for (auto& q : queries) {
            int l = q[0], r = q[1] + 1; // 注意这里已经把 r 加一了
            int length = sum_non_zero[r] - sum_non_zero[l];
            long long x = pre_num[r] - 1LL * pre_num[l] * pow10[length] % MOD + MOD; // +MOD 保证结果非负
            ans.push_back(x * (sum_d[r] - sum_d[l]) % MOD);
        }
        return ans;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const maxN = 100_001

var pow10 = [maxN]int{1}

func init() {
	// 预处理 10 的幂
	for i := 1; i < maxN; i++ {
		pow10[i] = pow10[i-1] * 10 % mod
	}
}

func sumAndMultiply(s string, queries [][]int) []int {
	n := len(s)
	sumD := make([]int, n+1)       // s 的前缀和
	preNum := make([]int, n+1)     // s 的前缀对应的数字（模 mod）
	sumNonZero := make([]int, n+1) // s 的前缀中的非零数字个数
	for i, ch := range s {
		d := int(ch - '0')
		sumD[i+1] = sumD[i] + d
		preNum[i+1] = preNum[i]
		sumNonZero[i+1] = sumNonZero[i]
		if d > 0 {
			preNum[i+1] = (preNum[i]*10 + d) % mod
			sumNonZero[i+1]++
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]+1
		length := sumNonZero[r] - sumNonZero[l]
		x := preNum[r] - preNum[l]*pow10[length]%mod + mod // +mod 保证结果非负
		ans[i] = x * (sumD[r] - sumD[l]) % mod
	}
	return ans
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(n+q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 相似题目

[2156. 查找给定哈希值的子串](https://leetcode.cn/problems/find-substring-with-given-hash-value/)

## 专题训练

见下面数据结构题单的「**一、前缀和**」。

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
