把偶数视作空位，奇数视作汽车。

例如 $[0,0,0,1,1,1]$，目标是空位 $0$ 和汽车 $1$ 交替，可以排列成 $[0,1,0,1,0,1]$，也可以排列成 $[1,0,1,0,1,0]$。两种情况的交换次数取最小值。

怎么交换？只需关注汽车，当所有汽车都移动到偶数（奇数）下标，所有空位就一定都位于奇数（偶数）下标。

比如汽车的位置为 $a=[3,4,5]$，目标位置为 $b=[1,3,5]$。贪心地，最左边的汽车移动到最左边的目标位置，比移动到其他位置更优。如果最左边的汽车没有移动到最左边的目标位置，那么其他更靠右的汽车就要移动到最左边，这样移动的总距离是更大的。

计算每辆车 $a[i]$ 到其目标位置 $b[i]$ 的距离 $|a[i]-b[i]|$，累加即为交换次数。

设车的个数为 $m$，那么空位的个数为 $n-m$，如果 $|(n-m)-m| > 1$，则无法形成有效排列，返回 $-1$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def minSwaps(self, nums: List[int]) -> int:
        pos1 = [i for i, x in enumerate(nums) if x % 2]

        n = len(nums)
        m = len(pos1)
        if abs(n - 2 * m) > 1:
            return -1

        # start=0 表示车要去偶数下标，start=1 表示车要去奇数下标
        def calc(start: int) -> int:
            # (n-start+1)//2 表示偶数（奇数）下标的个数
            if (n - start + 1) // 2 != m:
                return inf
            return sum(abs(i - j) for i, j in zip(range(start, n, 2), pos1))

        return min(calc(0), calc(1))
```

```java [sol-Java]
class Solution {
    public int minSwaps(int[] nums) {
        int n = nums.length;
        List<Integer> pos1 = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            if (nums[i] % 2 != 0) {
                pos1.add(i);
            }
        }

        int m = pos1.size();
        if (Math.abs(n - 2 * m) > 1) {
            return -1;
        }

        return Math.min(calc(0, pos1, n), calc(1, pos1, n));
    }

    // start=0 表示车要去偶数下标，start=1 表示车要去奇数下标
    private int calc(int start, List<Integer> pos1, int n) {
        // (n-start+1)/2 表示偶数（奇数）下标的个数
        if ((n - start + 1) / 2 != pos1.size()) {
            return Integer.MAX_VALUE;
        }
        int res = 0;
        for (int i = 0; i < pos1.size(); i++) {
            res += Math.abs(i * 2 + start - pos1.get(i));
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minSwaps(vector<int>& nums) {
        vector<int> pos1;
        int n = nums.size();
        for (int i = 0; i < n; i++) {
            if (nums[i] % 2) {
                pos1.push_back(i);
            }
        }

        int m = pos1.size();
        if (abs(n - 2 * m) > 1) {
            return -1;
        }

        // start=0 表示车要去偶数下标，start=1 表示车要去奇数下标
        auto calc = [&](int start) {
            // (n-start+1)/2 表示偶数（奇数）下标的个数
            if ((n - start + 1) / 2 != m) {
                return INT_MAX;
            }
            int res = 0;
            for (int i = 0; i < m; i++) {
                res += abs(i * 2 + start - pos1[i]);
            }
            return res;
        };

        return min(calc(0), calc(1));
    }
};
```

```go [sol-Go]
func minSwaps(nums []int) int {
	pos1 := []int{}
	for i, x := range nums {
		if x%2 != 0 {
			pos1 = append(pos1, i)
		}
	}

	n := len(nums)
	m := len(pos1)
	if abs(n-m*2) > 1 {
		return -1
	}

	// start=0 表示车要去偶数下标，start=1 表示车要去奇数下标
	calc := func(start int) (res int) {
		if (n-start+1)/2 != m { // (n-start+1)/2 表示偶数（奇数）下标的个数
			return math.MaxInt
		}
		for i, j := range pos1 {
			res += abs(i*2 + start - j)
		}
		return
	}
	return min(calc(0), calc(1))
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。**注**：如果一边遍历 $\textit{nums}$ 一边计算答案，则空间复杂度为 $\mathcal{O}(1)$。

## 相似题目

[1864. 构成交替字符串需要的最小交换次数](https://leetcode.cn/problems/minimum-number-of-swaps-to-make-the-binary-string-alternating/)

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
