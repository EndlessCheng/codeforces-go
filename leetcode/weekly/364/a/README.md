由于奇数的二进制末尾一定是 $1$，我们可以把一个 $1$ 放在末尾，其余的 $1$ 全部放在开头，这样构造出的奇数尽量大。

```py [sol-Python3]
class Solution:
    def maximumOddBinaryNumber(self, s: str) -> str:
        cnt1 = s.count('1')
        return '1' * (cnt1 - 1) + '0' * (len(s) - cnt1) + '1'
```

```java [sol-Java]
public class Solution {
    public String maximumOddBinaryNumber(String s) {
        int cnt1 = (int) s.chars().filter(c -> c == '1').count();
        return "1".repeat(cnt1 - 1) + "0".repeat(s.length() - cnt1) + "1";
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string maximumOddBinaryNumber(string s) {
        int cnt1 = ranges::count(s, '1');
        return string(cnt1 - 1, '1') + string(s.length() - cnt1, '0') + '1';
    }
};
```

```go [sol-Go]
func maximumOddBinaryNumber(s string) string {
	cnt1 := strings.Count(s, "1")
	return strings.Repeat("1", cnt1-1) + strings.Repeat("0", len(s)-cnt1) + "1"
}
```

```js [sol-JavaScript]
var maximumOddBinaryNumber = function (s) {
    const cnt1 = _.sumBy(s, c => c === '1')
    return '1'.repeat(cnt1 - 1) + '0'.repeat(s.length - cnt1) + '1';
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_odd_binary_number(s: String) -> String {
        let cnt1 = s.chars().filter(|&c| c == '1').count();
        "1".repeat(cnt1 - 1) + &"0".repeat(s.len() - cnt1) + "1"
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
