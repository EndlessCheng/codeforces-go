### 本题视频讲解

见[【双周赛 100】](https://www.bilibili.com/video/BV1WM411H7UE/)。

# 方法一：贪心 + 排序 + 双指针

### 提示 1

田忌赛马。

### 提示 2-1

想一想，$\textit{nums}$ 的最小值是否要参与贡献伟大值？要和谁匹配？

### 提示 2-2

$\textit{nums}$ 的最小值要参与匹配，否则更大的数字更难匹配上。

$\textit{nums}$ 的最小值要与次小值匹配，这样后面的数字才能取匹配更大的数。

### 提示 3

为了方便实现，对 $\textit{nums}$ 从小到大排序。（为什么可以排序？因为只在乎匹配关系，与下标无关。）

例如示例 1 排序后为 $[1,1,1,2,3,3,5]$。那么前三个 $1$ 分别与 $2,3,3$ 匹配，$2$ 与 $5$ 匹配，后面就没有数字能匹配了。

```py [sol1-Python3]
class Solution:
    def maximizeGreatness(self, nums: List[int]) -> int:
        nums.sort()
        i = 0
        for x in nums:
            if x > nums[i]:
                i += 1
        return i
```

```java [sol1-Java]
class Solution {
    public int maximizeGreatness(int[] nums) {
        Arrays.sort(nums);
        int i = 0;
        for (int x : nums)
            if (x > nums[i])
                ++i;
        return i;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int maximizeGreatness(vector<int> &nums) {
        sort(nums.begin(), nums.end());
        int i = 0;
        for (int x : nums)
            if (x > nums[i])
                ++i;
        return i;
    }
};
```

```go [sol1-Go]
func maximizeGreatness(nums []int) (i int) {
	sort.Ints(nums)
	for _, x := range nums {
		if x > nums[i] {
			i++
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。忽略排序时的栈开销，仅用到若干额外变量。

# 方法二：利用两个指针的距离

考虑无法匹配的个数 $m$，答案为 $\textit{nums}$ 的长度减去 $m$。

再来看方法一的双指针，设另一个指针为 $j$，即 $x=\textit{nums}[j]$。

每次 `x > nums[i]` **不成立**时，两个指针错开的距离 $j-i$ 就增加 $1$。

那么循环结束后，两个指针最终错开的距离 $j-i$，就是 $m$。

以 $[1,1,2,2,2,2,3,3]$ 为例，初始 $i=0,j=0$：

- $j=0$，无法匹配，$j$ 加一。
- $j=1$，无法匹配，$j$ 加一。
- $j=2$，可以匹配，$i,j$ 都加一。
- $j=3$，可以匹配，$i,j$ 都加一。
- $j=4$，无法匹配，$j$ 加一。注意此时 $i$ 指向第一个 $2$。
- $j=5$，无法匹配，$j$ 加一。
- $j=6$，可以匹配，$i,j$ 都加一。注意此时 $j$ 指向最后一个 $2$ 的右侧相邻元素。
- $j=7$，可以匹配，$i,j$ 都加一。循环结束。

可以发现，当 $i$ 指向出现次数最多的数 $p$ 的时候，$j$ 要一直移动到下一个不等于 $p$ 的数，此时错开的距离是最大的。

由于后面的数出现次数不会超过 $p$，所以不会出现无法匹配的情况。（可以用反证法证明，如果出现了，说明这个数的出现次数大于 $p$，矛盾。）

所以 $m$ 就是 $\textit{nums}$ 中元素出现次数的最大值。用哈希表统计即可，无需排序。

```py [sol2-Python3]
class Solution:
    def maximizeGreatness(self, nums: List[int]) -> int:
        return len(nums) - max(Counter(nums).values())
```

```java [sol2-Java]
class Solution {
    public int maximizeGreatness(int[] nums) {
        int mx = 0;
        var cnt = new HashMap<Integer, Integer>();
        for (int x : nums)
            mx = Math.max(mx, cnt.merge(x, 1, Integer::sum));
        return nums.length - mx;
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    int maximizeGreatness(vector<int> &nums) {
        int mx = 0;
        unordered_map<int, int> cnt;
        for (int x : nums)
            mx = max(mx, ++cnt[x]);
        return nums.size() - mx;
    }
};
```

```go [sol2-Go]
func maximizeGreatness(nums []int) int {
	maxCnt := 0
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
		maxCnt = max(maxCnt, cnt[v])
	}
	return len(nums) - maxCnt
}

func max(a, b int) int { if a < b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。
