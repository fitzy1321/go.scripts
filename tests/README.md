# Golang testing

How do golang unit tests?

Testing idea

Which is "better", using a channel or mutex to update values, somewhere?

Just say some youtube golang concurrency videos, and one of those "creators" used a Mutex to update a map.

My thought: "wouldn't it be easier and or better to use a channel, taking single KV pairs, and update a map in blocking code, i.e. outside go routine?

```go
import "sync"

type Pair[T, U any] struct {
  First  T
  Second U
}

func runSomeFunc(c chan Pair[string, int]) {
  for i := range 10 {
    c <- Pair[string, int]{First: fmt.Sprintf("key_{%d}", i), Second: i}
  }
}

func main() {
  kv_chan := make(chan Pair[string, int])

  values := map[string]int{}

  go runSomeFunc(kv_chan)

  for kv := range kv_chan {
    values[kv.First] = kv.Second
  }
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
```

that's the gist. How does a channel of KV Pairs compare to Mutexing around a map?

Performance, does it matter if not at scale? What even would be 'at scale' for this kinda code?
