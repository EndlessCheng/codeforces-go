## 分析

因为 $10$ 的倍数不可能等于 $115$，所以面额为 $10$ 的硬币不能单独选，至少要选 $1$ 个面额为 $75$ 的硬币。

又由于 $75\cdot 2=150>115$，所以面额为 $75$ 的硬币要**恰好**选 $1$ 个。

由于 $75+10\cdot 4 = 115$，所以面额为 $10$ 的硬币要**恰好**选 $4$ 个。

> 本质上来说，我们在求解二元一次不定方程 $75a+10b=115$，它有唯一正整数解 $a=1,b=4$。

如果一开始 Alice 就没法选，或者偶数轮后 Alice 没法选，那么 Bob 胜出，否则 Alice 胜出。

## 结论

设 $k = \min(x, \lfloor y/4 \rfloor)$，这是能玩的回合数。

判断 $k$ 的奇偶性，奇数 Alice 胜，偶数 Bob 胜。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1JE4m1d7br/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def losingPlayer(self, x: int, y: int) -> str:
        return "Alice" if min(x, y // 4) % 2 else "Bob"
```

```java [sol-Java]
class Solution {
    public String losingPlayer(int x, int y) {
        return Math.min(x, y / 4) % 2 != 0 ? "Alice" : "Bob";
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string losingPlayer(int x, int y) {
        return min(x, y / 4) % 2 ? "Alice" : "Bob";
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

const char* losingPlayer(int x, int y) {
    return MIN(x, y / 4) % 2 ? "Alice" : "Bob";
}
```

```go [sol-Go]
func losingPlayer(x, y int) string {
	return [2]string{"Bob", "Alice"}[min(x, y/4)%2]
}
```

```js [sol-JavaScript]
var losingPlayer = function(x, y) {
    return Math.min(x, Math.floor(y / 4)) % 2 ? "Alice" : "Bob";
};
```

```rust [sol-Rust]
impl Solution {
    pub fn losing_player(x: i32, y: i32) -> String {
        if x.min(y / 4) % 2 != 0 {
            "Alice".to_string()
        } else {
            "Bob".to_string()
        }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
