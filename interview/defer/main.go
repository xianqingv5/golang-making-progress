// Go：return 与 defer的执行顺序
// https://www.jianshu.com/p/67e40938d6b2
package main

import "fmt"

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0 // result = 0, call defer, ret
}

func f2() (r int) {
	t := 3
	defer func() {
		t = t + 3
	}()
	return t // t = r = 3, call defer, ret
}

func f3() (r int) {
	defer func(r int) {
		r = r + 2
	}(r)
	return 1 // r = 1, call defer, ret
}

func test() (x int) {
	defer println("defer")
	return 200
}

func test2() (x int) {
	defer func() {
		x = 100
	}()
	x = 200
	return x
}

func test3() (z int) {
	defer func() {
		println("defer:", z)
		z += 100 // 修改命名返回值。
	}()
	return 100 // 实际执行次序：z = 100, call defer, ret
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())

	fmt.Println(test())
	fmt.Println(test2())
	fmt.Println(test3())
}

// 先对返回值进行赋值：MOVQ $0xc8, 0x30(SP)
// 执行 defer 语句: CALL runtime.deferreturn(SB)
// 执行RET指令(函数返回): RET

// 考点：defer的一些坑
// return xxx可以改写成以下规则：
// 返回值 = xxx
// 调用defer函数
// 空的return
// 输出：
// 1
// 3
// 1
