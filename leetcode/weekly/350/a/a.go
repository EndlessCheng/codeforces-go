package main

// https://space.bilibili.com/206214
func distanceTraveled(mainTank, additionalTank int) int {
	return (mainTank + min((mainTank-1)/4, additionalTank)) * 10
}
