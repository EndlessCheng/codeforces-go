和 [2582. 递枕头](https://leetcode.cn/problems/pass-the-pillow/) 是一样的，那题下标从 $1$ 开始，本题下标从 $0$ 开始。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1h7421R78s/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def numberOfChild(self, n: int, k: int) -> int:
        k, t = divmod(k, n - 1)
        return n - t - 1 if k % 2 else t
```

```java [sol-Java]
class Solution {
    public int numberOfChild(int n, int k) {
        int t = k % (n - 1);
        return k / (n - 1) % 2 > 0 ? n - t - 1 : t;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfChild(int n, int k) {
        int t = k % (n - 1);
        return k / (n - 1) % 2 ? n - t - 1 : t;
    }
};
```

```go [sol-Go]
func numberOfChild(n, k int) int {
	t := k % (n - 1)
	if k/(n-1)%2 > 0 {
		return n - t - 1
	}
	return t
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
