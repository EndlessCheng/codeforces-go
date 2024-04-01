根据题意，当 $n$ 为奇数时，答案为 $2n$，当 $n$ 为偶数时，答案为 $n$。

因此答案为

$$
(n\bmod 2 + 1) \cdot n
$$

```py [sol-Python3]
class Solution:
    def smallestEvenMultiple(self, n: int) -> int:
        return (n % 2 + 1) * n
```

```java [sol-Java]
class Solution {
    public int smallestEvenMultiple(int n) {
        return (n % 2 + 1) * n;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int smallestEvenMultiple(int n) {
        return (n % 2 + 1) * n;
    }
};
```

```go [sol-Go]
func smallestEvenMultiple(n int) int {
	return (n%2 + 1) * n
}
```

```js [sol-JavaScript]
var smallestEvenMultiple = function(n) {
    return (n % 2 + 1) * n;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn smallest_even_multiple(n: i32) -> i32 {
        (n % 2 + 1) * n
    }
}
```

也可以看成是 $n$ 为奇数时，$n$ 左移一位，否则不变。因此可以用位运算解决。

```py [sol-Python3]
class Solution:
    def smallestEvenMultiple(self, n: int) -> int:
        return n << (n & 1)
```

```java [sol-Java]
class Solution {
    public int smallestEvenMultiple(int n) {
        return n << (n & 1);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int smallestEvenMultiple(int n) {
        return n << (n & 1);
    }
};
```

```go [sol-Go]
func smallestEvenMultiple(n int) int {
	return n << (n & 1)
}
```

```js [sol-JavaScript]
var smallestEvenMultiple = function(n) {
    return n << (n & 1);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn smallest_even_multiple(n: i32) -> i32 {
        n << (n & 1)
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
