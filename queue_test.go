package buffer_queue

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQueue_Put(t *testing.T) {
	var q Queue
	q.Init()
	for len(q.Message) < Max {
		data := rand.Intn(100)
		if l, err := q.Put(data); err != nil {
			t.Error(err)
		} else {
			fmt.Println("queue's length:", l, ",added:", data)
		}
	}
	close(q.Message)
	for len(q.Message) > 0 {
		<-q.Message
	}
}

func BenchmarkQueue_Put(b *testing.B) {
	var q Queue
	q.Init()
	for i := 0; i < b.N; i++ {
		data := rand.Intn(100)
		if l, err := q.Put(data); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("queue's length:", l, ",added:", data)
		}
	}
	close(q.Message)
	for len(q.Message) > 0 {
		<-q.Message
	}
}

func TestQueue_Pop(t *testing.T) {
	var q Queue
	q.Init()
	for len(q.Message) < Max {
		data := rand.Intn(100)
		if l, err := q.Put(data); err != nil {
			t.Error(err)
		} else {
			fmt.Println("queue's length:", l, ",added:", data)

		}
	}
	close(q.Message)
	for len(q.Message) > 0 {
		if data, l, err := q.Pop(); err != nil {
			t.Error(err)
		} else {
			fmt.Println("queue's length:", l, ",consumed:", data)
		}
	}
}


func PopInit() *Queue{
	var q Queue
	q.Init()
	for i := 0; i < Max; i++ {
		data := rand.Intn(100)
		if l, err := q.Put(data); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("queue's length:", l, ",added:", data)
		}
	}
	close(q.Message)
	return &q
}
var mq *Queue
func init()  {
	mq = PopInit()
}
func BenchmarkQueue_Pop(b *testing.B) {
	q := mq
	for i := 0; i < b.N; i++ {
		if data, l, err := q.Pop(); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("queue's length:", l, ",consumed:", data)
		}
	}
}
