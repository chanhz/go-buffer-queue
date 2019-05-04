package buffer_queue

import (
	"math/rand"
	"sync"
)

type Queue struct {
	Data sync.Pool
}

func (self *Queue) Init() {
	self.Data = sync.Pool{
		New: func() interface{} {
			return rand.Intn(100)
		},
	}
}

func (self *Queue) Put(item interface{}) {
	self.Data.Put(item)
}

func(self *Queue) Pop() interface{} {
	return self.Data.Get()
}