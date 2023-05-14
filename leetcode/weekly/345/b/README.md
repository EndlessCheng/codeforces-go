下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

### 前置知识：异或运算的性质

$a\oplus a = 0$。

$a\oplus b = c$ 两边同时异或 $a$ 可以得到 $b = c\oplus a$。

### 思路

如果知道 $\textit{original}[0]$，利用 $\textit{derived}[i]$ 可以推出其余 $\textit{original}[i]$ 的值，即

$$
\textit{original}[i+1] = \textit{original}[i]\oplus \textit{derived}[i]
$$

那么有 

$$
\textit{original}[n-1] = \textit{original}[0] \oplus \textit{derived}[0] \oplus \textit{derived}[1]\oplus \cdots \oplus \textit{derived}[n-2]
$$

由于 

$$
\textit{original}[0]\oplus \textit{original}[n-1] =\textit{derived}[n-1]
$$

联立得

$$
\textit{derived}[0] \oplus \textit{derived}[1] \oplus\cdots \oplus \textit{derived}[n-1] = 0
$$

所以如果上式成立，$\textit{original}$ 必然存在。

```py [sol1-Python3]
class Solution:
    def doesValidArrayExist(self, derived: List[int]) -> bool:
        return reduce(xor, derived) == 0
```

```java [sol1-Java]
class Solution {
    public boolean doesValidArrayExist(int[] derived) {
        int xor = 0;
        for (int x : derived)
            xor ^= x;
        return xor == 0;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    bool doesValidArrayExist(vector<int> &derived) {
        int xor_ = 0;
        for (int x: derived)
            xor_ ^= x;
        return xor_ == 0;
    }
};
```

```go [sol1-Go]
func doesValidArrayExist(derived []int) bool {
	xor := 0
	for _, x := range derived {
		xor ^= x
	}
	return xor == 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{derived}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
