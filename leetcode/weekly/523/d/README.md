本题是**恰好型 0-1 背包**。请先完成模板题 [2915. 和为目标值的最长子序列的长度](https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/)，[我的题解](https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/solutions/2502839/mo-ban-qia-hao-zhuang-man-xing-0-1-bei-b-0nca/)。

本题在 0-1 背包的基础上，要输出具体方案（选了那些数），还要求字典序尽量小。

字典序尽量小，那么第一个数越小越好。例如 $s = 18 = 5+13 = 7+11$，对于质数 $5,7,11,13$，在考虑「选或不选」时，要**优先选靠前的质数**。

设质数之和的最大值为 $M$。对于第一个质数 $2$，我们需要知道，如果选了 $2$，能否从大于 $2$ 的质数中得到和 $M-2$。

所以本题在定义状态时，要定义成质数数组的**后缀**。

具体地，设 $[2,n]$ 中的质数数组为 $\textit{primes}$，长为 $m$。

定义 $f[i][j]$ 表示从 $\textit{primes}$ 的下标 $[i,m-1]$ 中选质数，所选质数之和**恰好**为 $j$ 时，所选质数个数的最小值。如果无解，那么 $f[i][j]=\infty$。

状态转移方程为

$$
f[i][j] =
\begin{cases} 
f[i+1][j], & j < \textit{primes}[i]     \\
\min(f[i+1][j], f[i+1][j-\textit{primes}[i]]+1), & j \ge \textit{primes}[i]      \\
\end{cases}
$$

初始值：$f[m][0] = 0$，其余 $f[m][j] = \infty$。

找到最大的 $j_0$，满足 $f[0][j_0] < \infty$。

从状态 $(i,j) = (0,j_0)$ 开始，还原状态转移的路径，也就是转移方程计算 $\min$ 的时候，是选，还是不选。

- 如果 $f[i+1][j-\textit{primes}[i]]+1 \le f[i+1][j]$，那么选 $\textit{primes}[i]$。然后更新 $(i,j)$ 为 $(i+1, j -\textit{primes}[i])$。
- 如果 $f[i+1][j-\textit{primes}[i]]+1 > f[i+1][j]$，那么不选 $\textit{primes}[i]$。然后更新 $(i,j)$ 为 $(i+1, j)$。

[本题视频讲解](https://www.bilibili.com/video/BV1wnpj6VEtP/?t=8m27s)，欢迎点赞关注~

```py [sol-Python3]
MX = 1001
is_prime = [True] * MX
primes = []
for i in range(2, MX):
    if is_prime[i]:
        primes.append(i)
        for j in range(i * i, MX, i):
            is_prime[j] = False  # j 是质数 i 的倍数

class Solution:
    def maxPrimeSubset(self, n: int, s: int) -> list[int]:
        m = bisect_right(primes, n)

        f = [[inf] * (s + 1) for _ in range(m + 1)]
        f[m][0] = 0

        # 倒着 DP，这样可以正序输出方案，且保证越靠前的质数越小
        for i in range(m - 1, -1, -1):
            x = primes[i]
            for j in range(s + 1):
                if j < x:
                    f[i][j] = f[i + 1][j]
                else:
                    f[i][j] = min(f[i + 1][j], f[i + 1][j - x] + 1)

        j = s
        while j > 0 and f[0][j] == inf:
            j -= 1

        ans = []
        for i in range(m):
            if j == 0:
                break
            x = primes[i]
            # 靠前的质数越小越好，f[i+1][j-x]+1 = f[i+1][j] 的时候要选
            if f[i + 1][j - x] + 1 <= f[i + 1][j]:
                ans.append(x)
                j -= x

        return ans
```

```java [sol-Java]
class Solution {
    private static final int MX = 1001;
    private static final boolean[] isPrime = new boolean[MX];
    private static final List<Integer> primes = new ArrayList<>();
    private static boolean initialized = false;

    // 这样写比 static block 快
    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        Arrays.fill(isPrime, true);
        for (int i = 2; i < MX; i++) {
            if (isPrime[i]) {
                primes.add(i);
                for (long j = (long) i * i; j < MX; j += i) {
                    isPrime[(int) j] = false; // j 是质数 i 的倍数
                }
            }
        }
    }

    public List<Integer> maxPrimeSubset(int n, int s) {
        int m = Collections.binarySearch(primes, n + 1);
        if (m < 0) m = ~m; // 见 Collections.binarySearch 源码

        int[][] f = new int[m + 1][s + 1];
        for (int[] row : f) {
            Arrays.fill(row, Integer.MAX_VALUE / 2);
        }
        f[m][0] = 0;

        // 倒着 DP，这样可以正序输出方案，且保证越靠前的质数越小
        for (int i = m - 1; i >= 0; i--) {
            int x = primes.get(i);
            for (int j = 0; j <= s; j++) {
                if (j < x) {
                    f[i][j] = f[i + 1][j];
                } else {
                    f[i][j] = Math.min(f[i + 1][j], f[i + 1][j - x] + 1);
                }
            }
        }

        int j = s;
        while (j > 0 && f[0][j] == Integer.MAX_VALUE / 2) {
            j--;
        }

        List<Integer> ans = new ArrayList<>();
        for (int i = 0; j > 0; i++) {
            int x = primes.get(i);
            // 靠前的质数越小越好，f[i+1][j-x]+1 = f[i+1][j] 的时候要选
            if (f[i + 1][j - x] + 1 <= f[i + 1][j]) {
                ans.add(x);
                j -= x;
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
constexpr int MX = 1001;
bool is_prime[MX];
vector<int> primes;

auto init = [] {
    ranges::fill(is_prime, true);
    for (int i = 2; i < MX; i++) {
        if (is_prime[i]) {
            primes.push_back(i);
            for (long long j = 1LL * i * i; j < MX; j += i) {
                is_prime[j] = false; // j 是质数 i 的倍数
            }
        }
    }
    return 0;
}();

class Solution {
public:
    vector<int> maxPrimeSubset(int n, int s) {
        int m = ranges::upper_bound(primes, n) - primes.begin();

        vector f(m + 1, vector<int>(s + 1, INT_MAX / 2));
        f[m][0] = 0;

        // 倒着 DP，这样可以正序输出方案，且保证越靠前的质数越小
        for (int i = m - 1; i >= 0; i--) {
            int x = primes[i];
            for (int j = 0; j <= s; j++) {
                if (j < x) {
                    f[i][j] = f[i + 1][j];
                } else {
                    f[i][j] = min(f[i + 1][j], f[i + 1][j - x] + 1);
                }
            }
        }

        int j = s;
        while (j > 0 && f[0][j] == INT_MAX / 2) {
            j--;
        }

        vector<int> ans;
        for (int i = 0; j > 0; i++) {
            int x = primes[i];
            // 靠前的质数越小越好，f[i+1][j-x]+1 = f[i+1][j] 的时候要选
            if (f[i + 1][j - x] + 1 <= f[i + 1][j]) {
                ans.push_back(x);
                j -= x;
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
const mx = 1001
var np = [mx]bool{}
var primes []int

func init() {
	for i := 2; i < mx; i++ {
		if !np[i] {
			primes = append(primes, i)
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func maxPrimeSubset(n int, s int) (ans []int) {
	m := sort.SearchInts(primes, n+1)

	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, s+1)
		for j := range f[i] {
			f[i][j] = math.MaxInt / 2 // 避免加法溢出
		}
	}
	f[m][0] = 0

	// 倒着 DP，这样可以正序输出方案，且保证越靠前的质数越小
	for i := m - 1; i >= 0; i-- {
		x := primes[i]
		for j := range s + 1 {
			if j < x {
				f[i][j] = f[i+1][j]
			} else {
				f[i][j] = min(f[i+1][j], f[i+1][j-x]+1)
			}
		}
	}

	j := s
	for j > 0 && f[0][j] == math.MaxInt/2 {
		j--
	}

	for i := 0; j > 0; i++ {
		x := primes[i]
		// 靠前的质数越小越好，f[i+1][j-x]+1 = f[i+1][j] 的时候要选
		if f[i+1][j-x]+1 <= f[i+1][j] {
			ans = append(ans, x)
			j -= x
		}
	}

	return
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(ms)$，其中 $m = \mathcal{O}(n/\log n)$ 是 $\le n$ 的质数个数。
- 空间复杂度：$\mathcal{O}(ms)$。

## 专题训练

见下面动态规划题单的「**§3.1 0-1 背包**」和「**专题：输出具体方案（打印方案）**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
