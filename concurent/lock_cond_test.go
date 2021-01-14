package concurent

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

type Queue struct {
	max    int
	buffer []int
	fill   int
	use    int
	count  int
	lock   sync.Mutex
	empty  sync.Cond
	full   sync.Cond
}

func (q *Queue) put(num int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	for q.count == q.max {
		q.empty.Wait()
	}
	q.buffer[q.fill] = num
	q.fill = (q.fill + 1) % q.max
	q.count += 1
	q.full.Signal()
}

func (q *Queue) get() (num int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	for q.count == 0 {
		q.full.Wait()
	}
	num = q.buffer[q.use]
	q.use = (q.use + 1) % q.max
	q.count -= 1
	q.empty.Signal()
	return
}

func NewQueue(max int) *Queue {
	q := &Queue{max: max, buffer: make([]int, max)}
	q.empty.L = &q.lock
	q.full.L = &q.lock
	return q
}

func produce(t *testing.T, stop *bool, queue *Queue, wg *sync.WaitGroup) {
	var i int
	for !*stop {
		queue.put(i)
		i += 1
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
	wg.Done()
}

func consume(t *testing.T, stop *bool, queue *Queue, wg *sync.WaitGroup) {
	for !*stop {
		t.Log(queue.get())
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
	wg.Done()
}

func TestProducerConsumer(t *testing.T) {
	stop := false
	q := NewQueue(5)
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go produce(t, &stop, q, wg)
		go consume(t, &stop, q, wg)
	}
	time.Sleep(6 * time.Second)
	stop = true
	wg.Wait()
}

func TestChannel(t *testing.T) {
	c := make(chan int)
	close(c)
	c = nil
}
