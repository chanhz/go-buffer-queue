package buffer_queue

import (
	"errors"
)

const Max = 1000000

type element interface{}
type message chan element

type Queue struct {
	Message message
}

func (self *Queue) Init() {
	self.Message = make(message, Max)
}

func (self *Queue) Put(item interface{}) (int, error) {
	l := len(self.Message)
	if l < Max {
		self.Message <- item
		return l + 1, nil
	} else {
		return l, errors.New("The queue was full")
	}
}

func (self *Queue) Pop() (interface{}, int, error) {
	l := len(self.Message)
	if l > 0 {
		res := <-self.Message
		return res, l - 1, nil
	} else {
		return nil, 0, errors.New("The queue was empty")
	}
}
