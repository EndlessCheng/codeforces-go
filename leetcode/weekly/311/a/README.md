根据题意，当 $n$ 为奇数时，答案为 $2n$，当 $n$ 为偶数时，答案为 $n$。

因此答案为

$$
(n\bmod 2 + 1) \cdot n
$$

```py [sol1-Python3]
class Solution:
    def smallestEvenMultiple(self, n: int) -> int:
        return (n % 2 + 1) * n
```

```java [sol1-Java]
class Solution {
    public int smallestEvenMultiple(int n) {
        return (n % 2 + 1) * n;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int smallestEvenMultiple(int n) {
        return (n % 2 + 1) * n;
    }
};
```

```go [sol1-Go]
func smallestEvenMultiple(n int) int {
	return (n%2 + 1) * n
}
```

也可以看成是 $n$ 为奇数时，$n$ 左移一位，否则不变。因此可以用位运算解决。

```py [sol2-Python3]
class Solution:
    def smallestEvenMultiple(self, n: int) -> int:
        return n << (n & 1)
```

```java [sol2-Java]
class Solution {
    public int smallestEvenMultiple(int n) {
        return n << (n & 1);
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    int smallestEvenMultiple(int n) {
        return n << (n & 1);
    }
};
```

```go [sol2-Go]
func smallestEvenMultiple(n int) int {
	return n << (n & 1)
}
```

#### 复杂度分析

- 时间复杂度：$O(1)$。
- 空间复杂度：$O(1)$。

---

欢迎关注[ biIibiIi@灵茶山艾府](https://space.bilibili.com/206214)，高质量算法教学，持续输出中~
