对于数字 $v=123$，我们可以通过模 $10$ 除以 $10$ 的方式，从低到高遍历其各个数位：

- $123\bmod10 =3$，得到个位数，然后更新 $v$ 为 $\left\lfloor\dfrac{v}{10}\right\rfloor=12$。
- $12\bmod10 =2$，得到十位数，然后更新 $v$ 为 $\left\lfloor\dfrac{12}{10}\right\rfloor=1$。
- $1\bmod10 =1$，得到百位数，然后更新 $v$ 为 $\left\lfloor\dfrac{1}{10}\right\rfloor=0$。遍历结束。

```py [sol-Python3]
class Solution:
    def sumOfTheDigitsOfHarshadNumber(self, x: int) -> int:
        s = 0
        v = x
        while v:
            v, d = divmod(v, 10)
            s += d
        return -1 if x % s else s
```

```java [sol-Java]
class Solution {
    public int sumOfTheDigitsOfHarshadNumber(int x) {
        int s = 0;
        for (int v = x; v > 0; v /= 10) {
            s += v % 10;
        }
        return x % s > 0 ? -1 : s;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int sumOfTheDigitsOfHarshadNumber(int x) {
        int s = 0;
        for (int v = x; v; v /= 10) {
            s += v % 10;
        }
        return x % s ? -1 : s;
    }
};
```

```c [sol-C]
int sumOfTheDigitsOfHarshadNumber(int x) {
    int s = 0;
    for (int v = x; v; v /= 10) {
        s += v % 10;
    }
    return x % s ? -1 : s;
}
```

```go [sol-Go]
func sumOfTheDigitsOfHarshadNumber(x int) int {
	s := 0
	for v := x; v > 0; v /= 10 {
		s += v % 10
	}
	if x%s == 0 {
		return s
	}
	return -1
}
```

```js [sol-JavaScript]
var sumOfTheDigitsOfHarshadNumber = function(x) {
    let s = 0;
    for (let v = x; v; v = Math.floor(v / 10)) {
        s += v % 10;
    }
    return x % s ? -1 : s;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn sum_of_the_digits_of_harshad_number(x: i32) -> i32 {
        let mut s = 0;
        let mut v = x;
        while v > 0 {
            s += v % 10;
            v /= 10;
        }
        if x % s == 0 { s } else { -1 }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log x)$。循环次数为 $x$ 的十进制长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
