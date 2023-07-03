贪心，按照 $1,0,-1$ 的顺序选：

- 先选 $1$，如果 $k\le \textit{numOnes}$ 那么答案就是 $k$。
- 再选 $0$，如果 $k\le \textit{numOnes}+\textit{numZeros}$ 那么答案为 $\textit{numOnes}$。
- 最后选 $-1$（题目要求恰好选 $k$ 个），那么剩余必须选 $k-\textit{numOnes}-\textit{numZeros}$ 个 $-1$，答案为

$$
\textit{numOnes} + (k-\textit{numOnes}-\textit{numZeros})  \cdot  (-1)= \textit{numOnes} \cdot 2 + \textit{numZeros} - k
$$

```py [sol-Python3]
class Solution:
    def kItemsWithMaximumSum(self, numOnes: int, numZeros: int, _: int, k: int) -> int:
        if k <= numOnes + numZeros:
            return min(k, numOnes)
        return numOnes * 2 + numZeros - k
```

```java [sol-Java]
class Solution {
    public int kItemsWithMaximumSum(int numOnes, int numZeros, int numNegOnes, int k) {
        if (k <= numOnes + numZeros)
            return Math.min(k, numOnes);
        return numOnes * 2 + numZeros - k;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int kItemsWithMaximumSum(int numOnes, int numZeros, int _, int k) {
        if (k <= numOnes + numZeros)
            return min(k, numOnes);
        return numOnes * 2 + numZeros - k;
    }
};
```

```go [sol-Go]
func kItemsWithMaximumSum(numOnes, numZeros, _, k int) int {
	if k <= numOnes {
		return k
	}
	if k <= numOnes+numZeros {
		return numOnes
	}
	return numOnes*2 + numZeros - k
}
```

```js [sol-JavaScript]
var kItemsWithMaximumSum = function(numOnes, numZeros, _, k) {
    if (k <= numOnes + numZeros)
        return Math.min(k, numOnes);
    return numOnes * 2 + numZeros - k;
};
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。
