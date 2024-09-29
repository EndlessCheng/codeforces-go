本题相当于周赛第四题所有 $\textit{operations}[i]=1$ 的版本，做法是一样的，请先看 [我的题解](https://leetcode.cn/problems/find-the-k-th-character-in-string-game-ii/solutions/2934284/liang-chong-zuo-fa-di-gui-die-dai-python-5f6z/)。

## 优化前

为了方便大家看出怎么优化，代码中先把 $k$ 减一。

```py [sol-Python3]
class Solution:
    def kthCharacter(self, k: int) -> str:
        k -= 1
        inc = 0
        for i in range(k.bit_length() - 1, -1, -1):
            if k >= 1 << i:  # k 在右半边
                inc += 1
                k -= 1 << i
        return ascii_lowercase[inc]
```

```java [sol-Java]
class Solution {
    public char kthCharacter(int k) {
        k--;
        char ans = 'a';
        for (int i = 31 - Integer.numberOfLeadingZeros(k); i >= 0; i--) {
            if (k >= (1 << i)) { // k 在右半边
                ans++;
                k -= (1 << i);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    char kthCharacter(long long k) {
        k--;
        char ans = 'a';
        for (int i = __lg(k); i >= 0; i--) {
            if (k >= (1 << i)) { // k 在右半边
                ans++;
                k -= (1 << i);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func kthCharacter(k int) byte {
	k--
	ans := byte('a')
	for i := bits.Len(uint(k)) - 1; i >= 0; i-- {
		if k >= 1<<i { // k 在右半边
			ans++
			k -= 1 << i
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log k)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 优化

上面的代码相当于，每次遇到 $k-1$ 二进制中的 $1$，就把答案加一。

所以答案为 $\texttt{a}$ 加上 $k-1$ 二进制中的 $1$ 的个数。

注意在本题的数据范围下，无需和 $26$ 取模。

```py [sol-Python3]
class Solution:
    def kthCharacter(self, k: int) -> str:
        return ascii_lowercase[(k - 1).bit_count()]
```

```java [sol-Java]
class Solution {
    public char kthCharacter(int k) {
        return (char) ('a' + Integer.bitCount(k - 1));
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    char kthCharacter(long long k) {
        return 'a' + __builtin_popcount(k - 1);
    }
};
```

```go [sol-Go]
func kthCharacter(k int) byte {
	return byte('a' + bits.OnesCount(uint(k-1)))
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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
