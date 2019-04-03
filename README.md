# golang-making-progress
golang精进

1，是否可以编译通过？如果通过，输出什么？
``` golang
func main() {
	i := GetValue()

	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}
func GetValue() int {
	return 1
}
```
解析
考点：type
编译失败，因为type只能使用在interface

2，下面函数有什么问题？
``` golang
func funcMui(x,y int)(sum int,error){
    return x+y,nil
}
```
解析
考点：函数返回值命名
在函数有多个返回值时，只要有一个返回值有指定命名，其他的也必须有命名。 如果返回值有有多个返回值必须加上括号； 如果只有一个返回值并且有命名也需要加上括号； 此处函数第一个返回值有sum名称，第二个未命名，所以错误。

3，是否可以编译通过？如果通过，输出什么？
``` golang
package main
func main() {

}
func DeferFunc1(i int) (t int) {
    t = i
    defer func() {
        t += 3
    }()
    return t
}
func DeferFunc2(i int) int {
    t := i
    defer func() {
        t += 3
    }()
    return t
}
func DeferFunc3(i int) (t int) {
    defer func() {
        t += i
    }()
    return 2
```
解析
考点:defer和函数返回值
需要明确一点是defer需要在函数结束前执行。 函数返回值名字会在函数起始处被初始化为对应类型的零值并且作用域为整个函数 DeferFunc1有函数返回值t作用域为整个函数，在return之前defer会被执行，所以t会被修改，返回4; DeferFunc2函数中t的作用域为函数，返回1; DeferFunc3返回3

3，是否可以编译通过？如果通过，输出什么？
``` golang
func main() {
    list := new([]int)
    list = append(list, 1)
    fmt.Println(list)
}
```
解析
考点：new
list:=make([]int,0)

4，是否可以编译通过？如果通过，输出什么？
``` golang
package main
import "fmt"
func main() {
    s1 := []int{1, 2, 3}
    s2 := []int{4, 5}
    s1 = append(s1, s2)
    fmt.Println(s1)
}
```
解析
考点：append
append切片时候别漏了'...'

5，是否可以编译通过？如果通过，输出什么？
``` golang
func main() {
    sn1 = struct{
        age int
        name string
    }{age: 11, name: "qq"}
    sn2 := struct {
        age int 
        name string
    }{age: 11, name: "qq"}
    if sn1 == sn2 {
        fmt.Println("sn1 == sn2")
    }

    sm1 := struct {
        age int
        m map[string]string
    }{age: 11, m: map[string]string{"a": "1"}}
    sm2 := struct {
        age int
        m map[string]string
    }{age: 11, m: map[string]string{"a": "1"}}
    if sm1 == sm2 {
        fmt.Println("sm1 == sm2")
    }
}
```
解析
考点:结构体比较
进行结构体比较时候，只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关。
``` golang
sn3:= struct {
    name string
    age  int
}{age:11,name:"qq"}
```
sn3与sn1就不是相同的结构体了，不能比较。 还有一点需要注意的是结构体是相同的，但是结构体属性中有不可以比较的类型，如map,slice。 如果该结构属性都是可以比较的，那么就可以使用“==”进行比较操作。
可以使用reflect.DeepEqual进行比较
``` golang
if reflect.DeepEqual(sn1, sm) {
    fmt.Println("sn1 ==sm")
}else {
    fmt.Println("sn1 !=sm")
}
```
所以编译不通过： invalid operation: sm1 == sm2

6，是否可以编译通过？如果通过，输出什么？
``` golang
func Foo(x interface{}) {
    if x == nil {
        fmt.Println("empty interface")
        return
    }
    fmt.Pritnln("non-empty interface")
}
func main() {
    var x *int = nil
    Foo(x)
}
```