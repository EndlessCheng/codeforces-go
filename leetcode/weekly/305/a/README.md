## 视频讲解

见[【周赛 305】](https://www.bilibili.com/video/BV1CN4y1V7uE)。

## 方法一：哈希表

> 先吐槽一下，arithmetic triplets 应该翻译成**等差三元组**更加贴切（例如 arithmetic progression 翻译成等差数列）。

由于 $\textit{nums}$ 是严格递增的，对于一个特定的 $\textit{nums}[j]$，如果它在等差三元组中，那么这样的等差三元组是唯一的，即 

$$
(\textit{nums}[j]-\textit{diff},\textit{nums}[j],\textit{nums}[j]+\textit{diff})
$$

我们可以用哈希表记录 $\textit{nums}$ 的每个元素，然后遍历 $\textit{nums}$，看 $\textit{nums}[j]-\textit{diff}$ 和 $\textit{nums}[j]+\textit{diff}$ 是否都在哈希表中。

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

### 前置知识：同向双指针

见[【基础算法精讲 01】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

### 思路

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

---

欢迎关注[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)，高质量算法教学，持续更新中~

附：[每日一题·高质量题解精选](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)。
