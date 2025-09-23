## 方法一：模拟

本题和 [1518. 换水问题](https://leetcode.cn/problems/water-bottles/description/) 几乎一样，唯一区别是每次循环要把 $\textit{numExchange}$ 加一。

```py [sol-Python3]
class Solution:
    def maxBottlesDrunk(self, numBottles: int, numExchange: int) -> int:
        ans = 0
        while numBottles >= numExchange:
            ans += numExchange  # 吨吨吨~
            numBottles -= numExchange - 1
            numExchange += 1
        return ans + numBottles
```

```java [sol-Java]
class Solution {
    public int maxBottlesDrunk(int numBottles, int numExchange) {
        int ans = 0;
        while (numBottles >= numExchange) {
            ans += numExchange; // 吨吨吨~
            numBottles -= numExchange - 1;
            numExchange++;
        }
        return ans + numBottles;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxBottlesDrunk(int numBottles, int numExchange) {
        int ans = 0;
        while (numBottles >= numExchange) {
            ans += numExchange; // 吨吨吨~
            numBottles -= numExchange - 1;
            numExchange++;
        }
        return ans + numBottles;
    }
};
```

```c [sol-C]
int maxBottlesDrunk(int numBottles, int numExchange) {
    int ans = 0;
    while (numBottles >= numExchange) {
        ans += numExchange; // 吨吨吨~
        numBottles -= numExchange - 1;
        numExchange++;
    }
    return ans + numBottles;
}
```

```go [sol-Go]
func maxBottlesDrunk(numBottles, numExchange int) (ans int) {
	for numBottles >= numExchange {
		ans += numExchange // 吨吨吨~
		numBottles -= numExchange - 1
		numExchange++;
	}
	return ans + numBottles
}
```

```js [sol-JavaScript]
var maxBottlesDrunk = function(numBottles, numExchange) {
    let ans = 0;
    while (numBottles >= numExchange) {
        ans += numExchange; // 吨吨吨~
        numBottles -= numExchange - 1;
        numExchange++;
    }
    return ans + numBottles;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn max_bottles_drunk(mut num_bottles: i32, mut num_exchange: i32) -> i32 {
        let mut ans = 0;
        while num_bottles >= num_exchange {
            ans += num_exchange; // 吨吨吨~
            num_bottles -= num_exchange - 1;
            num_exchange += 1;
        }
        ans + num_bottles
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\sqrt \textit{numBottles})$。$\textit{numExchange}=1$ 为最坏情况，根据方法二，需要循环 $\mathcal{O}(\sqrt \textit{numBottles})$ 次。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：数学公式

设 $n = \textit{numBottles}$，$e = \textit{numExchange}$。

设 $k$ 为循环次数（兑换得到的水瓶数），那么答案就是 $n+k$。

我们需要计算最小的 $k$，满足

$$
n - ((e-1) + e + (e+1) + \cdots + (e+k-2)) < e + k
$$

利用等差数列求和公式，上式化简为

$$
k^2 + (2e-1) k - 2(n-e) > 0
$$

解得

$$
k > \dfrac{-(2e-1) + \sqrt{(2e-1)^2+8(n-e)}}{2}
$$

设 $b = 2e-1$。由于 $k$ 是整数，所以 $k$ 最小为

$$
\left\lfloor\dfrac{\sqrt{b^2+8(n-e)} - b}{2}\right\rfloor + 1 = \left\lfloor\dfrac{\sqrt{b^2+8(n-e)} - b + 2}{2}\right\rfloor
$$

为了减少浮点运算次数，减少舍入误差，根据 [下取整恒等式及其应用](https://zhuanlan.zhihu.com/p/1893240318645732760)，上式等于

$$
\left\lfloor\dfrac{\lfloor\sqrt{b^2+8(n-e)}\rfloor - b + 2}{2}\right\rfloor
$$

> 注：部分编程语言是向零取整的，必须把 $+1$ 放到分数中，否则计算出负数，向零取整就是向上取整了。

```py [sol-Python3]
class Solution:
    def maxBottlesDrunk(self, n: int, e: int) -> int:
        b = e * 2 - 1
        k = (isqrt(b * b + (n - e) * 8) - b + 2) // 2
        return n + k
```

```java [sol-Java]
class Solution {
    public int maxBottlesDrunk(int n, int e) {
        int b = e * 2 - 1;
        int k = ((int) Math.sqrt(b * b + (n - e) * 8) - b + 2) / 2;
        return n + k;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxBottlesDrunk(int n, int e) {
        int b = e * 2 - 1;
        int k = ((int) sqrt(b * b + (n - e) * 8) - b + 2) / 2;
        return n + k;
    }
};
```

```c [sol-C]
int maxBottlesDrunk(int n, int e) {
    int b = e * 2 - 1;
    int k = ((int) sqrt(b * b + (n - e) * 8) - b + 2) / 2;
    return n + k;
}
```

```go [sol-Go]
func maxBottlesDrunk(n, e int) int {
	b := e*2 - 1
	k := (int(math.Sqrt(float64(b*b+(n-e)*8))) - b + 2) / 2
	return n + k
}
```

```js [sol-JavaScript]
var maxBottlesDrunk = function(n, e) {
    const b = e * 2 - 1;
    const k = Math.floor((Math.sqrt(b * b + (n - e) * 8) - b) / 2) + 1;
    return n + k;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn max_bottles_drunk(n: i32, e: i32) -> i32 {
        let b = e * 2 - 1;
        let delta = b * b + (n - e) * 8;
        let k = ((delta as f64).sqrt() as i32 - b + 2) / 2;
        n + k
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
