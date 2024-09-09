请看 [视频讲解](https://www.bilibili.com/video/BV1Yu4y1v7H6/) 第三题。

## 方法一：最长非递减子序列

转换成最多可以**保留**多少个元素不变。这些保留的元素必须是非递减的，请看 [最长递增子序列【基础算法精讲 20】](https://www.bilibili.com/video/BV1ub411Q7sB/)，视频末尾讲了如何处理非递减的情况。

```py [sol-Python3]
class Solution:
    def minimumOperations(self, nums: List[int]) -> int:
        g = []
        for x in nums:
            j = bisect_right(g, x)
            if j == len(g):
                g.append(x)
            else:
                g[j] = x
        return len(nums) - len(g)
```

```java [sol-Java]
class Solution {
    public int minimumOperations(List<Integer> nums) {
        List<Integer> g = new ArrayList<>();
        for (int x : nums) {
            int j = upperBound(g, x);
            if (j == g.size()) g.add(x);
            else g.set(j, x);
        }
        return nums.size() - g.size();
    }

    // 开区间写法
    private int upperBound(List<Integer> g, int target) {
        int left = -1, right = g.size(); // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] <= target
            // nums[right] > target
            int mid = (left + right) >>> 1;
            if (g.get(mid) <= target)
                left = mid; // 范围缩小到 (mid, right)
            else
                right = mid; // 范围缩小到 (left, mid)
        }
        return right; // 或者 left+1
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumOperations(vector<int>& nums) {
        vector<int> g;
        for (int x : nums) {
            auto it = ranges::upper_bound(g, x);
            if (it == g.end()) g.push_back(x);
            else *it = x;
        }
        return nums.size() - g.size();
    }
};
```

```go [sol-Go]
func minimumOperations(nums []int) int {
	g := []int{}
	for _, x := range nums {
		p := sort.SearchInts(g, x+1)
		if p < len(g) {
			g[p] = x
		} else {
			g = append(g, x)
		}
	}
	return len(nums) - len(g)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：状态机 DP

请看[【基础算法精讲 21】](https://www.bilibili.com/video/BV1ho4y1W7QK/)。

定义 $f[i+1][j]$ 表示把 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 这段变成非递减，且 $\textit{nums}[i]=j$ 的最小修改次数。

枚举第 $i-1$ 个数改成了 $k$，有

$$
f[i+1][j] = \min_{1\le k\le j} f[i][k] + [j \ne \textit{nums}[i]]
$$

初始值 $f[0][j] = 0$。

答案为 $\min(f[n])$。

代码实现时，第一个维度可以省略。为了避免状态被覆盖，需要倒序枚举 $j$。

```py [sol-Python3]
class Solution:
    def minimumOperations(self, nums: List[int]) -> int:
        f = [0] * 4
        for x in nums:
            for j in range(3, 0, -1):
                f[j] = min(f[k] for k in range(1, j + 1)) + (j != x)
        return min(f[1:])
```

```java [sol-Java]
class Solution {
    public int minimumOperations(List<Integer> nums) {
        var f = new int[4];
        for (int x : nums) {
            for (int j = 3; j > 0; j--) {
                for (int k = 1; k <= j; k++)
                    f[j] = Math.min(f[j], f[k]);
                if (j != x) f[j]++;
            }
        }
        int ans = nums.size();
        for (int j = 1; j < 4; j++)
            ans = Math.min(ans, f[j]);
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumOperations(vector<int>& nums) {
        int f[4]{};
        for (int x : nums)
            for (int j = 3; j; j--)
                f[j] = *min_element(f + 1, f + j + 1) + (j != x);
        return *min_element(f + 1, f + 4);
    }
};
```

```go [sol-Go]
func minimumOperations(nums []int) int {
	f := [4]int{}
	for _, x := range nums {
		for j := 3; j > 0; j-- {
			for k := 1; k <= j; k++ {
				f[j] = min(f[j], f[k])
			}
			if j != x {
				f[j]++
			}
		}
	}
	return min(len(nums), f[1], f[2], f[3])
}
```

也可以计算至多保留多少个元素：

$$
f[j] = \max(f[j], f[j-1]) + [j = \textit{nums}[i]]
$$

```py [sol-Python3]
class Solution:
    def minimumOperations(self, nums: List[int]) -> int:
        f = [0] * 4
        for x in nums:
            f[x] += 1
            f[2] = max(f[2], f[1])
            f[3] = max(f[3], f[2])
        return len(nums) - max(f)
```

```java [sol-Java]
class Solution {
    public int minimumOperations(List<Integer> nums) {
        var f = new int[4];
        for (int x : nums) {
            f[x]++;
            f[2] = Math.max(f[2], f[1]);
            f[3] = Math.max(f[3], f[2]);
        }
        return nums.size() - Math.max(Math.max(f[1], f[2]), f[3]);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumOperations(vector<int>& nums) {
        int f[4]{};
        for (int x: nums) {
            f[x]++;
            f[2] = max(f[2], f[1]);
            f[3] = max(f[3], f[2]);
        }
        return nums.size() - *max_element(f + 1, f + 4);
    }
};
```

```go [sol-Go]
func minimumOperations(nums []int) int {
	f := [4]int{}
	for _, x := range nums {
		f[x]++
		f[2] = max(f[2], f[1])
		f[3] = max(f[3], f[2])
	}
	return len(nums) - max(f[1], f[2], f[3])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

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
