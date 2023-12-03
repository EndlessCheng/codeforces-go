下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

---

遍历下标在 $[1,n-2]$ 内的所有数，如果其大于其左右两侧相邻数字，则把 $i$ 加入答案。

```py [sol-Python3]
class Solution:
    def findPeaks(self, a: List[int]) -> List[int]:
        return [i for i in range(1, len(a) - 1) if a[i - 1] < a[i] > a[i + 1]]
```

```java [sol-Java]
class Solution {
    public List<Integer> findPeaks(int[] mountain) {
        List<Integer> ans = new ArrayList<>();
        for (int i = 1; i < mountain.length - 1; i++) {
            if (mountain[i] > mountain[i - 1] && mountain[i] > mountain[i + 1]) {
                ans.add(i);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findPeaks(vector<int> &mountain) {
        vector<int> ans;
        for (int i = 1; i + 1 < mountain.size(); i++) {
            if (mountain[i] > mountain[i - 1] && mountain[i] > mountain[i + 1]) {
                ans.push_back(i);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findPeaks(mountain []int) (ans []int) {
	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			ans = append(ans, i)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{mountain}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。
