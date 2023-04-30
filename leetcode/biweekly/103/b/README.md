下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

用位运算表示集合，两个数的 AND 就代表集合的交集，交集的大小就是二进制中 $1$ 的个数。

```py [sol1-Python3]
class Solution:
    def findThePrefixCommonArray(self, a: List[int], b: List[int]) -> List[int]:
        ans = []
        p = q = 0
        for x, y in zip(a, b):
            p |= 1 << x
            q |= 1 << y
            ans.append((p & q).bit_count())
        return ans
```

```java [sol1-Java]
class Solution {
    public int[] findThePrefixCommonArray(int[] a, int[] b) {
        int n = a.length;
        var ans = new int[n];
        long p = 0, q = 0;
        for (int i = 0; i < n; ++i) {
            p |= 1L << a[i];
            q |= 1L << b[i];
            ans[i] = Long.bitCount(p & q);
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<int> findThePrefixCommonArray(vector<int> &a, vector<int> &b) {
        int n = a.size();
        vector<int> ans(n);
        long long p = 0, q = 0;
        for (int i = 0; i < n; ++i) {
            p |= 1LL << a[i];
            q |= 1LL << b[i];
            ans[i] = __builtin_popcountll(p & q);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func findThePrefixCommonArray(a, b []int) []int {
	ans := make([]int, len(a))
	var p, q uint
	for i, x := range a {
		p |= 1 << x
		q |= 1 << b[i]
		ans[i] = bits.OnesCount(p & q)
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
