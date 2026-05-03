用筛法（例如埃氏筛）预处理每个数是不是质数，然后把非质数视作 $0$，计算 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

算出 $r$ 后，用前缀和可以 $\mathcal{O}(1)$ 求出 $[\min(n,r),\max(n,r)]$ 中的质数之和。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
MX = 1001
is_prime = [0, 0] + [1] * (MX - 2)
for i in range(2, isqrt(MX) + 1):
    if is_prime[i]:
        for j in range(i * i, MX, i):
            is_prime[j] = 0

# 原地计算 is_prime 的质数前缀和
for i in range(1, MX):
    is_prime[i] = is_prime[i - 1] + (i if is_prime[i] else 0)

class Solution:
    def sumOfPrimesInRange(self, n: int) -> int:
        r = 0
        x = n
        while x > 0:
            r = r * 10 + x % 10
            x //= 10
        return is_prime[max(n, r)] - is_prime[min(n, r) - 1]
```

```java [sol-Java]
class Solution {
    private static final int MX = 1001;
    private static final int[] isPrime = new int[MX];
    private static boolean initialized = false;

    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        Arrays.fill(isPrime, 2, MX, 1);
        for (int i = 2; i * i < MX; i++) {
            if (isPrime[i] > 0) {
                for (int j = i * i; j < MX; j += i) {
                    isPrime[j] = 0;
                }
            }
        }

        // 原地计算 isPrime 的质数前缀和
        for (int i = 1; i < MX; i++) {
            isPrime[i] = isPrime[i - 1] + (isPrime[i] > 0 ? i : 0);
        }
    }

    public int sumOfPrimesInRange(int n) {
        int r = 0;
        for (int x = n; x > 0; x /= 10) {
            r = r * 10 + x % 10;
        }
        return isPrime[Math.max(n, r)] - isPrime[Math.min(n, r) - 1];
    }
}
```

```cpp [sol-C++]
constexpr int MX = 1001;
int is_prime[MX];

auto init = [] {
    fill(is_prime + 2, is_prime + MX, 1);
    for (int i = 2; i * i < MX; i++) {
        if (is_prime[i]) {
            for (int j = i * i; j < MX; j += i) {
                is_prime[j] = 0;
            }
        }
    }

    // 原地计算 is_prime 的质数前缀和
    for (int i = 1; i < MX; i++) {
        is_prime[i] = is_prime[i - 1] + (is_prime[i] ? i : 0);
    }
    return 0;
}();

class Solution {
public:
    int sumOfPrimesInRange(int n) {
        int r = 0;
        for (int x = n; x > 0; x /= 10) {
            r = r * 10 + x % 10;
        }
        return is_prime[max(n, r)] - is_prime[min(n, r) - 1];
    }
};
```

```go [sol-Go]
const mx = 1001
var isPrime [mx]int

func init() {
	for i := 2; i < mx; i++ {
		isPrime[i] = 1
	}
	for i := 2; i*i < mx; i++ {
		if isPrime[i] > 0 {
			for j := i * i; j < mx; j += i {
				isPrime[j] = 0
			}
		}
	}

	// 原地计算 isPrime 的质数前缀和
	for i := 1; i < mx; i++ {
		if isPrime[i] > 0 {
			isPrime[i] = isPrime[i-1] + i
		} else {
			isPrime[i] = isPrime[i-1]
		}
	}
}

func sumOfPrimesInRange(n int) int {
	r := 0
	for x := n; x > 0; x /= 10 {
		r = r*10 + x%10
	}
	return isPrime[max(n, r)] - isPrime[min(n, r)-1]
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

1. 数学题单的「**§1.2 预处理质数（筛质数）**」。
2. 数据结构题单的「**一、前缀和**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
