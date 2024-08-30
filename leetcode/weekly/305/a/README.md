## 方法一：哈希表

由于 $\textit{nums}$ 是严格递增的，对于一个特定的 $\textit{nums}[j]$，如果它在等差三元组中，那么这样的等差三元组是唯一的，即 

$$
(\textit{nums}[j]-\textit{diff},\textit{nums}[j],\textit{nums}[j]+\textit{diff})
$$

我们可以用哈希表记录 $\textit{nums}$ 的每个元素，然后遍历 $\textit{nums}$，看 $\textit{nums}[j]-\textit{diff}$ 和 $\textit{nums}[j]+\textit{diff}$ 是否都在哈希表中。

[视频讲解【周赛 305】](https://www.bilibili.com/video/BV1CN4y1V7uE)

```py [sol1-Python3]
class Solution:
    def arithmeticTriplets(self, nums: List[int], diff: int) -> int:
        s = set(nums)
        return sum(x - diff in s and x + diff in s for x in nums)
```

```java [sol1-Java]
class Solution {
    public int arithmeticTriplets(int[] nums, int diff) {
        int ans = 0;
        var set = new HashSet<Integer>();
        for (int x : nums) set.add(x);
        for (int x : nums)
            if (set.contains(x - diff) && set.contains(x + diff))
                ++ans;
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int arithmeticTriplets(vector<int> &nums, int diff) {
        int ans = 0;
        unordered_set<int> s{nums.begin(), nums.end()};
        for (int x: nums)
            if (s.count(x - diff) && s.count(x + diff))
                ++ans;
        return ans;
    }
};
```

```go [sol1-Go]
func arithmeticTriplets(nums []int, diff int) (ans int) {
	set := map[int]bool{}
	for _, x := range nums {
		set[x] = true
	}
	for _, x := range nums {
		if set[x-diff] && set[x+diff] {
			ans++
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

### 优化

等差三元组也可以用 $\textit{nums}[k]$ 表示：

$$
(\textit{nums}[k]-2\cdot\textit{diff},\textit{nums}[k]-\textit{diff},\textit{nums}[k])
$$

所以还可以一边查询哈希表中是否有 $\textit{nums}[k]-2\cdot\textit{diff}$ 和 $\textit{nums}[k]-\textit{diff}$，一边把 $\textit{nums}[k]$ 插入哈希表，从而做到一次遍历。

```py [sol12-Python3]
class Solution:
    def arithmeticTriplets(self, nums: List[int], diff: int) -> int:
        ans, s = 0, set()
        for x in nums:
            if x - diff in s and x - diff * 2 in s:
                ans += 1
            s.add(x)
        return ans
```

```java [sol12-Java]
class Solution {
    public int arithmeticTriplets(int[] nums, int diff) {
        int ans = 0;
        var set = new HashSet<Integer>();
        for (int x : nums) {
            if (set.contains(x - diff) && set.contains(x - diff * 2))
                ++ans;
            set.add(x);
        }
        return ans;
    }
}
```

```cpp [sol12-C++]
class Solution {
public:
    int arithmeticTriplets(vector<int> &nums, int diff) {
        int ans = 0;
        unordered_set<int> s;
        for (int x: nums) {
            if (s.count(x - diff) && s.count(x - diff * 2))
                ++ans;
            s.insert(x);
        }
        return ans;
    }
};
```

```go [sol12-Go]
func arithmeticTriplets(nums []int, diff int) (ans int) {
	set := map[int]bool{}
	for _, x := range nums {
		if set[x-diff] && set[x-diff*2] {
			ans++
		}
		set[x] = true
	}
	return
}
```

## 方法二：三指针

由于 $\textit{nums}$ 是严格递增的，遍历 $k$ 时，$i$ 和 $j$ 只增不减，因此可以用类似同向双指针的做法来移动指针：

1. 枚举 $x=\textit{nums}[k]$；
2. 移动 $j$ 直到 $\textit{nums}[j] + \textit{diff}\ge x$；
3. 如果 $\textit{nums}[j] + \textit{diff}= x$，则移动 $i$ 直到 $\textit{nums}[i] + 2\cdot\textit{diff}\ge x$；
4. 如果 $\textit{nums}[i] + 2\cdot\textit{diff}= x$，则找到了一对等差三元组。

注意下面代码在循环时没有判断 $j<k$ 和 $i<j$，因为一旦 $j=k$，$\textit{nums}[j] + \textit{diff}\ge x$ 必然成立，所以 $j<k$ 无需判断，$i$ 也同理。

```py [sol3-Python3]
class Solution:
    def arithmeticTriplets(self, nums: List[int], diff: int) -> int:
        ans, i, j = 0, 0, 1
        for x in nums:  # x = nums[k]
            while nums[j] + diff < x:
                j += 1
            if nums[j] + diff > x:
                continue
            while nums[i] + diff * 2 < x:
                i += 1
            if nums[i] + diff * 2 == x:
                ans += 1
        return ans
```

```java [sol3-Java]
class Solution {
    public int arithmeticTriplets(int[] nums, int diff) {
        int ans = 0, i = 0, j = 1;
        for (int x : nums) { // x = nums[k]
            while (nums[j] + diff < x)
                ++j;
            if (nums[j] + diff > x)
                continue;
            while (nums[i] + diff * 2 < x)
                ++i;
            if (nums[i] + diff * 2 == x)
                ++ans;
        }
        return ans;
    }
}
```

```cpp [sol3-C++]
class Solution {
public:
    int arithmeticTriplets(vector<int> &nums, int diff) {
        int ans = 0, i = 0, j = 1;
        for (int x: nums) { // x = nums[k]
            while (nums[j] + diff < x)
                ++j;
            if (nums[j] + diff > x)
                continue;
            while (nums[i] + diff * 2 < x)
                ++i;
            if (nums[i] + diff * 2 == x)
                ++ans;
        }
        return ans;
    }
};
```

```go [sol3-Go]
func arithmeticTriplets(nums []int, diff int) (ans int) {
	i, j := 0, 1
	for _, x := range nums[2:] { // x = nums[k]
		for nums[j]+diff < x {
			j++
		}
		if nums[j]+diff > x {
			continue
		}
		for nums[i]+diff*2 < x {
			i++
		}
		if nums[i]+diff*2 == x {
			ans++
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。虽然写了个二重循环，但是 `i++` 和 `j++` 的执行次数不会超过 $n$ 次，所以总的时间复杂度为 $O(n)$。
- 空间复杂度：$O(1)$，仅用到若干额外变量。

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
