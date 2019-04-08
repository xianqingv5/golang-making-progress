// package main

// import (
// 	"errors"
// 	"time"
// )

// // Message 消息
// type Message struct {
// 	// Contents
// 	// 消息内容
// }

// // Subscription 订阅
// type Subscription struct {
// 	ch    chan<- Message // 只能发送值
// 	Inbox chan Message
// }

// func (s *Subscription) Publish(msg Message) error {
// 	if _, ok := <-s.Inbox; !ok {
// 		return errors.New("Topic has been closed")
// 	}

// 	s.ch <- msg
// 	return nil
// }

// // 话题
// type Topic struct {
// 	Subscribers    []Session
// 	MessageHistory []Message
// }

// // Subscribe 订阅
// func (t *Topic) Subscribe(uid uint64) (Subscription, error) {
// 	// Get session and create one if it's the first
// 	// 给第一个获取会话并创建会话.
// 	// Add session to the Topic & MessageHistory
// 	// 加入会议的主题和messagehistory
// 	// Create a subscription
// 	// 创建一个订阅

// 	return nil, nil
// }

// // 取消订阅
// func (t *Topic) Unsubscribe(Subscription) error {
// 	// Implementation
// 	return nil
// }

// // 删除
// func (t *Topic) Delete() error {
// 	// Implementation
// 	return nil
// }

// // 用户
// type User struct {
// 	ID   uint64
// 	Name string
// }

// // 会话
// type Session struct {
// 	User      User
// 	Timestamp time.Time
// }

// func main() {

// }
