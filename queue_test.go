package buffer_queue

import (
	"fmt"
	"math/rand"
	"runtime/debug"
	"testing"
)

func TestQueue_Put(t *testing.T) {
	//defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var q Queue
	q.Init()
	for i := 0 ; i < 100; i++ {
		q.Put(rand.Intn(100))
	}
	//debug.SetGCPercent(100)
	for i := 0 ; i < 100; i++ {
		data := q.Pop()
		fmt.Println(data)
	}


}

func BenchmarkQueue_Put(b *testing.B) {
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var q Queue
	q.Init()
	for i := 0 ; i < b.N; i ++ {
		q.Put(rand.Intn(100))
	}

}




func BenchmarkQueue_Pop(b *testing.B) {
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var q Queue
	q.Init()
	for i := 0; i < b.N; i++  {
		fmt.Println(q.Pop())
	}

}
