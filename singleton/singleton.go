package counter

import (
	"sync"
)

type SingleCounter struct {
	count int
}

var instance *SingleCounter
var once sync.Once

func GetInstance() *SingleCounter {
	once.Do(func() {
		instance = &SingleCounter{}
	})
	return instance
}

func (s *SingleCounter) Increment() int {
	s.count++
	return s.count
}
