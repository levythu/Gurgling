package kvstore

type MemKvStore map[string]map[string]string
// implementing kvstore

func (this MemKvStore)Set(key string, val map[string]string) error {
    this[key]=val
    return nil
}

func (this MemKvStore)Get(key string) map[string]string {
    var res, ok=this[key]
    if ok {
        return res
    }
    return nil
}
