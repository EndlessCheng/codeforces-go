按照题目要求模拟即可，详见代码注释。

[本题视频讲解](https://www.bilibili.com/video/BV1W2zQBnE3g/)，欢迎点赞关注~

## 写法一

```py [sol-Python3]
class Solution:
    def rotateElements(self, nums: List[int], k: int) -> List[int]:
        # 取出非负数
        a = [x for x in nums if x >= 0]

        # 没有非负数，无需操作
        if not a:
            return nums

        # 向左轮替 k 个位置
        k %= len(a)
        a = a[k:] + a[:k]

        # 双指针，把 a 填入 nums，跳过负数
        j = 0
        for i, x in enumerate(nums):
            if x >= 0:
                nums[i] = a[j]
                j += 1
        return nums
```

```java [sol-Java]
class Solution {
    public int[] rotateElements(int[] nums, int k) {
        // 取出非负数
        List<Integer> a = new ArrayList<>();
        for (int x : nums) {
            if (x >= 0) {
                a.add(x);
            }
        }

        // 向左轮替 k 个位置（原地操作）
        Collections.rotate(a, -k);

        // 双指针，把 a 填入 nums，跳过负数
        int j = 0;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] >= 0) {
                nums[i] = a.get(j++);
            }
        }
        return nums;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> rotateElements(vector<int>& nums, int k) {
        // 取出非负数
        vector<int> a;
        for (int x : nums) {
            if (x >= 0) {
                a.push_back(x);
            }
        }

        // 没有非负数，无需操作
        if (a.empty()) {
            return nums;
        }

        // 向左轮替 k 个位置（原地操作）
        ranges::rotate(a, a.begin() + k % a.size());

        // 双指针，把 a 填入 nums，跳过负数
        int j = 0;
        for (int& x : nums) {
            if (x >= 0) {
                x = a[j++];
            }
        }
        return nums;
    }
};
```

```go [sol-Go]
// 原理类似 189. 轮转数组
// https://leetcode.cn/problems/rotate-array/
func rotateLeft(a []int, k int) {
	slices.Reverse(a[:k])
	slices.Reverse(a[k:])
	slices.Reverse(a)
}

func rotateElements(nums []int, k int) []int {
	// 取出非负数
	a := []int{}
	for _, x := range nums {
		if x >= 0 {
			a = append(a, x)
		}
	}

	m := len(a)
	// 没有非负数，无需操作
	if m == 0 {
		return nums
	}

	// 向左轮替 k 个位置（原地操作）
	rotateLeft(a, k%m)

	// 双指针，把 a 填入 nums，跳过负数
	j := 0
	for i, x := range nums {
		if x >= 0 {
			nums[i] = a[j]
			j++
		}
	}
	return nums
}
```

## 写法二

无需轮替，直接把双指针的 $j$ 初始化成 $k$，即 $a$ 轮替后的第一个数。

```py [sol-Python3]
class Solution:
    def rotateElements(self, nums: List[int], k: int) -> List[int]:
        # 取出非负数
        a = [x for x in nums if x >= 0]
        m = len(a)

        # 双指针，把 a 填入 nums，跳过负数
        j = k
        for i, x in enumerate(nums):
            if x >= 0:
                nums[i] = a[j % m]
                j += 1
        return nums
```

```java [sol-Java]
class Solution {
    public int[] rotateElements(int[] nums, int k) {
        // 取出非负数
        List<Integer> a = new ArrayList<>();
        for (int x : nums) {
            if (x >= 0) {
                a.add(x);
            }
        }

        // 双指针，把 a 填入 nums，跳过负数
        int j = k;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] >= 0) {
                nums[i] = a.get(j++ % a.size());
            }
        }
        return nums;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> rotateElements(vector<int>& nums, int k) {
        // 取出非负数
        vector<int> a;
        for (int x : nums) {
            if (x >= 0) {
                a.push_back(x);
            }
        }

        // 双指针，把 a 填入 nums，跳过负数
        int j = k;
        for (int& x : nums) {
            if (x >= 0) {
                x = a[j++ % a.size()];
            }
        }
        return nums;
    }
};
```

```go [sol-Go]
func rotateElements(nums []int, k int) []int {
	// 取出非负数
	a := []int{}
	for _, x := range nums {
		if x >= 0 {
			a = append(a, x)
		}
	}

	// 双指针，把 a 填入 nums，跳过负数
	j := k
	for i, x := range nums {
		if x >= 0 {
			nums[i] = a[j%len(a)]
			j++
		}
	}
	return nums
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(m)$，其中 $m$ 是 $\textit{nums}$ 中的非负数的个数。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
