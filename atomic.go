package atomic

import "sync"

// AtomicValue holds an arbitrary
// value whose writes and reads
// are synchronized against a
// read-write lock
type AtomicValue struct {
    Value interface{}
    lock *sync.RWMutex
}

func New() *AtomicValue {
    return &AtomicValue{
        lock: &sync.RWMutex{},
    }
}

func NewWithValue(v interface{}) *AtomicValue {
    return &AtomicValue{
        Value: v,
        lock: &sync.RWMutex{},
    }
}

func (a *AtomicValue) Set(v interface{}) {
    a.lock.Lock()
    defer a.lock.Unlock()
    a.Value = v
}

func (a *AtomicValue) Get() interface{} {
    a.lock.RLock()
    defer a.lock.RUnlock()
    return a.Value
}
