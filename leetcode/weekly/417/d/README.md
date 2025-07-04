## 方法一：递归

设 $\textit{operations}$ 的长度为 $n$。

由于每次操作，都会把字符串的长度扩大一倍，所以 $n$ 次操作后，最终字符串 $S$ 的长度为 $2^n$。

分类讨论：

- 如果 $k \le 2^{n-1}$，那么第 $k$ 个字母在 $S$ 的左半段，不会受到 $\textit{operations}[n-1]$ 的影响，问题等价于去掉 $\textit{operations}[n-1]$ 的子问题。
- 如果 $k > 2^{n-1}$，那么第 $k$ 个字母在 $S$ 的右半段，问题等价于去掉 $\textit{operations}[n-1]$，计算右半段的第 $k-2^{n-1}$ 个字母的子问题。如果 $\textit{operations}[n-1]=1$，那么子问题返回的字母需要加 $1$（变成下一个字母），否则不变。也相当于子问题返回的字母需要增加 $\textit{operations}[n-1]$。

递归边界：如果 $n=0$，没有操作，返回 Alice 最初的字母 $\texttt{a}$。 

具体请看 [视频讲解](https://www.bilibili.com/video/BV1TqxCeZEmb/?t=14m29s) 第四题，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def kthCharacter(self, k: int, operations: List[int]) -> str:
        if not operations:
            return 'a'
        op = operations.pop()
        # 注意 pop 之后 operations 的长度减少了 1，所以下面写的是 1<<n 而不是 1<<(n-1)
        m = 1 << len(operations)
        if k <= m:  # k 在左半段
            return self.kthCharacter(k, operations)
        # k 在右半段
        ans = self.kthCharacter(k - m, operations)
        return ascii_lowercase[(ord(ans) - ord('a') + op) % 26]
```

```java [sol-Java]
class Solution {
    public char kthCharacter(long k, int[] operations) {
        // 从 k-1 的二进制长度减一开始，详细解释见下文
        return f(k, operations, 63 - Long.numberOfLeadingZeros(k - 1));
    }

    private char f(long k, int[] operations, int i) {
        if (i < 0) {
            return 'a';
        }
        int op = operations[i];
        if (k <= (1L << i)) { // k 在左半段
            return f(k, operations, i - 1);
        }
        // k 在右半段
        char ans = f(k - (1L << i), operations, i - 1);
        return (char) ('a' + (ans - 'a' + op) % 26);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    char kthCharacter(long long k, vector<int>& operations) {
        if (operations.empty()) {
            return 'a';
        }
        int op = operations.back();
        operations.pop_back();
        int n = operations.size(); // 注意这里是减一后的 n
        if (n >= 63 || k <= (1LL << n)) { // k 在左半段
            return kthCharacter(k, operations);
        }
        // k 在右半段
        char ans = kthCharacter(k - (1LL << n), operations);
        return 'a' + (ans - 'a' + op) % 26;
    }
};
```

```go [sol-Go]
func kthCharacter(k int64, operations []int) byte {
	n := len(operations)
	if n == 0 {
		return 'a'
	}
	n-- // 注意这里减一了
	op := operations[n]
	operations = operations[:n]
	if n >= 63 || k <= 1<<n { // k 在左半段
		return kthCharacter(k, operations)
	}
	// k 在右半段
	ans := kthCharacter(k-1<<n, operations)
	return 'a' + (ans-'a'+byte(op))%26
}
```

## 方法二：迭代

写出上面的递归代码后，可以发现：

1. 本质上，我们在计算 $\texttt{a}$ 需要增加的次数，这可以用一个变量 $\textit{inc}$ 记录。
2. 我们在倒序遍历 $\textit{operations}$。当 $k$ 在字符串的右半段，也就是 $k > 2^i$ 时，我们会把 $\textit{inc}$ 增加 $\textit{operations}[i]$。 

由于 $k > 2^i$ 等价于 $k-1\ge 2^i$，解得

$$
i\le \lfloor \log_2 (k-1) \rfloor
$$

也就是 $i$ 小于等于 $k-1$ 的二进制长度减一。

注意题目保证执行完所有操作后字符串至少有 $k$ 个字母，所以无需担心下标 $i$ 越界的情况。

### 写法一

```py [sol-Python3]
class Solution:
    def kthCharacter(self, k: int, operations: List[int]) -> str:
        m = (k - 1).bit_length()
        inc = 0
        for i in range(m - 1, -1, -1):
            if k > 1 << i:  # k 在右半段
                inc += operations[i]
                k -= 1 << i
        return ascii_lowercase[inc % 26]
```

```java [sol-Java]
class Solution {
    public char kthCharacter(long k, int[] operations) {
        int inc = 0;
        for (int i = 63 - Long.numberOfLeadingZeros(k - 1); i >= 0; i--) {
            if (k > (1L << i)) { // k 在右半段
                inc += operations[i];
                k -= 1L << i;
            }
        }
        return (char) ('a' + inc % 26);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    char kthCharacter(long long k, vector<int>& operations) {
        int m = bit_width((uint64_t) k - 1);
        int inc = 0;
        for (int i = m - 1; i >= 0; i--) {
            if (k > (1LL << i)) { // k 在右半段
                inc += operations[i];
                k -= 1LL << i;
            }
        }
        return 'a' + inc % 26;
    }
};
```

```go [sol-Go]
func kthCharacter(k int64, operations []int) byte {
	inc := 0
	for i := bits.Len(uint(k-1)) - 1; i >= 0; i-- {
		if k > 1<<i { // k 在右半段
			inc += operations[i]
			k -= 1 << i
		}
	}
	return 'a' + byte(inc%26)
}
```

### 写法二

本质上，我们相当于在遍历 $k-1$ 二进制的每个比特 $1$，累加 $1$ 对应的 $\textit{operations}[i]$。

```py [sol-Python3]
class Solution:
    def kthCharacter(self, k: int, operations: List[int]) -> str:
        k -= 1
        m = k.bit_length()
        inc = sum(op for i, op in enumerate(operations[:m]) if k >> i & 1)
        return ascii_lowercase[inc % 26]
```

```java [sol-Java]
class Solution {
    public char kthCharacter(long k, int[] operations) {
        k--;
        int inc = 0;
        for (int i = 63 - Long.numberOfLeadingZeros(k); i >= 0; i--) {
            if ((k >> i & 1) > 0) {
                inc += operations[i];
            }
        }
        return (char) ('a' + inc % 26);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    char kthCharacter(long long k, vector<int>& operations) {
        k--;
        int m = bit_width((uint64_t) k);
        int inc = 0;
        for (int i = m - 1; i >= 0; i--) {
            if (k >> i & 1) {
                inc += operations[i];
            }
        }
        return 'a' + inc % 26;
    }
};
```

```go [sol-Go]
func kthCharacter(k int64, operations []int) byte {
	k--
	inc := 0
	for i, op := range operations[:bits.Len(uint(k))] {
		if k>>i&1 > 0 {
			inc += op
		}
	}
	return 'a' + byte(inc%26)
}
```

### 写法三

当「$k-1$ 第 $i$ 位是 $1$」和「$\textit{operations}[i]=1$」两个条件同时满足时，才把 $\textit{inc}$ 加一。

这两个条件可以合并为：把 $k-1$ 右移 $i$ 位，与 $\textit{operations}[i]$ 计算 AND，如果结果是 $1$，把 $\textit{inc}$ 加一。也可以直接把结果加到 $\textit{inc}$ 中。

```py [sol-Python3]
class Solution:
    def kthCharacter(self, k: int, operations: List[int]) -> str:
        k -= 1
        m = k.bit_length()
        inc = sum(k >> i & op for i, op in enumerate(operations[:m]))
        return ascii_lowercase[inc % 26]
```

```java [sol-Java]
class Solution {
    public char kthCharacter(long k, int[] operations) {
        k--;
        int inc = 0;
        for (int i = 63 - Long.numberOfLeadingZeros(k); i >= 0; i--) {
            inc += k >> i & operations[i];
        }
        return (char) ('a' + inc % 26);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    char kthCharacter(long long k, vector<int>& operations) {
        k--;
        int m = bit_width((uint64_t) k);
        int inc = 0;
        for (int i = m - 1; i >= 0; i--) {
            inc += k >> i & operations[i];
        }
        return 'a' + inc % 26;
    }
};
```

```go [sol-Go]
func kthCharacter(k int64, operations []int) byte {
	k--
	inc := 0
	for i, op := range operations[:bits.Len(uint(k))] {
		inc += int(k) >> i & op
	}
	return 'a' + byte(inc%26)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log k)$。注意题目保证 $\textit{operations}$ 数组足够长。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

- [1545. 找出第 N 个二进制字符串中的第 K 位](https://leetcode.cn/problems/find-kth-bit-in-nth-binary-string/) 做到 $\mathcal{O}(\log k)$

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
