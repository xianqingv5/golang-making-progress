// https://github.com/jeanphorn/go-design-patterns/tree/master/behavioral_patterns
package main

import (
	"fmt"
	"time"
)

// Event 事件类型，发生事件驱动时将事件传递给观察者们
type Event struct {
	Data string
}

// Observer 观察者
type Observer interface {
	// 更新事件
	Update(*Event)
}

// Subject 被观察的对象接口
type Subject interface {
	// 注册观察者
	Regist(Observer)
	// 注销观察者
	Deregist(Observer)
	// 通知观察者事件
	Notify(*Event)
}

// ConcreteObserver 实现Subject接口
type ConcreteObserver struct {
	ID int
}

// Update update
func (co *ConcreteObserver) Update(e *Event) {
	fmt.Printf("observer [%d] recieved msg: %s.\n", co.ID, e.Data)
}

// ConcreteSubject ConcreteSubject
type ConcreteSubject struct {
	Observers map[Observer]struct{} // 该主题的观察者
}

// Regist 注册
func (cs *ConcreteSubject) Regist(ob Observer) {
	cs.Observers[ob] = struct{}{}
}

// Deregist 注销观察者
func (cs *ConcreteSubject) Deregist(ob Observer) {
	delete(cs.Observers, ob)
}

// Notify 通知每个观察者事件
func (cs *ConcreteSubject) Notify(e *Event) {
	for ob, _ := range cs.Observers {
		ob.Update(e)
	}
}

func main() {
	cs := &ConcreteSubject{
		Observers: make(map[Observer]struct{}),
	}

	// 实例化两个观察者
	cobserver1 := &ConcreteObserver{1}
	cobserver2 := &ConcreteObserver{2}

	// 注册观察者
	cs.Regist(cobserver1)
	cs.Regist(cobserver2)

	for i := 0; i < 5; i++ {
		e := &Event{fmt.Sprintf("msg [%d]", i)}
		cs.Notify(e)
		time.Sleep(time.Duration(3) * time.Second)
	}
}
