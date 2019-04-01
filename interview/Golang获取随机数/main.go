package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 获取随机整数
	for i := 0; i < 5; i++ {
		fmt.Printf("%v ", rand.Int())
	}
	fmt.Println()

	// 生成随机的32位整数
	for i := 0; i < 5; i++ {
		fmt.Printf("%v ", rand.Int31())
	}
	fmt.Println()

	// 获取浮点型数[0.0, 1.0)之间
	for i := 0; i < 5; i++ {
		fmt.Printf("%v ", rand.Float32())
	}
	fmt.Println()

	// 注意，如果上面的代码重复执行几遍，会得到相同的随机数，这时候需要根据时间设置随机数的种子。
	// 根据时间设置随机数种子
	rand.Seed(int64(time.Now().Nanosecond()))
	// 获取指定范围内的随机数
	for i := 0; i < 5; i++ {
		fmt.Printf("%v ", rand.Intn(100))
	}
	fmt.Println()

}
