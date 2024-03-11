由于 $n \bmod (n-1) = 1$ 一定满足要求，我们可以从 $n$ 开始，不断生成 $n-1,n-2,\cdots$，最后 $[2,n]$ 中的数字都会在桌面上，这有 $n-1$ 个。

注意特判 $n=1$ 的情况，此时答案为 $1$。

```py [sol-Python3]
class Solution:
    def distinctIntegers(self, n: int) -> int:
        return n - 1 if n > 1 else 1  # max(n - 1, 1)
```

```java [sol-Java]
class Solution {
    public int distinctIntegers(int n) {
        return Math.max(n - 1, 1);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int distinctIntegers(int n) {
        return max(n - 1, 1);
    }
};
```

```go [sol-Go]
func distinctIntegers(n int) int {
	return max(n-1, 1)
}
```

```js [sol-JavaScript]
var distinctIntegers = function(n) {
    return Math.max(n - 1, 1);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn distinct_integers(n: i32) -> i32 {
        1.max(n - 1)
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
