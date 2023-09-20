### 提示 1

从小到大枚举答案。

### 提示 2

假设操作了 $k$ 次，那么操作后 $\textit{num}_1$ 变成 $\textit{num}_1 - \textit{num}_2\cdot k$ 再减去 $k$ 个 $2^i$。

此时问题变成：$\textit{num}_1 - \textit{num}_2\cdot k$ 能否拆分成 $k$ 个 $2^i$ 之和？

### 提示 3

设 $x=\textit{num}_1 - \textit{num}_2\cdot k$。

- 如果 $x<0$，无解。
- 否则如果 $x<k$，那么即使每次操作取 $i=0$，也至少要把 $x$ 拆分成 $k$ 个 $1$ 之和，这是不可能的。
- 否则如果 $x$ 中二进制 $1$ 的个数大于 $k$，也无法拆分成 $k$ 个 $2^i$ 之和，无解。
- 否则分解方案一定存在，返回 $k$。（因为可以把一个 $2^j$ 分解成两个 $2^{j-1}$，所以分解出的 $2^i$ 的**个数**可以从「$x$ 中二进制 $1$ 的个数」一直到 $x$，$k$ 只要属于这个范围，分解方案就是存在的。）

代码实现时，如果出现 $x<k$ 的情况，说明 $\textit{num}_2\ge 0$，那么对于更大的 $k$，$x$ 无法变得更大，所以后面都无解，直接退出循环。在 [视频讲解](https://www.bilibili.com/video/BV1du41187ZN/) 中，我画出了两个关于 $k$ 的一次函数的图像，数形结合，可以更容易地理解这一做法的正确性。

```py [sol-Python3]
class Solution:
    def makeTheIntegerZero(self, num1: int, num2: int) -> int:
        for k in count(1):
            x = num1 - num2 * k
            if x < k: return -1
            if k >= x.bit_count(): return k
```

```java [sol-Java]
class Solution {
    public int makeTheIntegerZero(int num1, int num2) {
        for (long k = 1; k <= num1 - num2 * k; k++)
            if (k >= Long.bitCount(num1 - num2 * k))
                return (int) k;
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int makeTheIntegerZero(int num1, int num2) {
        for (long long k = 1; k <= num1 - num2 * k; k++)
            if (k >= __builtin_popcountll(num1 - num2 * k))
                return k;
        return -1;
    }
};
```

```go [sol-Go]
func makeTheIntegerZero(num1, num2 int) int {
	for k := 1; k <= num1-num2*k; k++ {
		if k >= bits.OnesCount(uint(num1-num2*k)) {
			return k
		}
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(f^{-1}(\textit{num}_1+|\textit{num}_2|))$，其中 $f^{-1}(x)$ 是 $f(x)=\dfrac{2^x}{x}$ 的反函数。具体来说，循环的次数不会超过 $36$ 次，详见 [视频讲解](https://www.bilibili.com/video/BV1du41187ZN/)。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

> 关于这个反函数的研究，见朗伯 W 函数。
