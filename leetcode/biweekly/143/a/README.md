至多循环 $10$ 次，一定会遇到个位数为 $0$ 的数字，数位乘积是 $0$，一定是 $t$ 的倍数。

所以暴力枚举即可。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1cgmBYqEhu/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def smallestNumber(self, n: int, t: int) -> int:
        for x in count(n):
            prod = reduce(mul, map(int, str(x)))
            if prod % t == 0:
                return x
```

```py [sol-Python3 不用字符串]
class Solution:
    def smallestNumber(self, n: int, t: int) -> int:
        for i in count(n):
            prod = 1
            x = i
            while x:
                x, d = divmod(x, 10)
                prod *= d
            if prod % t == 0:
                return i
```

```java [sol-Java]
class Solution {
    int smallestNumber(int n, int t) {
        for (int i = n; ; i++) {
            int prod = 1;
            for (int x = i; x > 0; x /= 10) {
                prod *= x % 10;
            }
            if (prod % t == 0) {
                return i;
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int smallestNumber(int n, int t) {
        for (int i = n; ; i++) {
            int prod = 1;
            for (int x = i; x; x /= 10) {
                prod *= x % 10;
            }
            if (prod % t == 0) {
                return i;
            }
        }
    }
};
```

```go [sol-Go]
func smallestNumber(n, t int) int {
	for i := n; ; i++ {
		prod := 1
		for x := i; x > 0; x /= 10 {
			prod *= x % 10
		}
		if prod%t == 0 {
			return i
		}
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。
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
