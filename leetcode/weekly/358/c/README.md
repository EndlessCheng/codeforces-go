请看 [视频讲解](https://www.bilibili.com/video/BV1wh4y1Q7XW/) 第三题。

```py [sol-Python3]
from sortedcontainers import SortedList

class Solution:
    def minAbsoluteDifference(self, nums: List[int], x: int) -> int:
        ans = inf
        sl = SortedList((-inf, inf))  # 哨兵
        for v, y in zip(nums, nums[x:]):
            sl.add(v)
            j = sl.bisect_left(y)
            ans = min(ans, sl[j] - y, y - sl[j - 1])
        return ans
```

```java [sol-Java]
class Solution {
    public int minAbsoluteDifference(List<Integer> nums, int x) {
        var a = nums.stream().mapToInt(i -> i).toArray();
        int ans = Integer.MAX_VALUE, n = a.length;
        var s = new TreeSet<Integer>();
        s.add(Integer.MAX_VALUE); // 哨兵
        s.add(Integer.MIN_VALUE / 2);
        for (int i = x; i < n; i++) {
            s.add(a[i - x]);
            int y = a[i];
            ans = Math.min(ans, Math.min(s.ceiling(y) - y, y - s.floor(y)));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minAbsoluteDifference(vector<int> &nums, int x) {
        int ans = INT_MAX, n = nums.size();
        set<int> s = {INT_MIN / 2, INT_MAX}; // 哨兵
        for (int i = x; i < n; i++) {
            s.insert(nums[i - x]);
            int y = nums[i];
            auto it = s.lower_bound(y); // 注意用 set 自带的 lower_bound，具体见视频中的解析
            ans = min(ans, min(*it - y, y - *prev(it))); // 注意不能写 *--it，这是未定义行为：万一先执行了 --it，前面的 *it-y 就错了
        }
        return ans;
    }
};
```

```go [sol-Go]
// import "github.com/emirpasic/gods/trees/redblacktree"
func minAbsoluteDifference(nums []int, k int) int {
	ans := math.MaxInt
	t := redblacktree.NewWithIntComparator()
	t.Put(math.MaxInt, nil) // 哨兵
	t.Put(math.MinInt/2, nil)
	for i, y := range nums[k:] {
		t.Put(nums[i], nil)
		c, _ := t.Ceiling(y)
		f, _ := t.Floor(y)
		ans = min(ans, min(c.Key.(int)-y, y-f.Key.(int)))
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n-x)\log (n-x))$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n-x)$。
