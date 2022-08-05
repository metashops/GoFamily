package main

import (
	"sync"
)
type RWMap struct {
	sync.Map
	sync.RWMutex
	m map[int]int
}
func NewRWMap(n int) *RWMap {
	return &RWMap{
		m: make(map[int]int,n),
	}
}
func (m *RWMap) Get(k int) (int,bool) {//获取值
	m.RLock()
	defer m.RUnlock()
	v,existed := m.m[k]
	return v,existed
}
func (m *RWMap) Set(k,v int) {
	m.RLock()
	defer m.RUnlock()
	m.m[k] = v
}
func (m *RWMap)Delete(k int){
	m.RLock()
	defer m.RUnlock()
	delete(m.m,k)
}
func (m *RWMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}
func (m *RWMap) Each(f func(k,v int) bool) {
	m.RLock()
	defer m.RUnlock()
	for k,v := range m.m {
		if !f (k,v) {
			return
		}
	}

}
