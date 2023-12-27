[本题视频讲解](https://www.bilibili.com/video/BV1rG411k72D/)

水平栅栏和垂直栅栏分开计算。

- 对于水平栅栏，计算出任意两个栅栏之间的距离，存到一个哈希表 $h$ 中。
- 对于垂直栅栏，计算出任意两个栅栏之间的距离，存到一个哈希表 $v$ 中。

答案就是 $h$ 和 $v$ 交集中的最大值的平方。如果不存在最大值，返回 $-1$。

```py [sol-Python3]
class Solution:
    def maximizeSquareArea(self, m: int, n: int, hFences: List[int], vFences: List[int]) -> int:
        h = self.f(hFences, m)
        v = self.f(vFences, n)
        ans = max(h & v, default=0)
        return ans ** 2 % 1_000_000_007 if ans else -1

    def f(self, a: List[int], mx: int) -> Set[int]:
        a.extend([1, mx])
        a.sort()
        return set(y - x for x, y in combinations(a, 2))
```

```java [sol-Java]
class Solution {
    public int maximizeSquareArea(int m, int n, int[] hFences, int[] vFences) {
        Set<Integer> h = f(hFences, m);
        Set<Integer> v = f(vFences, n);
        int ans = 0;
        for (int x : h) {
            if (v.contains(x)) {
                ans = Math.max(ans, x);
            }
        }
        return ans > 0 ? (int) ((long) ans * ans % 1_000_000_007) : -1;
    }

    private Set<Integer> f(int[] a, int mx) {
        int len = a.length;
        a = Arrays.copyOf(a, len + 2);
        a[len] = 1;
        a[len + 1] = mx;
        Arrays.sort(a);

        Set<Integer> set = new HashSet<>();
        for (int i = 0; i < a.length; i++) {
            for (int j = i + 1; j < a.length; j++) {
                set.add(a[j] - a[i]);
            }
        }
        return set;
    }
}
```

```cpp [sol-C++]
class Solution {
    unordered_set<int> f(vector<int> &a, int mx) {
        a.push_back(1);
        a.push_back(mx);
        sort(a.begin(), a.end());
        unordered_set<int> set;
        for (int i = 0; i < a.size(); i++) {
            for (int j = i + 1; j < a.size(); j++) {
                set.insert(a[j] - a[i]);
            }
        }
        return set;
    }

public:
    int maximizeSquareArea(int m, int n, vector<int> &hFences, vector<int> &vFences) {
        auto h = f(hFences, m);
        auto v = f(vFences, n);
        if (h.size() > v.size()) {
            swap(h, v);
        }
        int ans = 0;
        for (int x: h) {
            if (v.contains(x)) {
                ans = max(ans, x);
            }
        }
        return ans ? (long long) ans * ans % 1'000'000'007 : -1;
    }
};
```

```go [sol-Go]
func f(a []int, mx int) map[int]bool {
	a = append(a, 1, mx)
	slices.Sort(a)
	set := map[int]bool{}
	for i, x := range a {
		for _, y := range a[i+1:] {
			set[y-x] = true
		}
	}
	return set
}

func maximizeSquareArea(m, n int, hFences, vFences []int) int {
	h := f(hFences, m)
	v := f(vFences, n)
	ans := 0
	for x := range h {
		if v[x] {
			ans = max(ans, x)
		}
	}
	if ans == 0 {
		return -1
	}
	return ans * ans % 1_000_000_007
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(h^2+v^2)$，其中 $h$ 为 $\textit{hFences}$ 的长度，$v$ 为 $\textit{vFences}$ 的长度。
- 空间复杂度：$\mathcal{O}(h^2+v^2)$。

#### 相似题目

- [2943. 最大化网格图中正方形空洞的面积](https://leetcode.cn/problems/maximize-area-of-square-hole-in-grid/)
