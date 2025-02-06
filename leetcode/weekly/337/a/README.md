## 方法一：遍历二进制数

把 $n$ 当成一个二进制数来遍历。

遍历的顺序是从低位到高位。具体来说，通过 `n & 1` 取二进制的最低位，然后把 $n$ 右移一位，继续计算 `n & 1`，这样可以取到次低位。如此循环，直到 $n=0$ 为止。

在遍历的过程中，统计奇偶下标比特位中的 $1$ 的个数。

```py [sol-Python3]
class Solution:
    def evenOddBit(self, n: int) -> List[int]:
        ans = [0, 0]
        i = 0
        while n:
            ans[i] += n & 1
            n >>= 1
            i ^= 1  # 切换奇偶
        return ans
```

```java [sol-Java]
class Solution {
    public int[] evenOddBit(int n) {
        int[] ans = new int[2];
        for (int i = 0; n > 0; n >>= 1) {
            ans[i] += n & 1;
            i ^= 1; // 切换奇偶
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> evenOddBit(int n) {
        vector<int> ans(2);
        for (int i = 0; n; n >>= 1) {
            ans[i] += n & 1;
            i ^= 1; // 切换奇偶
        }
        return ans;
    }
};
```

```go [sol-Go]
func evenOddBit(n int) []int {
    ans := make([]int, 2)
    for i := 0; n > 0; n >>= 1 {
        ans[i] += n & 1
        i ^= 1 // 切换奇偶
    }
    return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：位掩码 + 库函数

利用位掩码 `0x55555555`（二进制的 $01010101\cdots$），与 $n$ 计算 AND，即可取出所有偶数下标比特，然后用库函数统计二进制中的 $1$ 的个数。

把 `0x55555555` 右移一位，与 $n$ 计算 AND，即可取出所有奇数下标比特，然后用库函数统计二进制中的 $1$ 的个数。

> 注：因为 $n$ 比较小，你也可以用 `0x555` 作为位掩码。

```py [sol-Python3]
class Solution:
    def evenOddBit(self, n: int) -> List[int]:
        MASK = 0x55555555
        return [(n & MASK).bit_count(), (n & (MASK >> 1)).bit_count()]
```

```java [sol-Java]
class Solution {
    public int[] evenOddBit(int n) {
        final int MASK = 0x55555555;
        return new int[]{Integer.bitCount(n & MASK), Integer.bitCount(n & (MASK >> 1))};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> evenOddBit(int n) {
        const unsigned MASK = 0x55555555;
        return {popcount(n & MASK), popcount(n & (MASK >> 1))};
    }
};
```

```go [sol-Go]
func evenOddBit(n int) []int {
    const mask = 0x55555555
    return []int{bits.OnesCount(uint(n & mask)), bits.OnesCount(uint(n & (mask >> 1)))}
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
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
