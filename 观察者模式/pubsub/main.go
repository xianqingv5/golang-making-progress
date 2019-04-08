// https://blog.csdn.net/xcl168/article/details/44355611
package pubsub

import "sync"

// Client 客户端
type Client struct {
	ID int
	IP string
}

type Server struct {
	// Dict map[string]*Channel
	sync.RWMutex
}
