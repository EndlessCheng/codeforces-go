package main

// https://space.bilibili.com/206214
func kItemsWithMaximumSum(numOnes, numZeros, _, k int) int {
	if k <= numOnes {
		return k
	}
	if k <= numOnes+numZeros {
		return numOnes
	}
	return numOnes*2 + numZeros - k
}
