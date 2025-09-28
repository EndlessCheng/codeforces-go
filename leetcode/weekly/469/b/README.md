根据题意，$\textit{nums}$ 必须**先严格递增，再严格递减**（山峰数组）。

有三种情况：

- $[1,2,3,1,3,2]$ 或者 $[1,2,3,3,3,2]$ 等。前者不是先严格递增，再严格递减；后者峰顶的重复元素太多，无法分割。
- $[1,2,3,3,2]$，峰顶恰好有两个相等元素，有唯一分割方案 $[1,2,3] + [3,2]$。
- $[1,2,3,2,1]$，峰顶没有重复元素，有两种分割方案：
  - 峰顶给前缀，$[1,2,3] + [2,1]$。
  - 峰顶给后缀，$[1,2] + [3,2,1]$。

从左到右，找最长严格递增前缀 $[0,i-1]$，最长严格递减后缀 $[j+1,n-1]$。

分类讨论：

- 如果 $i\le j$，对应第一种情况，返回 $-1$。
- 如果 $i-1=j$，对应第二种情况，返回 $[0,i-1]$ 的元素和与 $[i,n-1]$ 的元素和的绝对差。
- 否则 $i-2=j$，对应第三种情况，返回如下两个数的最小值：
  - $[0,i-1]$ 的元素和与 $[i,n-1]$ 的元素和的绝对差。
  - $[0,i-2]$ 的元素和与 $[i-1,n-1]$ 的元素和的绝对差。

[本题视频讲解](https://www.bilibili.com/video/BV156n9z7E9o/?t=5m38s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def splitArray(self, nums: List[int]) -> int:
        n = len(nums)
        # 最长严格递增前缀
        pre = nums[0]
        i = 1
        while i < n and nums[i] > nums[i - 1]:
            pre += nums[i]
            i += 1

        # 最长严格递减后缀
        suf = nums[-1]
        j = n - 2
        while j >= 0 and nums[j] > nums[j + 1]:
            suf += nums[j]
            j -= 1

        # 情况一
        if i <= j:
            return -1

        d = pre - suf
        # 情况二
        if i - 1 == j:
            return abs(d)

        # 情况三，suf 多算了一个 nums[i-1]，或者 pre 多算了一个 nums[i-1]
        return min(abs(d + nums[i - 1]), abs(d - nums[i - 1]))
```

```java [sol-Java]
class Solution {
    public long splitArray(int[] nums) {
        int n = nums.length;
        // 最长严格递增前缀
        long pre = nums[0];
        int i = 1;
        while (i < n && nums[i] > nums[i - 1]) {
            pre += nums[i];
            i++;
        }

        // 最长严格递减后缀
        long suf = nums[n - 1];
        int j = n - 2;
        while (j >= 0 && nums[j] > nums[j + 1]) {
            suf += nums[j];
            j--;
        }

        // 情况一
        if (i <= j) {
            return -1;
        }

        long d = pre - suf;
        // 情况二
        if (i - 1 == j) {
            return Math.abs(d);
        }

        // 情况三，suf 多算了一个 nums[i-1]，或者 pre 多算了一个 nums[i-1]
        return Math.min(Math.abs(d + nums[i - 1]), Math.abs(d - nums[i - 1]));
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long splitArray(vector<int>& nums) {
        int n = nums.size();
        // 最长严格递增前缀
        long long pre = nums[0];
        int i = 1;
        while (i < n && nums[i] > nums[i - 1]) {
            pre += nums[i];
            i++;
        }

        // 最长严格递减后缀
        long long suf = nums[n - 1];
        int j = n - 2;
        while (j >= 0 && nums[j] > nums[j + 1]) {
            suf += nums[j];
            j--;
        }

        // 情况一
        if (i <= j) {
            return -1;
        }

        long long d = pre - suf;
        // 情况二
        if (i - 1 == j) {
            return abs(d);
        }

        // 情况三，suf 多算了一个 nums[i-1]，或者 pre 多算了一个 nums[i-1]
        return min(abs(d + nums[i - 1]), abs(d - nums[i - 1]));
    }
};
```

```go [sol-Go]
func splitArray(nums []int) int64 {
	n := len(nums)
	// 最长严格递增前缀
	pre := nums[0]
	i := 1
	for i < n && nums[i] > nums[i-1] {
		pre += nums[i]
		i++
	}

	// 最长严格递减后缀
	suf := nums[n-1]
	j := n - 2
	for j >= 0 && nums[j] > nums[j+1] {
		suf += nums[j]
		j--
	}

	// 情况一
	if i <= j {
		return -1
	}

	d := pre - suf
	// 情况二
	if i-1 == j {
		return int64(abs(d))
	}

	// 情况三，suf 多算了一个 nums[i-1]，或者 pre 多算了一个 nums[i-1]
	return int64(min(abs(d+nums[i-1]), abs(d-nums[i-1])))
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
