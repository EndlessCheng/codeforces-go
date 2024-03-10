题目说，同一个包裹中的苹果可以分装到不同的箱子中。

那么按照容量从大到小选择箱子装苹果，直到所有苹果均装入箱子为止。

注意题目保证可以将包裹中的苹果重新分装到箱子中。

```py [sol-Python3]
class Solution:
    def minimumBoxes(self, apple: List[int], capacity: List[int]) -> int:
        s = sum(apple)
        capacity.sort(reverse=True)
        for i, x in enumerate(capacity, 1):
            s -= x
            if s <= 0:  # 所有苹果都装入了箱子
                return i
```

```java [sol-Java]
class Solution {
    public int minimumBoxes(int[] apple, int[] capacity) {
        int s = 0;
        for (int x : apple) {
            s += x;
        }
        Arrays.sort(capacity);
        int m = capacity.length;
        int i = m - 1;
        for (; s > 0; i--) {
            s -= capacity[i];
        }
        return m - 1 - i;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumBoxes(vector<int> &apple, vector<int> &capacity) {
        int s = accumulate(apple.begin(), apple.end(), 0);
        ranges::sort(capacity, greater<>());
        int i = 0;
        for (; s > 0; i++) {
            s -= capacity[i];
        }
        return i;
    }
};
```

```go [sol-Go]
func minimumBoxes(apple, capacity []int) int {
	s := 0
	for _, x := range apple {
		s += x
	}
	slices.SortFunc(capacity, func(a, b int) int { return b - a })
	for i, c := range capacity {
		s -= c
		if s <= 0 { // 所有苹果都装入了箱子
			return i + 1 // 0 到 i 有 i+1 个箱子
		}
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m\log m)$，其中 $n$ 为 $\textit{apple}$ 的长度，$m$ 为 $\textit{capacity}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
