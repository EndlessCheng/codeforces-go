统计每个数组的出现次数，记在一个 $\textit{cnt}$ 哈希表（或者数组）中。

然后遍历 $\textit{cnt}$，找到其中出现次数最小的数字。如果出现次数相同，取其中最小的数字。

如何把遍历一个数字的每一位？

一种方法是转成字符串，然后遍历字符串。

另一种方法是不断地把 $n$ 除以 $10$（下取整）直到 $0$，例如 $123\to 12\to 1\to 0$。在这个过程中的 $n\bmod 10$，即为每个数位。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1aCaGzWEm4/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def getLeastFrequentDigit(self, n: int) -> int:
        # 统计每个数字的出现次数
        cnt = Counter(str(n))

        # 找出现次数最小的数字
        min_ch = min((c, ch) for ch, c in cnt.items())[1]
        return int(min_ch)
```

```py [sol-Python3 写法二]
class Solution:
    def getLeastFrequentDigit(self, n: int) -> int:
        # 统计每个数字的出现次数
        cnt = [0] * 10
        while n > 0:
            cnt[n % 10] += 1
            n //= 10

        # 找出现次数最小的数字
        min_c = inf
        ans = 0
        for i, c in enumerate(cnt):
            if 0 < c < min_c:
                min_c = c
                ans = i
        return ans
```

```java [sol-Java]
class Solution {
    public int getLeastFrequentDigit(int n) {
        // 统计每个数字的出现次数
        int[] cnt = new int[10];
        while (n > 0) {
            cnt[n % 10]++;
            n /= 10;
        }

        // 找出现次数最小的数字
        int minC = Integer.MAX_VALUE;
        int ans = 0;
        for (int i = 0; i < 10; i++) {
            int c = cnt[i];
            if (c > 0 && c < minC) {
                minC = c;
                ans = i;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int getLeastFrequentDigit(int n) {
        // 统计每个数字的出现次数
        int cnt[10]{};
        while (n > 0) {
            cnt[n % 10]++;
            n /= 10;
        }

        // 找出现次数最小的数字
        int min_c = INT_MAX;
        int ans = 0;
        for (int i = 0; i < 10; i++) {
            int c = cnt[i];
            if (c > 0 && c < min_c) {
                min_c = c;
                ans = i;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func getLeastFrequentDigit(n int) (ans int) {
	// 统计每个数字的出现次数
	cnt := [10]int{}
	for n > 0 {
		cnt[n%10]++
		n /= 10
	}

	// 找出现次数最小的数字
	minC := math.MaxInt
	for i, c := range cnt {
		if c > 0 && c < minC {
			minC = c
			ans = i
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$ 或 $\mathcal{O}(D + \log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$D=10$ 是 $\textit{cnt}$ 数组的大小。
- 空间复杂度：$\mathcal{O}(\log n)$ 或 $\mathcal{O}(D)$。

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
