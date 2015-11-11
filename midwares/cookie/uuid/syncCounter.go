package uuid

import (
    "sync/atomic"
)

type SyncCounter struct {
    counter int64
}

func (this *SyncCounter) Inc() int64 {
    return atomic.AddInt64(&this.counter, 1)
}

func (this *SyncCounter) Get() int64 {
    return atomic.LoadInt64(&this.counter)
}

func (this *SyncCounter) Dec() int64 {
    return atomic.AddInt64(&this.counter, -1)
}

func (this *SyncCounter) Set(value int64) {
    atomic.StoreInt64(&this.counter, value)
}
