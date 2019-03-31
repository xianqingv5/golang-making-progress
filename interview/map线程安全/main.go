package main

import "sync"

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	ua := UserAges{}
	ua.Get("yourname")
}

// 考点：map线程安全
// 有可能出现：fatal error: concurrent map read and map write
// 修改， Get方法加锁：
//func (ua *UserAges) Get(name string) int {
//    ua.Lock()
//    defer ua.Unlock()
//    if age, ok := ua.ages[name]; ok {
//        return age
//    }
//    return -1
//}
