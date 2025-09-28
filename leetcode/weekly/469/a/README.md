不断地把 $n$ 除以 $10$（下取整）直到 $0$，例如 $537\to 53\to 5\to 0$。这个过程中的 $n\bmod 10$，即为每个数位。

我们需要对其中不为 $0$ 的数位，乘以 $10^k$。例如百位数是 $5$，那么就计算 $5\times 10^2 = 500$，加入答案。

由于我们是从低位到高位计算的，最后要把答案反转。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def decimalRepresentation(self, n: int) -> List[int]:
        ans = []
        pow10 = 1
        while n:
            n, d = divmod(n, 10)
            if d:
                ans.append(d * pow10)
            pow10 *= 10
        ans.reverse()
        return ans
```

```java [sol-Java]
class Solution {
    public int[] decimalRepresentation(int n) {
        List<Integer> a = new ArrayList<>();
        int pow10 = 1;
        for (; n > 0; n /= 10) {
            int d = n % 10;
            if (d > 0) {
                a.add(d * pow10);
            }
            pow10 *= 10;
        }
        Collections.reverse(a);

        int[] ans = new int[a.size()];
        for (int i = 0; i < a.size(); i++) {
            ans[i] = a.get(i);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> decimalRepresentation(int n) {
        vector<int> ans;
        long long pow10 = 1;
        for (; n > 0; n /= 10) {
            int d = n % 10;
            if (d > 0) {
                ans.push_back(d * pow10);
            }
            pow10 *= 10;
        }
        ranges::reverse(ans);
        return ans;
    }
};
```

```go [sol-Go]
func decimalRepresentation(n int) (ans []int) {
	pow10 := 1
	for ; n > 0; n /= 10 {
		d := n % 10
		if d > 0 {
			ans = append(ans, d*pow10)
		}
		pow10 *= 10
	}
	slices.Reverse(ans)
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

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
