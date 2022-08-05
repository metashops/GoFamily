package main

import (
	"container/list"
	"sync"
)

// Weighted 定义个weighted结构体
type Weighted struct {
	size int64
	cur int64
	mu sync.Mutex
	waiters list.List
}
func Semaphore(n int64) *Weighted {
	W := &Weighted{size: n}
	return W
}

// TryAcquire Acquire 获取权重
func (s *Weighted) TryAcquire(n int64) bool {
	s.mu.Lock()
	success := s.size-s.cur >= n && s.waiters.Len() == 0
	if success {
		s.cur += n
	}
	s.mu.Unlock()
	return success
}
func (s *Weighted) Release(n int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cur -= n
}