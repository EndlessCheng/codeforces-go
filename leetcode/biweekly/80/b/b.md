对于正整数，$xy\ge\textit{success}$ 等价于 $y\ge\left\lceil\dfrac{\textit{success}}{x}\right\rceil$，也等价于 $y>\left\lfloor\dfrac{\textit{success}-1}{x}\right\rfloor$

这样对 $\textit{potions}$ 排序后，就可以愉快地使用库函数二分 $\textit{potions}$ 了。

除了可以使用库函数二分外，另一个好处是每次二分只需要做一次除法，避免在二分内部做乘法，效率更高。

```Python [sol1-Python3]
class Solution:
    def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:
        potions.sort()
        return [len(potions) - bisect_right(potions, (success - 1) // x) for x in spells]
```

```C++ [sol1-C++]
class Solution {
public:
    vector<int> successfulPairs(vector<int> &spells, vector<int> &potions, long long success) {
        sort(potions.begin(), potions.end());
        for (auto &x : spells)
            x = potions.end() - upper_bound(potions.begin(), potions.end(), (success - 1) / x);
        return spells;
    }
};
```

```go [sol1-Go]
func successfulPairs(spells, potions []int, success int64) []int {
	sort.Ints(potions)
	for i, x := range spells {
		spells[i] = len(potions) - sort.SearchInts(potions, (int(success)-1)/x+1)
	}
	return spells
}
```
