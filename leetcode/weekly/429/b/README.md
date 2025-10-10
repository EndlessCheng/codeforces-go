一个现实中的例子：

- 军训的某一天，同学们在操场上。现在教官吹响了口哨，同学们集合，排成一排。对于最靠左的同学 $A$，他需要**尽量往左移动**，给其他同学腾出位置。$A$ 旁边的同学，可以**紧挨着** $A$。依此类推。

把 $\textit{nums}$ 视作 $n$ 个同学在一维数轴中的位置，从最左边的同学（$\textit{nums}$ 的最小值）开始思考。

设最左边的同学的位置为 $a$，他尽量往左移，位置变成 $a-k$。

$\textit{nums}$ 的次小值 $b$ 呢？这位同学也尽量往左移：

- 比如 $a=4,b=6,k=3$，那么 $a$ 变成 $a-k=1$，$b$ 变成 $b-k=3$。
- 比如 $a=4,b=4,k=3$，那么 $a$ 变成 $a'=a-k=1$，$b$ 变成 $b-k=1$ 就和 $a'$ 一样了，可以稍微大一点（紧挨着 $a'$），把 $b$ 变成 $a'+1=2$。

一般地，$b$ 变成

$$
\max(b-k,a'+1)
$$

但这不能超过 $b+k$，所以最终要变成

$$
\min(\max(b-k,a'+1),b+k)
$$

> 相当于让 $a'+1$ 落在 $[b-k,b+k]$ 中，若超出范围则修正。

第三小的数也同理，通过前一个数可以算出当前元素能变成多少。

最后答案为 $\textit{nums}$ 中的不同元素个数。我们可以在修改的同时统计，如果发现当前元素修改后的值，比上一个元素修改后的值大，那么答案加一。

为方便计算，把 $\textit{nums}$ 从小到大排序。排序后，从左到右遍历数组，就相当于从最左边的人开始计算了。

[本题视频讲解](https://www.bilibili.com/video/BV1wmkqYREnP/?t=4m20s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxDistinctElements(self, nums: List[int], k: int) -> int:
        nums.sort()
        ans = 0
        pre = -inf  # 记录每个人左边的人的位置
        for x in nums:
            x = min(max(x - k, pre + 1), x + k)
            if x > pre:
                ans += 1
                pre = x
        return ans
```

```java [sol-Java]
class Solution {
    public int maxDistinctElements(int[] nums, int k) {
        Arrays.sort(nums);
        int ans = 0;
        int pre = Integer.MIN_VALUE; // 记录每个人左边的人的位置
        for (int x : nums) {
            x = Math.min(Math.max(x - k, pre + 1), x + k);
            if (x > pre) {
                ans++;
                pre = x;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxDistinctElements(vector<int>& nums, int k) {
        ranges::sort(nums);
        int ans = 0;
        int pre = INT_MIN; // 记录每个人左边的人的位置
        for (int x : nums) {
            x = clamp(pre + 1, x - k, x + k); // min(max(x - k, pre + 1), x + k)
            if (x > pre) {
                ans++;
                pre = x;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxDistinctElements(nums []int, k int) (ans int) {
	slices.Sort(nums)
	pre := math.MinInt // 记录每个人左边的人的位置
	for _, x := range nums {
		x = min(max(x-k, pre+1), x+k)
		if x > pre {
			ans++
			pre = x
		}
	}
	return
}
```

## 优化

什么情况下，可以直接返回 $n$？

先考虑 $\textit{nums}$ 所有元素都相同的情况（同学们都挤在一起）。我们可以把元素 $x$ 变成 $[x-k,x+k]$ 中的整数，这一共有 $2k+1$ 个。如果 $2k+1 \ge n$，就可以让所有元素互不相同。

如果 $\textit{nums}$ 有不同元素，当 $2k+1 \ge n$ 时，更加可以让所有元素互不相同。

所以只要 $2k+1 \ge n$，就可以直接返回 $n$。

```py [sol-Python3]
class Solution:
    def maxDistinctElements(self, nums: List[int], k: int) -> int:
        if k * 2 + 1 >= len(nums):
            return len(nums)

        nums.sort()
        ans = 0
        pre = -inf  # 记录每个人左边的人的位置
        for x in nums:
            x = min(max(x - k, pre + 1), x + k)
            if x > pre:
                ans += 1
                pre = x
        return ans
```

```java [sol-Java]
class Solution {
    public int maxDistinctElements(int[] nums, int k) {
        int n = nums.length;
        if (k * 2 + 1 >= n) {
            return n;
        }

        Arrays.sort(nums);
        int ans = 0;
        int pre = Integer.MIN_VALUE; // 记录每个人左边的人的位置
        for (int x : nums) {
            x = Math.min(Math.max(x - k, pre + 1), x + k);
            if (x > pre) {
                ans++;
                pre = x;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxDistinctElements(vector<int>& nums, int k) {
        int n = nums.size();
        if (k * 2 + 1 >= n) {
            return n;
        }

        ranges::sort(nums);
        int ans = 0;
        int pre = INT_MIN; // 记录每个人左边的人的位置
        for (int x : nums) {
            x = clamp(pre + 1, x - k, x + k); // min(max(x - k, pre + 1), x + k)
            if (x > pre) {
                ans++;
                pre = x;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxDistinctElements(nums []int, k int) (ans int) {
	n := len(nums)
	if k*2+1 >= n {
		return n
	}

	slices.Sort(nums)
	pre := math.MinInt // 记录每个人左边的人的位置
	for _, x := range nums {
		x = min(max(x-k, pre+1), x+k)
		if x > pre {
			ans++
			pre = x
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 专题训练

见下面贪心题单的「**§1.1 从最小/最大开始贪心**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
