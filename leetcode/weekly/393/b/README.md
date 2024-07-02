## 方法一：判断质数

如何判断一个数 $n$ 是不是质数？

在本题数据范围下，可以判断 $n$ 能否能被 $10$ 以内（或者 $\sqrt n$ 以内）的某个大于 $1$ 的数整除，如果不能则说明 $n$ 是质数。

为什么？反证法：如果 $n$ 不能被 $10$ 以内的大于 $1$ 的数整除，但可以被大于 $10$ 的数整除，例如 $11$，那么必然还有一个数 $\dfrac{n}{11}$ 也能整除 $n$，但是 $\dfrac{n}{11}\le \dfrac{100}{11} < 10$，说明有一个小于 $10$ 的数也能整除 $n$，矛盾。

注意 $1$ 不是质数。

```py [sol-Python3]
class Solution:
    def maximumPrimeDifference(self, nums: List[int]) -> int:
        is_prime = lambda n: n >= 2 and all(n % i for i in range(2, isqrt(n) + 1))
        i = 0
        while not is_prime(nums[i]):
            i += 1
        j = len(nums) - 1
        while not is_prime(nums[j]):
            j -= 1
        return j - i
```

```java [sol-Java]
class Solution {
    public int maximumPrimeDifference(int[] nums) {
        int i = 0;
        while (!isPrime(nums[i])) {
            i++;
        }
        int j = nums.length - 1;
        while (!isPrime(nums[j])) {
            j--;
        }
        return j - i;
    }

    private boolean isPrime(int n) {
        for (int i = 2; i * i <= n; i++) {
            if (n % i == 0) {
                return false;
            }
        }
        return n >= 2;
    }
}
```

```cpp [sol-C++]
class Solution {
    bool is_prime(int n) {
        for (int i = 2; i * i <= n; i++) {
            if (n % i == 0) {
                return false;
            }
        }
        return n >= 2;
    }

public:
    int maximumPrimeDifference(vector<int>& nums) {
        int i = 0;
        while (!is_prime(nums[i])) {
            i++;
        }
        int j = nums.size() - 1;
        while (!is_prime(nums[j])) {
            j--;
        }
        return j - i;
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

func maximumPrimeDifference(nums []int) int {
	i := 0
	for !isPrime(nums[i]) {
		i++
	}
	j := len(nums)-1
	for !isPrime(nums[j]) {
		j--
	}
	return j - i
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})\le 100$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：筛质数

遍历 $\sqrt {100} = 10$ 以内的质数 $2,3,5,7$，标记其倍数为合数。

注意标记 $i$ 的倍数时，只需从 $j = i^2$ 开始，因为小于 $i^2$ 的合数 $j$ 一定有一个小于 $i$ 的质因子。可以用反证法证明：如果小于 $i^2$ 的合数 $j$ 没有小于 $i$ 的质因子，那么合数 $j$ 至少是两个 $\ge i$ 的数的乘积，这与 $j < i^2$ 矛盾。

```py [sol-Python3]
not_prime = [False] * 101
not_prime[1] = True
for i in 2, 3, 5, 7:
    for j in range(i * i, 101, i):
        not_prime[j] = True

class Solution:
    def maximumPrimeDifference(self, nums: List[int]) -> int:
        i = 0
        while not_prime[nums[i]]:
            i += 1
        j = len(nums) - 1
        while not_prime[nums[j]]:
            j -= 1
        return j - i
```

```py [sol-Python3 通用写法]
MX = 101
not_prime = [True, True] + [False] * (MX - 2)
for i in range(2, isqrt(MX) + 1):
    if not_prime[i]: continue
    for j in range(i * i, MX, i):
        not_prime[j] = True  # j 是质数 i 的倍数

class Solution:
    def maximumPrimeDifference(self, nums: List[int]) -> int:
        i = 0
        while not_prime[nums[i]]:
            i += 1
        j = len(nums) - 1
        while not_prime[nums[j]]:
            j -= 1
        return j - i
```

```java [sol-Java]
class Solution {
    private static final int MX = 101;
    private static final boolean[] NOT_PRIME = new boolean[MX];

    static {
        NOT_PRIME[1] = true;
        for (int i = 2; i * i < MX; i++) {
            if (NOT_PRIME[i]) continue;
            for (int j = i * i; j < MX; j += i) {
                NOT_PRIME[j] = true; // j 是质数 i 的倍数
            }
        }
    }

    public int maximumPrimeDifference(int[] nums) {
        int i = 0;
        while (NOT_PRIME[nums[i]]) {
            i++;
        }
        int j = nums.length - 1;
        while (NOT_PRIME[nums[j]]) {
            j--;
        }
        return j - i;
    }
}
```

```cpp [sol-C++]
const int MX = 101;
bool not_prime[MX];

auto init = [] {
    not_prime[1] = true;
    for (int i = 2; i * i < MX; i++) {
        if (not_prime[i]) continue;
        for (int j = i * i; j < MX; j += i) {
            not_prime[j] = true; // j 是质数 i 的倍数
        }
    }
    return 0;
}();

class Solution {
public:
    int maximumPrimeDifference(vector<int>& nums) {
        int i = 0;
        while (not_prime[nums[i]]) {
            i++;
        }
        int j = nums.size() - 1;
        while (not_prime[nums[j]]) {
            j--;
        }
        return j - i;
    }
};
```

```go [sol-Go]
const mx = 101
var notPrime = [mx]bool{true, true}
func init() {
	for i := 2; i*i < mx; i++ {
		if !notPrime[i] {
			for j := i * i; j < mx; j += i {
				notPrime[j] = true // j 是质数 i 的倍数
			}
		}
	}
}

func maximumPrimeDifference(nums []int) int {
	i := 0
	for notPrime[nums[i]] {
		i++
	}
	j := len(nums) - 1
	for notPrime[nums[j]] {
		j--
	}
	return j - i
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。忽略预处理的时间和空间。
- 空间复杂度：$\mathcal{O}(1)$。

**注**：从两边向中间找的好处是，在随机数据下，根据质数密度，期望 $\mathcal{O}(\log U)$ 的时间就能找到质数，其中 $U=\max(\textit{nums})\le 100$。

## 方法三：位运算

把 $100$ 内的**质数组成的集合**，「压缩」成两个 $64$ 位整数。（Python 只需要一个 `int`）

原理见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def maximumPrimeDifference(self, nums: List[int]) -> int:
        PRIME_MASK = 0x20208828828208a20a08a28ac
        i = 0
        while PRIME_MASK >> nums[i] & 1 == 0:
            i += 1
        j = len(nums) - 1
        while PRIME_MASK >> nums[j] & 1 == 0:
            j -= 1
        return j - i
```

```java [sol-Java]
class Solution {
    private static final long[] PRIME_MASK = {0x28208a20a08a28acL, 0x202088288L};

    public int maximumPrimeDifference(int[] nums) {
        int i = 0;
        while ((PRIME_MASK[nums[i] / 64] >> (nums[i] % 64) & 1) == 0) {
            i++;
        }
        int j = nums.length - 1;
        while ((PRIME_MASK[nums[j] / 64] >> (nums[j] % 64) & 1) == 0) {
            j--;
        }
        return j - i;
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr uint64_t PRIME_MASK[2]{0x28208a20a08a28ac, 0x202088288};
public:
    int maximumPrimeDifference(vector<int>& nums) {
        int i = 0;
        while ((PRIME_MASK[nums[i] / 64] >> (nums[i] % 64) & 1) == 0) {
            i++;
        }
        int j = nums.size() - 1;
        while ((PRIME_MASK[nums[j] / 64] >> (nums[j] % 64) & 1) == 0) {
            j--;
        }
        return j - i;
    }
};
```

```go [sol-Go]
func maximumPrimeDifference(nums []int) int {
	primeMask := [2]uint{0x28208a20a08a28ac, 0x202088288}
	i := 0
	for primeMask[nums[i]/64]>>(nums[i]%64)&1 == 0 {
		i++
	}
	j := len(nums) - 1
	for primeMask[nums[j]/64]>>(nums[j]%64)&1 == 0 {
		j--
	}
	return j - i
}
```

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
