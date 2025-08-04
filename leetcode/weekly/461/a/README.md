## 写法一：三个循环

三段式子数组必须满足「严格递增 - 严格递减 - 严格递增」，一共三段，每一段**至少要有两个数**。

每一段分别用一个循环寻找，具体请看 [视频讲解](https://www.bilibili.com/video/BV1BEh3zZEoM/) 中的图，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def isTrionic(self, nums: List[int]) -> bool:
        n = len(nums)
        # 第一段
        i = 1
        while i < n and nums[i - 1] < nums[i]:
            i += 1
        if i == 1:  # 第一段至少要有两个数
            return False

        # 第二段
        i0 = i
        while i < n and nums[i - 1] > nums[i]:
            i += 1
        if i == i0 or i == n:  # 第二段至少要有两个数，第三段至少要有两个数
            return False

        # 第三段
        while i < n and nums[i - 1] < nums[i]:
            i += 1
        return i == n
```

```java [sol-Java]
class Solution {
    public boolean isTrionic(int[] nums) {
        int n = nums.length;
        // 第一段
        int i = 1;
        while (i < n && nums[i - 1] < nums[i]) {
            i++;
        }
        if (i == 1) { // 第一段至少要有两个数
            return false;
        }

        // 第二段
        int i0 = i;
        while (i < n && nums[i - 1] > nums[i]) {
            i++;
        }
        if (i == i0 || i == n) { // 第二段至少要有两个数，第三段至少要有两个数
            return false;
        }

        // 第三段
        while (i < n && nums[i - 1] < nums[i]) {
            i++;
        }
        return i == n;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isTrionic(vector<int>& nums) {
        int n = nums.size();
        // 第一段
        int i = 1;
        while (i < n && nums[i - 1] < nums[i]) {
            i++;
        }
        if (i == 1) { // 第一段至少要有两个数
            return false;
        }

        // 第二段
        int i0 = i;
        while (i < n && nums[i - 1] > nums[i]) {
            i++;
        }
        if (i == i0 || i == n) { // 第二段至少要有两个数，第三段至少要有两个数
            return false;
        }

        // 第三段
        while (i < n && nums[i - 1] < nums[i]) {
            i++;
        }
        return i == n;
    }
};
```

```go [sol-Go]
func isTrionic(nums []int) bool {
	n := len(nums)
	// 第一段
	i := 1
	for i < n && nums[i-1] < nums[i] {
		i++
	}
	if i == 1 { // 第一段至少要有两个数
		return false
	}

	// 第二段
	i0 := i
	for i < n && nums[i-1] > nums[i] {
		i++
	}
	if i == i0 || i == n { // 第二段至少要有两个数，第三段至少要有两个数
		return false
	}

	// 第三段
	for i < n && nums[i-1] < nums[i] {
		i++
	}
	return i == n
}
```

## 写法二：一个循环

```py [sol-Python3]
class Solution:
    def isTrionic(self, nums: List[int]) -> bool:
        if nums[0] >= nums[1]:  # 一开始必须是递增的
            return False
        cnt = 1
        for i in range(2, len(nums)):
            if nums[i - 1] == nums[i]:
                return False
            if (nums[i - 2] < nums[i - 1]) != (nums[i - 1] < nums[i]):
                cnt += 1
        return cnt == 3  # 一定是增减增
```

```java [sol-Java]
class Solution {
    public boolean isTrionic(int[] nums) {
        if (nums[0] >= nums[1]) { // 一开始必须是递增的
            return false;
        }
        int cnt = 1;
        for (int i = 2; i < nums.length; i++) {
            if (nums[i - 1] == nums[i]) {
                return false;
            }
            if ((nums[i - 2] < nums[i - 1]) != (nums[i - 1] < nums[i])) {
                cnt++;
            }
        }
        return cnt == 3; // 一定是增减增
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isTrionic(vector<int>& nums) {
        if (nums[0] >= nums[1]) { // 一开始必须是递增的
            return false;
        }
        int cnt = 1;
        for (int i = 2; i < nums.size(); i++) {
            if (nums[i - 1] == nums[i]) {
                return false;
            }
            if ((nums[i - 2] < nums[i - 1]) != (nums[i - 1] < nums[i])) {
                cnt++;
            }
        }
        return cnt == 3; // 一定是增减增
    }
};
```

```go [sol-Go]
func isTrionic(nums []int) bool {
	if nums[0] >= nums[1] { // 一开始必须是递增的
		return false
	}
	cnt := 1
	for i := 2; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return false
		}
		if (nums[i-2] < nums[i-1]) != (nums[i-1] < nums[i]) {
			cnt++
		}
	}
	return cnt == 3 // 一定是增减增
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面滑动窗口与双指针题单的「**六、分组循环**」。

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
