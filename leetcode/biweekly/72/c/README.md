由于 $\textit{finalSum}$ 只能分解成偶数之和，而偶数+偶数=偶数，所以 $\textit{finalSum}$ 也必须是偶数。

既然要尽量多的分解，且分解出的偶数互不相同，那么可以按照 $2,4,6,8,\cdots$ 的顺序分解，一边分解一边减少 $\textit{finalSum}$，直到 $\textit{finalSum}$ 小于要分解出的数为止。

最后把剩余的 $\textit{finalSum}$ 加到最后一个分解出的偶数上，即完成了分解。

```py [sol-Python3]
class Solution:
    def maximumEvenSplit(self, finalSum: int) -> List[int]:
        if finalSum % 2:
            return []
        ans = []
        i = 2
        while i <= finalSum:
            ans.append(i)
            finalSum -= i
            i += 2
        ans[-1] += finalSum
        return ans
```

```java [sol-Java]
class Solution {
    public List<Long> maximumEvenSplit(long finalSum) {
        if (finalSum % 2 > 0) return List.of();
        var ans = new ArrayList<Long>();
        for (long i = 2; i <= finalSum; i += 2) {
            ans.add(i);
            finalSum -= i;
        }
        int back = ans.size() - 1;
        ans.set(back, ans.get(back) + finalSum);
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> maximumEvenSplit(long long finalSum) {
        if (finalSum % 2) return {};
        vector<long long> ans;
        for (long long i = 2; i <= finalSum; i += 2) {
            ans.push_back(i);
            finalSum -= i;
        }
        ans.back() += finalSum;
        return ans;
    }
};
```

```go [sol-Go]
func maximumEvenSplit(finalSum int64) (ans []int64) {
	if finalSum%2 == 0 {
		for i := int64(2); i <= finalSum; i += 2 {
			ans = append(ans, i)
			finalSum -= i
		}
		ans[len(ans)-1] += finalSum
	}
	return
}
```

```js [sol-JavaScript]
var maximumEvenSplit = function (finalSum) {
    if (finalSum % 2) 
        return [];
    let ans = [];
    for (let i = 2; i <= finalSum; i += 2) {
        ans.push(i);
        finalSum -= i;
    }
    ans[ans.length - 1] += finalSum;
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\sqrt{\textit{finalSum}})$。设数组长度为 $x$，根据等差数列之和，$2+4+6+8+\cdots+2x = x(x+1) \le \textit{finalSum}$，可知 $x$ 约为 $\sqrt{\textit{finalSum}}$，所以循环次数为 $\mathcal{O}(\sqrt{\textit{finalSum}})$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。
