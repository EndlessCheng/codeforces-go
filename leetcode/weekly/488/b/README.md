由于每次操作的都是最左边的相邻相等元素，我们可以从左到右遍历 $\textit{nums}$，一旦发现 $\textit{nums}[i]$ 和左侧剩余最近元素相等，就执行合并操作。如果合并后，新元素与左侧剩余最近元素相等，就继续执行合并操作。

如何快速找到左侧剩余最近元素、删除（合并）左侧剩余最近元素？

我们需要一个后进先出的数据结构——栈，模拟上述过程。

[本题视频讲解](https://www.bilibili.com/video/BV1idFoz3Efi/)，欢迎点赞关注~

## 写法一：先入栈，再出栈

```py [sol-Python3]
class Solution:
    def mergeAdjacent(self, nums: List[int]) -> List[int]:
        st = []
        for x in nums:
            st.append(x)
            while len(st) > 1 and st[-1] == st[-2]:
                st.pop()
                st[-1] *= 2
        return st
```

```java [sol-Java]
class Solution {
    public List<Long> mergeAdjacent(int[] nums) {
        List<Long> st = new ArrayList<>();
        for (int x : nums) {
            st.add((long) x);
            while (st.size() > 1 && st.getLast().equals(st.get(st.size() - 2))) {
                st.removeLast();
                int i = st.size() - 1;
                st.set(i, st.get(i) * 2);
            }
        }
        return st;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> mergeAdjacent(vector<int>& nums) {
        vector<long long> st;
        for (int x : nums) {
            st.push_back(x);
            while (st.size() > 1 && st.back() == st[st.size() - 2]) {
                st.pop_back();
                st.back() *= 2;
            }
        }
        return st;
    }
};
```

```go [sol-Go]
func mergeAdjacent(nums []int) []int64 {
	st := nums[:0] // 原地
	for _, x := range nums {
		st = append(st, x)
		for len(st) > 1 && st[len(st)-1] == st[len(st)-2] {
			st = st[:len(st)-1]
			st[len(st)-1] *= 2
		}
	}
	// 力扣的 int 就是 int64，直接 O(1) 转成 []int64
	return *(*[]int64)(unsafe.Pointer(&st))
}
```

## 写法二：先出栈，再入栈

```py [sol-Python3]
class Solution:
    def mergeAdjacent(self, nums: List[int]) -> List[int]:
        st = []
        for x in nums:
            while st and st[-1] == x:
                st.pop()
                x *= 2
            st.append(x)
        return st
```

```java [sol-Java]
class Solution {
    public List<Long> mergeAdjacent(int[] nums) {
        List<Long> st = new ArrayList<>();
        for (long x : nums) {
            while (!st.isEmpty() && st.getLast() == x) {
                st.removeLast();
                x *= 2;
            }
            st.add(x);
        }
        return st;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> mergeAdjacent(vector<int>& nums) {
        vector<long long> st;
        for (long long x : nums) {
            while (!st.empty() && st.back() == x) {
                st.pop_back();
                x *= 2;
            }
            st.push_back(x);
        }
        return st;
    }
};
```

```go [sol-Go]
func mergeAdjacent(nums []int) []int64 {
	st := nums[:0] // 原地
	for _, x := range nums {
		for len(st) > 0 && st[len(st)-1] == x {
			st = st[:len(st)-1]
			x *= 2
		}
		st = append(st, x)
	}
	// 力扣的 int 就是 int64，直接 O(1) 转成 []int64
	return *(*[]int64)(unsafe.Pointer(&st))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。虽然我们写了个二重循环，但站在每个元素的视角看，这个元素在二重循环中最多入栈出栈各一次，因此循环次数**之和**是 $\mathcal{O}(n)$，所以时间复杂度是 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。部分语言可以直接把 $\textit{nums}$ 当作栈，从而做到 $\mathcal{O}(1)$ 空间。

## 专题训练

见下面数据结构题单的「**§3.3 邻项消除**」。

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
