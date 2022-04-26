### 方法一：差分

#### 提示

把 $\textit{flowers}_i$ 看成是将区间 $[\textit{start}_i,\textit{end}_i]$ 上的每个时间点都增加一朵花。

那么对于第 $i$ 个人，我们就需要计算出 $\textit{person}_i$ 时间点上有多少朵花。

#### 算法

用变化量表示一段区间上的更新，即在时间点 $\textit{start}_i$ 变化量增加了 $1$，在时间点 $\textit{end}_i+1$ 变化量减少了 $1$（类比导数的概念）。

遍历 $\textit{flowers}$，统计这些区间端点产生的变化量，记录在有序集合 $\textit{diff}$ 中。

然后从小到大遍历 $\textit{diff}$，累加变化量（类比积分的概念）。第 $i$ 个人到达时，花的数目即为不超过 $\textit{person}_i$ 时间点的变化量的累加值。

为了快速计算每个人的答案，我们需要将 $\textit{person}$ 从小到大排序，这样可以在遍历 $\textit{person}$ 的同时从小到大遍历 $\textit{diff}$。

- 时间复杂度：$O(n\log n + m\log m)$，其中 $n$ 是 $\textit{flowers}$ 的长度，$m$ 是 $\textit{persons}$ 的长度。
- 空间复杂度：$O(n+m)$。

```Python [sol1-Python3]
class Solution:
    def fullBloomFlowers(self, flowers: List[List[int]], persons: List[int]) -> List[int]:
        diff = defaultdict(int)  # 也可以用 SortedDict
        for start, end in flowers:
            diff[start] += 1
            diff[end + 1] -= 1
        times = sorted(diff.keys())

        n = len(persons)
        ans = [0] * n
        i = sum = 0
        for p, id in sorted(zip(persons, range(n))):
            while i < len(times) and times[i] <= p:
                sum += diff[times[i]]  # 累加变化量
                i += 1
            ans[id] = sum
        return ans
```

```java [sol1-Java]
class Solution {
    public int[] fullBloomFlowers(int[][] flowers, int[] persons) {
        var diff = new HashMap<Integer, Integer>();
        for (var f : flowers) {
            diff.put(f[0], diff.getOrDefault(f[0], 0) + 1);
            diff.put(f[1] + 1, diff.getOrDefault(f[1] + 1, 0) - 1);
        }
        var times = diff.keySet().stream().mapToInt(Integer::intValue).sorted().toArray();

        var n = persons.length;
        var ids = IntStream.range(0, n).boxed().toArray(Integer[]::new);
        Arrays.sort(ids, (i, j) -> persons[i] - persons[j]);

        var ans = new int[n];
        int i = 0, sum = 0;
        for (var id : ids) {
            while (i < times.length && times[i] <= persons[id])
                sum += diff.get(times[i++]); // 累加变化量
            ans[id] = sum;
        }
        return ans;
    }
}
```

```C++ [sol1-C++]
class Solution {
public:
    vector<int> fullBloomFlowers(vector<vector<int>> &flowers, vector<int> &persons) {
        map<int, int> diff;
        for (auto &f : flowers) {
            ++diff[f[0]];
            --diff[f[1] + 1];
        }

        int n = persons.size();
        vector<int> id(n);
        iota(id.begin(), id.end(), 0);
        sort(id.begin(), id.end(), [&](int i, int j) { return persons[i] < persons[j]; });

        vector<int> ans(n);
        auto it = diff.begin();
        int sum = 0;
        for (int i : id) {
            while (it != diff.end() && it->first <= persons[i])
                sum += it++->second; // 累加变化量
            ans[i] = sum;
        }
        return ans;
    }
};
```

```go [sol1-Go]
func fullBloomFlowers(flowers [][]int, persons []int) []int {
	diff := map[int]int{}
	for _, f := range flowers {
		diff[f[0]]++
		diff[f[1]+1]--
	}

	n := len(diff)
	times := make([]int, 0, n)
	for t := range diff {
		times = append(times, t)
	}
	sort.Ints(times)

	for i, p := range persons {
		persons[i] = p<<32 | i
	}
	sort.Ints(persons)

	ans := make([]int, len(persons))
	i, sum := 0, 0
	for _, p := range persons {
		for ; i < n && times[i] <= p>>32; i++ {
			sum += diff[times[i]] // 累加变化量
		}
		ans[uint32(p)] = sum
	}
	return ans
}
```

### 方法二：转换 + 二分

第 $i$ 个人能看到的花的数目，等价于 $\textit{start}$ 不晚于 $\textit{persons}_i$ 的花的数目，减去 $\textit{end}$ 早于 $\textit{persons}_i$ 的花的数目，即开花数减去凋落数。

所以单独统计开花时间和凋落时间，排序后二分就得到了答案。

- 时间复杂度：$O((n+m)\log n)$，其中 $n$ 是 $\textit{flowers}$ 的长度，$m$ 是 $\textit{persons}$ 的长度。
- 空间复杂度：$O(n)$。不计返回值的空间。

```Python [sol2-Python3]
class Solution:
    def fullBloomFlowers(self, flowers: List[List[int]], persons: List[int]) -> List[int]:
        starts = sorted(s for s, _ in flowers)
        ends = sorted(e for _, e in flowers)
        return [bisect_right(starts, p) - bisect_left(ends, p) for p in persons]
```

```java [sol2-Java]
class Solution {
    public int[] fullBloomFlowers(int[][] flowers, int[] persons) {
        var starts = Arrays.stream(flowers).mapToInt(f -> f[0]).sorted().toArray();
        var ends = Arrays.stream(flowers).mapToInt(f -> f[1]).sorted().toArray();
        return Arrays.stream(persons).map(p -> lowerBound(starts, p + 1) - lowerBound(ends, p)).toArray();
    }

    int lowerBound(int[] arr, int x) {
        int left = 0, right = arr.length;
        while (left < right) {
            var mid = (left + right) / 2;
            if (arr[mid] >= x) right = mid;
            else left = mid + 1;
        }
        return left;
    }
}
```

```C++ [sol2-C++]
class Solution {
public:
    vector<int> fullBloomFlowers(vector<vector<int>> &flowers, vector<int> &persons) {
        int n = flowers.size();
        vector<int> starts(n), ends(n);
        for (int i = 0; i < n; ++i) {
            starts[i] = flowers[i][0];
            ends[i] = flowers[i][1];
        }
        sort(starts.begin(), starts.end());
        sort(ends.begin(), ends.end());

        n = persons.size();
        vector<int> ans(n);
        for (int i = 0; i < n; ++i)
            ans[i] = (upper_bound(starts.begin(), starts.end(), persons[i]) - starts.begin()) -
                     (lower_bound(ends.begin(), ends.end(), persons[i]) - ends.begin());
        return ans;
    }
};
```

```go [sol2-Go]
func fullBloomFlowers(flowers [][]int, persons []int) []int {
	n := len(flowers)
	starts := make([]int, n)
	ends := make([]int, n)
	for i, f := range flowers {
		starts[i] = f[0]
		ends[i] = f[1]
	}
	sort.Ints(starts)
	sort.Ints(ends)

	ans := make([]int, len(persons))
	for i, p := range persons {
		ans[i] = sort.SearchInts(starts, p+1) - sort.SearchInts(ends, p)
	}
	return ans
}
```
