预处理 $[2,10^6]$ 中的质数列表 $\textit{primes}$，其 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 数组为 $\textit{sum}$（不含 $0$）。

贪心地，对于一个最优方案，如果其不是 $\textit{primes}$ 的前缀，那么把其中的一个质数减小，和变小，也满足 $\le s$ 的要求。

所以**一定存在一个最优方案，选的是 $\textit{primes}$ 的前缀**。

在 $\textit{primes}$ 中 [二分查找](https://www.bilibili.com/video/BV1AP41137w7/) 最后一个 $\le n$ 的数的下标 $i$。

然后，在 $\textit{sum}$ 的前缀 $[0,i]$ 中二分查找最后一个 $\le s$ 的数的下标 $j$。那么 $\textit{primes}$ 的前缀 $[0,j]$ 即为答案。

> 注：如果有多个答案，返回其中任意一个都可以。

[本题视频讲解](https://www.bilibili.com/video/BV1wnpj6VEtP/?t=3m3s)，欢迎点赞关注~

```py [sol-Python3]
MX = 1_000_001
is_prime = [True] * MX
primes = []
p_sum = []
s = 0
for i in range(2, MX):
    if is_prime[i]:
        primes.append(i)
        s += i
        p_sum.append(s)
        for j in range(i * i, MX, i):
            is_prime[j] = False  # j 是质数 i 的倍数

class Solution:
    def maxPrimes(self, n: int, s: int) -> list[int]:
        i = bisect_right(primes, n)
        j = bisect_right(p_sum, s, 0, i)
        return primes[:j]
```

```java [sol-Java]
class Solution {
    private static final int MX = 1_000_001;
    private static final boolean[] isPrime = new boolean[MX];
    private static final List<Integer> primes = new ArrayList<>();
    private static final List<Long> sum = new ArrayList<>();
    private static boolean initialized = false;

    // 这样写比 static block 快
    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        Arrays.fill(isPrime, true);
        long s = 0;
        for (int i = 2; i < MX; i++) {
            if (isPrime[i]) {
                primes.add(i);
                s += i;
                sum.add(s);
                for (long j = (long) i * i; j < MX; j += i) {
                    isPrime[(int) j] = false; // j 是质数 i 的倍数
                }
            }
        }
    }

    public List<Integer> maxPrimes(int n, int s) {
        int i = Collections.binarySearch(primes, n + 1);
        if (i < 0) i = ~i; // 见 Collections.binarySearch 源码

        int j = Collections.binarySearch(sum.subList(0, i), (long) s + 1);
        if (j < 0) j = ~j;

        return primes.subList(0, j);
    }
}
```

```cpp [sol-C++]
constexpr int MX = 1'000'001;
bool is_prime[MX];
vector<int> primes;
vector<long long> sum;

auto init = [] {
    ranges::fill(is_prime, true);
    long long s = 0;
    for (int i = 2; i < MX; i++) {
        if (is_prime[i]) {
            primes.push_back(i);
            s += i;
            sum.push_back(s);
            for (long long j = 1LL * i * i; j < MX; j += i) {
                is_prime[j] = false; // j 是质数 i 的倍数
            }
        }
    }
    return 0;
}();

class Solution {
public:
    vector<int> maxPrimes(int n, int s) {
        int i = ranges::upper_bound(primes, n) - primes.begin();
        int j = ranges::upper_bound(sum.begin(), sum.begin() + i, s) - sum.begin();
        return vector(primes.begin(), primes.begin() + j);
    }
};
```

```go [sol-Go]
const mx = 1_000_001

var np = [mx]bool{}
var primes []int
var sum []int

func init() {
	s := 0
	for i := 2; i < mx; i++ {
		if !np[i] {
			primes = append(primes, i)
			s += i
			sum = append(sum, s)
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func maxPrimes(n int, s int) []int {
	i := sort.SearchInts(primes, n+1)
	j := sort.SearchInts(sum[:i], s+1)
	return primes[:j]
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}\left(\log \dfrac{U}{\log U}\right)$，其中 $U=10^6$。$[2,U]$ 中的质数个数为 $\mathcal{O}\left(\dfrac{U}{\log U}\right)$。不计入返回答案的时间。
- 空间复杂度：$\mathcal{O}(1)$。不计入答案的空间。

## 附：二分查找常用转化表

| **需求**  | **写法**  |  **如果不存在** | 
|---|---|---|
| $\ge x$ 的第一个整数的下标  | $\texttt{lowerBound}(\textit{nums},x)$  | 结果为 $n$ |
| $> x$ 的第一个整数的下标 | $\texttt{lowerBound}(\textit{nums},x+1)$  | 结果为 $n$ |
| $< x$ 的最后一个整数的下标  | $\texttt{lowerBound}(\textit{nums},x)-1$  | 结果为 $-1$ |
| $\le x$ 的最后一个整数的下标  | $\texttt{lowerBound}(\textit{nums},x+1)-1$  | 结果为 $-1$ |

## 专题训练

1. 数学题单的「**§1.2 预处理质数（筛质数）**」。
2. 贪心题单的「**§1.1 从最小/最大开始贪心**」。

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
