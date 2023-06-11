## 方法一：模拟

根据题意，$3n$ 必须是一个三位数，即 $3n\le 999$，所以 $n\le 333$。

同时，由于 $n$ 中不能有 $0$ 和重复数字，所以如果 $n<123$ 或者 $n>329$，直接返回 `false`。

否则可以用一个哈希表，或者二进制数 $\textit{mask}$ 记录有哪些数字出现过，最后判断 $\textit{mask}$ 是否只有 $1$ 到 $9$，即 $2^{10}-2=1111111110_{(2)}$。

```py [sol-Python3]
class Solution:
    def isFascinating(self, n: int) -> bool:
        if n < 123 or n > 329:
            return False
        s = str(n) + str(n * 2) + str(n * 3)
        return '0' not in s and len(set(s)) == 9
```

```java [sol-Java]
class Solution {
    public boolean isFascinating(int n) {
        if (n < 123 || n > 329) return false;
        int mask = 0;
        for (char c : ("" + n + (n * 2) + (n * 3)).toCharArray())
            mask |= 1 << (c - '0');
        return mask == (1 << 10) - 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isFascinating(int n) {
        if (n < 123 || n > 329) return false;
        int mask = 0;
        for (char c: to_string(n) + to_string(n * 2) + to_string(n * 3))
            mask |= 1 << (c - '0');
        return mask == (1 << 10) - 2;
    }
};
```

```go [sol-Go]
func isFascinating(n int) bool {
	if n < 123 || n > 329 {
		return false
	}
	mask := 0
	for _, c := range strconv.Itoa(n) + strconv.Itoa(n*2) + strconv.Itoa(n*3) {
		mask |= 1 << (c - '0')
	}
	return mask == 1<<10-2
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(\log n)$。

## 方法二：打表

遍历所有三位数 $n$，发现只有如下 $4$ 个数满足要求：

$$
192,219,273,327
$$

```py [sol-Python3]
class Solution:
    def isFascinating(self, n: int) -> bool:
        return n in (192, 219, 273, 327)
```

```java [sol-Java]
class Solution {
    public boolean isFascinating(int n) {
        return n == 192 || n == 219 || n == 273 || n == 327;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isFascinating(int n) {
        return n == 192 || n == 219 || n == 273 || n == 327;
    }
};
```

```go [sol-Go]
func isFascinating(n int) bool {
	return n == 192 || n == 219 || n == 273 || n == 327
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。
