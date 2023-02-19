把 $n$ 看成二进制数，那么更高位的比特 $1$ 是会受到更低位的比特 $1$ 的加减影响的，但是，最小的比特 $1$ 没有这个约束。

那么考虑优先消除最小的比特 $1$，设它对应的数字为 $\textit{lowbit}$。

消除方法只能是加上 $\textit{lowbit}$，或者减去 $\textit{lowbit}$。

$\textit{lowbit}$ 的计算方法见本题 [视频讲解](https://www.bilibili.com/video/BV1jM411J7y7/)。

贪心的策略是：如果有多个连续 $1$，那么采用加法是更优的，可以一次消除多个 $1$；否则对于单个 $1$，减法更优。

```py [sol2-Python3]
class Solution:
    def minOperations(self, n: int) -> int:
        ans = 1
        while n & (n - 1):  # 不是 2 的幂次
            lb = n & -n
            if n & (lb << 1): n += lb  # 多个连续 1
            else: n -= lb  # 单个 1
            ans += 1
        return ans
```

```java [sol2-Java]
class Solution {
    public int minOperations(int n) {
        int ans = 1;
        while ((n & (n - 1)) > 0) { // n 不是 2 的幂次
            int lb = n & -n;
            if ((n & (lb << 1)) > 0) n += lb; // 多个连续 1
            else n -= lb; // 单个 1
            ++ans;
        }
        return ans;
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    int minOperations(int n) {
        int ans = 1;
        while (n & (n - 1)) { // n 不是 2 的幂次
            int lb = n & -n;
            if (n & (lb << 1)) n += lb; // 多个连续 1
            else n -= lb; // 单个 1
            ++ans;
        }
        return ans;
    }
};
```

```go [sol2-Go]
func minOperations(n int) int {
	ans := 1
	for n&(n-1) > 0 { // n 不是 2 的幂次
		lb := n & -n
		if n&(lb<<1) > 0 { // 多个连续 1
			n += lb
		} else {
			n -= lb // 单个 1
		}
		ans++
	}
	return ans
}
```

### 位运算优化

对于多个连续 $1$，如果它和前面的 $1$ 由至少两个 $0$ 隔开的话，那么就需要先加上 $\textit{lowbit}$，产生单个 $1$，再减去 $\textit{lowbit}$ 去掉这个 $1$，那么需要操作两次。

注意到

$$
\begin{aligned} 
n&=00111111\\
3n&=10111101\\
n\oplus 3n&=10000010
\end{aligned}
$$

刚好可以得到两个 $1$（$\oplus$ 表示异或）。

另外，对于单个 $1$，有

$$
\begin{aligned}
n&=0100\\
3n&=1100\\
n\oplus 3n&=1000
\end{aligned}
$$

刚好可以得到一个 $1$。

因此答案就是 $n\oplus 3n$ 二进制中 $1$ 的个数。

```py [sol3-Python3]
class Solution:
    def minOperations(self, n: int) -> int:
        return (3 * n ^ n).bit_count()
```

```java [sol3-Java]
class Solution {
    public int minOperations(int n) {
        return Integer.bitCount(3 * n ^ n);
    }
}
```

```cpp [sol3-C++]
class Solution {
public:
    int minOperations(int n) {
        return __builtin_popcount(3 * n ^ n);
    }
};
```

```go [sol3-Go]
func minOperations(n int) int {
	return bits.OnesCount(uint(3*n ^ n))
}
```

### 复杂度分析

- 时间复杂度：$O(1)$。
- 空间复杂度：$O(1)$。仅用到若干变量。

---

### 附：记忆化搜索写法

理论讲解请看[【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)。

```py [sol1-Python3]
@cache
def dfs(x: int) -> int:
    if (x & (x - 1)) == 0:  # x 是 2 的幂次
        return 1
    lb = x & -x
    return 1 + min(dfs(x + lb), dfs(x - lb))

class Solution:
    def minOperations(self, n: int) -> int:
        return dfs(n)
```

```go [sol1-Go]
var cache = map[int]int{}

func minOperations(n int) int {
	if n&(n-1) == 0 { // n 是 2 的幂次
		return 1
	}
	if res, ok := cache[n]; ok {
		return res
	}
	lb := n & -n
	res := 1 + min(minOperations(n+lb), minOperations(n-lb))
	cache[n] = res
	return res
}

func min(a, b int) int { if a > b { return b }; return a }
```
