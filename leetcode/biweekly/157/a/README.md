**前置知识**：[如何判断质数](https://leetcode.cn/problems/prime-in-diagonal/solutions/2216347/pan-duan-zhi-shu-by-endlesscheng-m6nt/)。

本题可以枚举所有 $\dfrac{n(n+1)}{2}$ 个子串，判断子串对应的十进制数是否为质数。

把质数加到一个列表中，排序去重后，计算前三大之和，即为答案。

也可以用哈希表或者有序集合去重。

注意 $1$ 不是质数。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1cqjgzdEPP/)，欢迎点赞关注~

```py [sol-Python3]
def is_prime(n: int) -> bool:
    for i in range(2, isqrt(n) + 1):
        if n % i == 0:
            return False
    return n >= 2

class Solution:
    def sumOfLargestPrimes(self, s: str) -> int:
        primes = set()
        for i in range(len(s)):
            x = 0
            for ch in s[i:]:
                x = x * 10 + int(ch)
                if is_prime(x):
                    primes.add(x)
        return sum(sorted(primes)[-3:])
```

```java [sol-Java]
class Solution {
    private boolean isPrime(long n) {
        for (long i = 2; i * i <= n; i++) {
            if (n % i == 0) {
                return false;
            }
        }
        return n >= 2;
    }

    public long sumOfLargestPrimes(String S) {
        char[] s = S.toCharArray();
        TreeSet<Long> set = new TreeSet<>();
        for (int i = 0; i < s.length; i++) {
            long x = 0;
            for (int j = i; j < s.length; j++) {
                x = x * 10 + (s[j] - '0');
                if (isPrime(x)) {
                    set.add(x);
                }
            }
        }

        return set.descendingSet()
                  .stream()
                  .limit(3)
                  .mapToLong(Long::longValue)
                  .sum();
    }
}
```

```cpp [sol-C++]
class Solution {
    bool is_prime(long long n) {
        for (long long i = 2; i * i <= n; i++) {
            if (n % i == 0) {
                return false;
            }
        }
        return n >= 2;
    }

public:
    long long sumOfLargestPrimes(string s) {
        vector<long long> primes;
        for (int i = 0; i < s.size(); i++) {
            long long x = 0;
            for (int j = i; j < s.size(); j++) {
                x = x * 10 + (s[j] - '0');
                if (is_prime(x)) {
                    primes.push_back(x);
                }
            }
        }

        // 排序，去重，计算前三大之和
        ranges::sort(primes, greater());
        int n = ranges::unique(primes).begin() - primes.begin();
        return reduce(primes.begin(), primes.begin() + min(n, 3), 0LL);
    }
};
```

```go [sol-Go]
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2
}

func sumOfLargestPrimes(s string) (ans int64) {
	primes := []int{}
	n := len(s)
	for i := range n {
		x := 0
		for _, b := range s[i:] {
			x = x*10 + int(b-'0')
			if isPrime(x) {
				primes = append(primes, x)
			}
		}
	}

	slices.Sort(primes)
	primes = slices.Compact(primes) // 去重

	for _, p := range primes[max(len(primes)-3, 0):] {
		ans += int64(p)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(10^{n/2})$，其中 $n$ 是 $s$ 的长度。乍一看，时间复杂度是 $\mathcal{O}(n^2\cdot 10^{n/2})$，但实际上有 $n$ 个长为 $1$ 的子串，$n-1$ 个长为 $2$ 的子串，……，$1$ 个长为 $n$ 的子串。每个长为 $k$ 的子串需要 $\mathcal{O}(10^{k/2})$ 的时间判断质数，累加得
   $$  
   S = n\cdot 10^{1/2} + (n-1)\cdot 10^{2/2} + \cdots + 1\cdot 10^{n/2}
   $$
   用**错位相减法**可以求出 $S=\mathcal{O}(10^{n/2})$。详细推导过程请看 [视频讲解](https://www.bilibili.com/video/BV1cqjgzdEPP/)。
- 空间复杂度：$\mathcal{O}(n^2)$。

更多相似题目，见下面数学题单的「**§1.1 判断质数**」。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
