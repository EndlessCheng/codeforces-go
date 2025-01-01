算法：

1. 把字符串 $\textit{date}$ 按照 `-` 分割。
2. 对于每个字符串，首先转成数字。
3. 然后转成二进制字符串。
4. 最后，用 `-` 拼接每个二进制字符串。

```py [sol-Py3]
class Solution:
    def convertDateToBinary(self, date: str) -> str:
        a = date.split('-')
        for i in range(len(a)):
            a[i] = bin(int(a[i]))[2:]
        return '-'.join(a)
```

```py [sol-Py3 一行]
class Solution:
    def convertDateToBinary(self, date: str) -> str:
        return '-'.join(bin(int(s))[2:] for s in date.split('-'))
```

```java [sol-Java]
class Solution {
    public String convertDateToBinary(String date) {
        String[] a = date.split("-");
        for (int i = 0; i < a.length; i++) {
            a[i] = Integer.toBinaryString(Integer.parseInt(a[i]));
        }
        return String.join("-", a);
    }
}
```

```java [sol-Java Stream]
class Solution {
    public String convertDateToBinary(String date) {
        return Arrays.stream(date.split("-"))
                .map(s -> Integer.toBinaryString(Integer.parseInt(s)))
                .collect(Collectors.joining("-"));
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string convertDateToBinary(string date) {
        return format("{:b}-{:b}-{:b}",
                      stoi(date.substr(0, 4)),
                      stoi(date.substr(5, 2)),
                      stoi(date.substr(8, 2)));
    }
};
```

```go [sol-Go]
func convertDateToBinary(date string) string {
    a := strings.Split(date, "-")
    for i := range a {
        x, _ := strconv.Atoi(a[i])
        a[i] = strconv.FormatUint(uint64(x), 2)
    }
    return strings.Join(a, "-")
}
```

```js [sol-JS]
var convertDateToBinary = function(date) {
    return date.split('-')
               .map(s => parseInt(s, 10).toString(2))
               .join('-');
};
```

```rust [sol-Rust]
impl Solution {
    pub fn convert_date_to_binary(date: String) -> String {
        date.split('-')
            .map(|s| format!("{:b}", s.parse::<u16>().unwrap()))
            .collect::<Vec<_>>()
            .join("-")
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{date}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
