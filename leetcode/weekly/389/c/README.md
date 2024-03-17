[视频讲解](https://www.bilibili.com/video/BV1RH4y1W7DP/) 第三题。

**定理**：必然有一种字母是不需要删除的。

**反证法**：如果每种字母都至少删除一个，那么可以都增加一，不影响字母数量之差。

统计 $\textit{word}$ 中每个字母的出现次数，记到一个数组 $\textit{cnt}$ 中。

枚举 $i$ 作为出现次数最小的字母，为了保留尽量多的字母，字母 $i$ 肯定不需要删除。此外，出现次数最多的字母，其出现次数不能超过 $\textit{cnt}[i]+k$。

分类讨论：

- 出现次数小于 $\textit{cnt}[i]$ 的字母，全部删除。
- 出现次数大于等于 $\textit{cnt}[i]$ 的字母 $j$，保留 $\min(\textit{cnt}[j], \textit{cnt}[i] + k)$ 个。累加保留的字母个数，更新最多保留的字母个数 $\textit{maxSave}$ 的最大值。

最后用 $\textit{word}$ 的长度，减去 $\textit{maxSave}$，即为答案。

代码实现时，为方便计算，把 $\textit{cnt}$ 从小到大排序。

```py [sol-Python3]
class Solution:
    def minimumDeletions(self, word: str, k: int) -> int:
        cnt = sorted(Counter(word).values())
        max_save = max(sum(min(c, base + k) for c in cnt[i:])
                       for i, base in enumerate(cnt))
        return len(word) - max_save
```

```java [sol-Java]
class Solution {
    public int minimumDeletions(String word, int k) {
        int[] cnt = new int[26];
        for (char c : word.toCharArray()) {
            cnt[c - 'a']++;
        }
        Arrays.sort(cnt);

        int maxSave = 0;
        for (int i = 0; i < 26; i++) {
            int sum = 0;
            for (int j = i; j < 26; j++) {
                sum += Math.min(cnt[j], cnt[i] + k); // 至多保留 cnt[i]+k 个
            }
            maxSave = Math.max(maxSave, sum);
        }
        return word.length() - maxSave;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumDeletions(string word, int k) {
        int cnt[26]{};
        for (char c: word) {
            cnt[c - 'a']++;
        }
        ranges::sort(cnt);

        int max_save = 0;
        for (int i = 0; i < 26; i++) {
            int sum = 0;
            for (int j = i; j < 26; j++) {
                sum += min(cnt[j], cnt[i] + k); // 至多保留 cnt[i]+k 个
            }
            max_save = max(max_save, sum);
        }
        return word.length() - max_save;
    }
};
```

```go [sol-Go]
func minimumDeletions(word string, k int) int {
	cnt := make([]int, 26)
	for _, b := range word {
		cnt[b-'a']++
	}
	slices.Sort(cnt)

	maxSave := 0
	for i, base := range cnt {
		sum := 0
		for _, c := range cnt[i:] {
			sum += min(c, base+k) // 至多保留 base+k 个
		}
		maxSave = max(maxSave, sum)
	}
	return len(word) - maxSave
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|^2)$，其中 $n$ 为 $\textit{word}$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)

更多题单，请点我个人主页 - 讨论发布。
