本题是 [2833. 距离原点最远的点](https://leetcode.cn/problems/furthest-point-from-origin/) 的二维版本，思路是一样的，请看 [我的题解](https://leetcode.cn/problems/furthest-point-from-origin/solutions/2413317/nao-jin-ji-zhuan-wan-yi-xing-dai-ma-by-e-yfn0/)。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def maxDistance(self, moves: str) -> int:
        x = y = free = 0

        for ch in moves:
            if ch == 'L': x -= 1
            elif ch == 'R': x += 1
            elif ch == 'D': y -= 1
            elif ch == 'U': y += 1
            else: free += 1

        return abs(x) + abs(y) + free
```

```java [sol-Java]
class Solution {
    public int maxDistance(String moves) {
        int x = 0;
        int y = 0;
        int free = 0;

        for (char ch : moves.toCharArray()) {
            switch (ch) {
                case 'L': x--; break;
                case 'R': x++; break;
                case 'D': y--; break;
                case 'U': y++; break;
                default: free++;
            }
        }

        return Math.abs(x) + Math.abs(y) + free;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxDistance(string moves) {
        int x = 0, y = 0, free = 0;

        for (char ch : moves) {
            switch (ch) {
                case 'L': x--; break;
                case 'R': x++; break;
                case 'D': y--; break;
                case 'U': y++; break;
                default: free++;
            }
        }

        return abs(x) + abs(y) + free;
    }
};
```

```go [sol-Go]
func maxDistance(moves string) int {
	x, y, free := 0, 0, 0

	for _, ch := range moves {
		switch ch {
		case 'L': x--
		case 'R': x++
		case 'D': y--
		case 'U': y++
		default: free++
		}
	}

	return abs(x) + abs(y) + free
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{moves}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

[3443. K 次修改后的最大曼哈顿距离](https://leetcode.cn/problems/maximum-manhattan-distance-after-k-changes/)

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
