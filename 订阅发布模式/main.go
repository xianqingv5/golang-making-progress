package main

import "fmt"

const (
	sub   int = iota // 订阅
	pub              // 发布
	unsub            // 取消订阅
)

type Pubsub struct {
	//command channel,multiplexing
	cmdChan  chan cmd
	capacity int
}

type cmd struct {
	op    int
	topic string
	ch    chan interface{}
	msg   interface{}
}

func PubServer(capacity int) *Pubsub {
	ps := &Pubsub{make(chan cmd), capacity}
	go ps.start()
	return ps
}

func (ps *Pubsub) Sub(topic string) chan interface{} {
	ch := make(chan interface{}, ps.capacity)
	ps.cmdChan <- cmd{op: sub, topic: topic, ch: ch}
	return ch
}

func (ps *Pubsub) Pub(msg interface{}, topic string) {
	ps.cmdChan <- cmd{op: pub, topic: topic, msg: msg}
}

//topic ->sublist
type registry struct {
	topics map[string][]chan interface{}
}

//bind sub channels to topics
func (reg *registry) add(topic string, ch chan interface{}) {
	if reg.topics[topic] == nil {
		reg.topics[topic] = make([]chan interface{}, 0, 5)
	}
	reg.topics[topic] = append(reg.topics[topic], ch)
}

//kick the ball
func (reg *registry) send(topic string, msg interface{}) {
	for _, ch := range reg.topics[topic] {
		ch <- msg
	}
}

//worker,ready to dispatch
func (ps *Pubsub) start() {
	reg := registry{
		topics: make(map[string][]chan interface{}),
	}
loop:
	for cmd := range ps.cmdChan {
		if cmd.topic == "" {
			continue loop
		}
		switch cmd.op {
		case sub:
			reg.add(cmd.topic, cmd.ch)
		case pub:
			reg.send(cmd.topic, cmd.msg)
		case unsub:
			//skipped,for it's no easy to remove an element given a slice
		}
	}

}

func main() {
	s := PubServer(1)
	ch1 := s.Sub("english")
	ch2 := s.Sub("french")
	ch3 := s.Sub("chinese")

	s.Pub("hello english", "english")
	s.Pub("hello french", "french")
	s.Pub(1.222222, "chinese")

	fmt.Printf("%v ", "hello english" == <-ch1)
	//shall be false
	fmt.Printf("%v ", "hello frenchxxx" == <-ch2)
	fmt.Printf("%v ", 1.222222 == <-ch3)
}
