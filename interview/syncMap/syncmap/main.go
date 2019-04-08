// https://www.cnblogs.com/ghj1976/archive/2013/04/27/3047528.html
package syncmap

import "sync"

// SyncMap SyncMap
type SyncMap struct {
	lock *sync.RWMutex
	bm   map[interface{}]interface{}
}

// NewSyncMap NewSyncMap
func NewSyncMap() *SyncMap {
	return &SyncMap{
		lock: new(sync.RWMutex),
		bm:   make(map[interface{}]interface{}),
	}
}

// Get Get
func (m *SyncMap) Get(k interface{}) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val, ok := m.bm[k]; ok {
		return val
	}
	return nil
}

// Set set
func (m *SyncMap) Set(k interface{}, v interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if val, ok := m.bm[k]; !ok {
		m.bm[k] = v
	} else if val != v {
		m.bm[k] = v
	} else {
		return false
	}
	return true
}

// Check check
func (m *SyncMap) Check(k interface{}) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if _, ok := m.bm[k]; !ok {
		return false
	}
	return true
}

// Delete Delete
func (m *SyncMap) Delete(k interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.bm, k)
}
