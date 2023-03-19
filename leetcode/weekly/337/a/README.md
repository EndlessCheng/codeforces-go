### 视频讲解

见[【周赛 337】](https://www.bilibili.com/video/BV1EL411C7YU/)。

# 方法一：二进制基本操作

不断取最低位，然后右移，直到等于 $0$ 为止，这样可以取到每个比特位。

```py [sol1-Python3]
class Solution:
    def evenOddBit(self, n: int) -> List[int]:
        ans = [0, 0]
        i = 0
        while n:
            ans[i] += n & 1
            n >>= 1
            i ^= 1
        return ans
```

```java [sol1-Java]
class Solution {
    public int[] evenOddBit(int n) {
        var ans = new int[2];
        for (int i = 0; n > 0; i ^= 1, n >>= 1)
            ans[i] += n & 1;
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<int> evenOddBit(int n) {
        vector<int> ans(2);
        for (int i = 0; n; i ^= 1, n >>= 1)
            ans[i] += n & 1;
        return ans;
    }
};
```

```go [sol1-Go]
func evenOddBit(n int) []int {
	ans := make([]int, 2)
	for i := 0; n > 0; i ^= 1 {
		ans[i] += n & 1
		n >>= 1
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$O(\log n)$。
- 空间复杂度：$O(1)$。仅用到若干额外变量。

# 方法二：位掩码 + 库函数

利用位掩码 `0x55555555`（二进制的 $010101\cdots$），取出偶数下标比特和奇数下标比特，分别用库函数统计 $1$ 的个数。

本题由于 $n$ 范围比较小，取 `0x5555` 作为位掩码。

```py [sol2-Python3]
class Solution:
    def evenOddBit(self, n: int) -> List[int]:
        MASK = 0x5555
        return [(n & MASK).bit_count(), (n & (MASK >> 1)).bit_count()]
```

```java [sol2-Java]
class Solution {
    public int[] evenOddBit(int n) {
        final int MASK = 0x5555;
        return new int[]{Integer.bitCount(n & MASK), Integer.bitCount(n & (MASK >> 1))};
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    vector<int> evenOddBit(int n) {
        const int MASK = 0x5555;
        return {__builtin_popcount(n & MASK), __builtin_popcount(n & (MASK >> 1))};
    }
};
```

```go [sol2-Go]
func evenOddBit(n int) []int {
	const mask = 0x5555
	return []int{bits.OnesCount16(uint16(n & mask)), bits.OnesCount16(uint16(n & (mask >> 1)))}
}
```

### 复杂度分析

- 时间复杂度：$O(1)$。
- 空间复杂度：$O(1)$。仅用到若干额外变量。
