package random

import (
	"math/rand"
	"time"
)

// GeneratesSpecifiedNumberOfDigitsRandom 生成指定位数的随机数
func GeneratesSpecifiedNumberOfDigitsRandom(digit int) int {
	// 设置随机种子
	seed := time.Now().UnixNano()
	src := rand.NewSource(seed)
	rnd := rand.New(src)

	// 生成指定位数的随机数
	min := int(pow10(digit - 1))
	max := int(pow10(digit) - 1)
	return rnd.Intn(max-min+1) + min
}

// 计算 10 的 n 次方
func pow10(n int) float64 {
	result := 1.0
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}

// GeneratesRandomNumbersOfSpecifiedSize 生成指定大小的随机数
func GeneratesRandomNumbersOfSpecifiedSize(min, max int) int {
	// 设置随机种子
	seed := time.Now().UnixNano()
	src := rand.NewSource(seed)
	rnd := rand.New(src)

	// 生成指定范围内的随机数
	return rnd.Intn(max-min+1) + min
}
