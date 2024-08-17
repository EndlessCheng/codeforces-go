核心思路：逐个遍历 $\textit{num}$ 的每个数位，统计能整除 $\textit{num}$ 的数位个数。

代码实现时，可以不用转成字符串处理，而是不断取最低位（模 $10$），去掉最低位（除以 $10$），直到数字为 $0$。

例如 $\textit{num}=123$：

1. 初始化 $x=\textit{num}$。
2. 通过 $x\bmod 10$ 取到个位数 $3$，然后把 $x$ 除以 $10$（下取整），得到 $x=12$。
3. 再次 $x\bmod 10$ 取到十位数 $2$，然后把 $x$ 除以 $10$（下取整），得到 $x=1$。
4. 最后 $x\bmod 10$ 取到百位数 $1$，然后把 $x$ 除以 $10$（下取整），得到 $x=0$。此时完成了遍历 $\textit{num}$ 的每个数位，退出循环。
5. 在这个过程中，设取到的数位为 $d$，每次遇到 $\textit{num}\bmod d = 0$ 的情况，就把答案加一。

```py [sol-Python3]
class Solution:
    def countDigits(self, num: int) -> int:
        ans = 0
        x = num
        while x:
            ans += num % (x % 10) == 0
            x //= 10
        return ans
```

```py [sol-Python3 写法二]
class Solution:
    def countDigits(self, num: int) -> int:
        ans = 0
        x = num
        while x:
            x, d = divmod(x, 10)
            ans += num % d == 0
        return ans
```

```java [sol-Java]
class Solution {
    public int countDigits(int num) {
        int ans = 0;
        for (int x = num; x != 0; x /= 10) {
            ans += num % (x % 10) == 0 ? 1 : 0;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countDigits(int num) {
        int ans = 0;
        for (int x = num; x; x /= 10) {
            ans += num % (x % 10) == 0;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countDigits(num int) (ans int) {
	for x := num; x > 0; x /= 10 {
		if num%(x%10) == 0 {
			ans++
		}
	}
	return
}
```

```js [sol-JavaScript]
var countDigits = function(num) {
    let ans = 0;
    for (let x = num; x != 0; x = Math.floor(x / 10)) {
        ans += num % (x % 10) === 0 ? 1 : 0;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_digits(num: i32) -> i32 {
        let mut ans = 0;
        let mut x = num;
        while x != 0 {
            if num % (x % 10) == 0 {
                ans += 1;
            }
            x /= 10;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log \textit{num})$，即 $\textit{num}$ 的十进制长度。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干变量。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
