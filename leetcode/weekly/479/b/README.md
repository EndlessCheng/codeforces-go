## 写法一

预处理 $5\times 10^5$ 以内的质数列表 $\textit{primes}$。

计算质数列表的前缀和，如果前缀和是质数，那么加到特殊质数列表 $\textit{specialPrimes}$ 中。

由于 $\textit{specialPrimes}$ 是有序的，我们可以在 $\textit{specialPrimes}$ 中**二分查找**最大的 $\le n$ 的质数。关于二分查找的原理，请看[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

可以在 $\textit{specialPrimes}$ 的最前面加个 $0$，从而简化判断逻辑。

[本题视频讲解](https://www.bilibili.com/video/BV1sv2fB4Evi/?t=21m57s)，欢迎点赞关注~

```py [sol-Python3]
MX = 500_000

is_prime = [False] * 2 + [True] * (MX - 1)  # 0 和 1 不是质数
primes = []
for i in range(2, MX + 1):
    if is_prime[i]:
        primes.append(i)
        for j in range(i * i, MX + 1, i):
            is_prime[j] = False  # j 是质数 i 的倍数

special_primes = [0]  # 哨兵
sum_p = 0
for p in primes:
    sum_p += p
    if sum_p > MX:
        break
    if is_prime[sum_p]:
        special_primes.append(sum_p)


class Solution:
    def largestPrime(self, n: int) -> int:
        # 二分找 <= n 的最大特殊质数
        i = bisect_right(special_primes, n) - 1
        return special_primes[i]
```

```java [sol-Java]
class Solution {
    private static final int MX = 1_000_000;
    private static final boolean[] isPrime = new boolean[MX + 1];
    private static final List<Integer> primes = new ArrayList<>();
    private static final List<Integer> specialPrimes = new ArrayList<>();
    private static boolean initialized = false;

    // 这样写比 static block 更快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        Arrays.fill(isPrime, true);
        isPrime[0] = isPrime[1] = false; // 0 和 1 不是质数
        for (int i = 2; i <= MX; i++) {
            if (isPrime[i]) {
                primes.add(i);
                for (long j = (long) i * i; j <= MX; j += i) {
                    isPrime[(int) j] = false; // j 是质数 i 的倍数
                }
            }
        }

        specialPrimes.add(0); // 哨兵
        int sum = 0;
        for (int p : primes) {
            sum += p;
            if (sum > MX) {
                break;
            }
            if (isPrime[sum]) {
                specialPrimes.add(sum);
            }
        }
    }

    public int largestPrime(int n) {
        init();

        // 二分找 <= n 的最大特殊质数
        int i = Collections.binarySearch(specialPrimes, n + 1);
        if (i < 0) i = ~i;
        return specialPrimes.get(i - 1);
    }
}
```

```cpp [sol-C++]
const int MX = 500'000;
bool is_prime[MX + 1];
vector<int> primes;
vector<int> special_primes = {0}; // 哨兵

auto init = [] {
    ranges::fill(is_prime, true);
    is_prime[0] = is_prime[1] = false; // 0 和 1 不是质数
    for (int i = 2; i <= MX; i++) {
        if (is_prime[i]) {
            primes.push_back(i);
            for (long long j = 1LL * i * i; j <= MX; j += i) {
                is_prime[j] = false; // j 是质数 i 的倍数
            }
        }
    }

    int sum = 0;
    for (int p : primes) {
        sum += p;
        if (sum > MX) {
            break;
        }
        if (is_prime[sum]) {
            special_primes.push_back(sum);
        }
    }
    return 0;
}();

class Solution {
public:
    int largestPrime(int n) {
        // 二分找 <= n 的最大特殊质数
        auto it = ranges::upper_bound(special_primes, n);
        return *--it;
    }
};
```

```go [sol-Go]
const mx = 500_000

var primes []int
var np [mx + 1]bool
var specialPrimes = []int{0} // 哨兵

func init() {
	for i := 2; i <= mx; i++ {
		if !np[i] {
			primes = append(primes, i)
			for j := i * i; j <= mx; j += i {
				np[j] = true
			}
		}
	}

	sum := 0
	for _, p := range primes {
		sum += p
		if sum > mx {
			break
		}
		if !np[sum] {
			specialPrimes = append(specialPrimes, sum)
		}
	}
}

func largestPrime(n int) int {
	// 二分找 <= n 的最大特殊质数
	i := sort.SearchInts(specialPrimes, n+1) - 1
	return specialPrimes[i]
}
```

#### 复杂度分析

> 预处理的时空复杂度为 $\mathcal{O}(U\log\log U)$，$U=5\times 10^5$，不计入。

- 时间复杂度：$\mathcal{O}(\log K)$，其中 $K=40$。在本题数据范围下，有 $40$ 个特殊质数。
- 空间复杂度：$\mathcal{O}(1)$。

## 写法二

也可以用双指针直接预处理所有答案。

```py [sol-Python3]
MX = 500_000

is_prime = [False] * 2 + [True] * (MX - 2)  # 0 和 1 不是质数
primes = []
for i in range(2, MX):
    if is_prime[i]:
        primes.append(i)
        for j in range(i * i, MX, i):
            is_prime[j] = False  # j 是质数 i 的倍数

ans = [0] * (MX + 1)
sum_p = last = j = 0
for i in range(2, MX + 1):
    if sum_p + primes[j] <= i:
        sum_p += primes[j]
        j += 1
        if is_prime[sum_p]:
            last = sum_p
    ans[i] = last


class Solution:
    def largestPrime(self, n: int) -> int:
        return ans[n]
```

```java [sol-Java]
class Solution {
    private static final int MX = 1_000_000;
    private static final boolean[] isPrime = new boolean[MX + 1];
    private static final List<Integer> primes = new ArrayList<>();
    private static final int[] ans = new int[MX + 1];
    private static boolean initialized = false;

    // 这样写比 static block 更快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        Arrays.fill(isPrime, true);
        isPrime[0] = isPrime[1] = false; // 0 和 1 不是质数
        for (int i = 2; i <= MX; i++) {
            if (isPrime[i]) {
                primes.add(i);
                for (long j = (long) i * i; j <= MX; j += i) {
                    isPrime[(int) j] = false; // j 是质数 i 的倍数
                }
            }
        }

        int sum = 0;
        int last = 0;
        int j = 0;
        for (int i = 2; i <= MX; i++) {
            if (sum + primes.get(j) <= i) {
                sum += primes.get(j);
                j++;
                if (isPrime[sum]) {
                    last = sum;
                }
            }
            ans[i] = last;
        }
    }

    public int largestPrime(int n) {
        init();
        return ans[n];
    }
}
```

```cpp [sol-C++]
const int MX = 500'000;
bool is_prime[MX + 1];
vector<int> primes;
int ans[MX + 1];

auto init = [] {
    ranges::fill(is_prime, true);
    is_prime[0] = is_prime[1] = false; // 0 和 1 不是质数
    for (int i = 2; i <= MX; i++) {
        if (is_prime[i]) {
            primes.push_back(i);
            for (long long j = 1LL * i * i; j <= MX; j += i) {
                is_prime[j] = false; // j 是质数 i 的倍数
            }
        }
    }

    int sum = 0, last = 0, j = 0;
    for (int i = 2; i <= MX; i++) {
        if (sum + primes[j] <= i) {
            sum += primes[j];
            j++;
            if (is_prime[sum]) {
                last = sum;
            }
        }
        ans[i] = last;
    }
    return 0;
}();

class Solution {
public:
    int largestPrime(int n) {
        return ans[n];
    }
};
```

```go [sol-Go]
const mx = 500_000

var primes []int
var np [mx + 1]bool
var ans [mx + 1]int

func init() {
	for i := 2; i <= mx; i++ {
		if !np[i] {
			primes = append(primes, i)
			for j := i * i; j <= mx; j += i {
				np[j] = true
			}
		}
	}

	sum, last, j := 0, 0, 0
	for i := 2; i <= mx; i++ {
		if sum+primes[j] <= i {
			sum += primes[j]
			j++
			if !np[sum] {
				last = sum
			}
		}
		ans[i] = last
	}
}

func largestPrime(n int) int {
	return ans[n]
}
```

#### 复杂度分析

> 预处理的时空复杂度为 $\mathcal{O}(U\log\log U)$，$U=5\times 10^5$，不计入。

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面数学题单的「**§1.2 预处理质数（筛质数）**」。

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
