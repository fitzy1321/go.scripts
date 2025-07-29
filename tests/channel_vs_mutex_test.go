package channel_vs_mutex_tests

import (
	"fmt"
	"sync"
	"testing"
)

type Pair[T, U any] struct {
	First  T
	Second U
}

type MutexMap[T comparable, U any] struct {
	Values map[T]U
	mu     sync.Mutex
}

func (m *MutexMap[T, U]) Add(k T, v U) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Values[k] = v
}

func NewMutecMap() *MutexMap[string, int] {
	return &MutexMap[string, int]{Values: make(map[string]int)}
}

func TestChannel(t *testing.T) {
	kv_chan := make(chan Pair[string, int])

	values := map[string]int{}

	go func() {
		defer close(kv_chan)
		wg := sync.WaitGroup{}
		for i := range 10 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				kv_chan <- Pair[string, int]{First: fmt.Sprintf("key_%d", i), Second: i}
			}()

			// golang v1.25 WaitGroup.Go()
			// replaces `wg.Add(1); go func() {defer wg.Done(); ... }();`.
			// wg.Go(func() {
			// 	kv_chan <- Pair[string, int]{First: fmt.Sprintf("key_%d", i), Second: i}
			// })
		}
		wg.Wait()
	}()

	for kv := range kv_chan {
		t.Logf("Key Value Pair in main goroutine: '%s', %d", kv.First, kv.Second)
		values[kv.First] = kv.Second
	}
	t.Log("Test Passed")
}

func TestMutex(t *testing.T) {
	values := NewMutecMap()

	wg := sync.WaitGroup{}
	for i := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			values.Add(fmt.Sprintf("key_%d", i), i)
		}()
	}
	wg.Wait()
	for k, v := range values.Values {
		t.Logf("KV Pair in main goroutine: '%s', %d", k, v)
	}
	t.Log("Test Passed")
}
