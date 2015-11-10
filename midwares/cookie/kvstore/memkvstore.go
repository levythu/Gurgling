package kvstore

import (
    "sync"
)

type MemKvStore struct {
    Map map[string]map[string]string
    lock sync.RWMutex
}
// implementing kvstore

func (this *MemKvStore)Set(key string, val map[string]string) error {
    this.lock.Lock()
    defer this.lock.Unlock()

    this.Map[key]=val
    return nil
}

func (this *MemKvStore)Get(key string) map[string]string {
    this.lock.RLock()
    defer this.lock.RUnlock()

    var res, ok=this.Map[key]
    if ok {
        return res
    }
    return nil
}

func (this *MemKvStore)Remove(key string) error {
    this.lock.Lock()
    defer this.lock.Unlock()

    delete(this.Map, key)
    return nil
}
