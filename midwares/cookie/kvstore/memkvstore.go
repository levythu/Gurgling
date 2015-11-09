package kvstore

type MemKvStore map[string]string
// implementing kvstore

func (this MemKvStore)Set(key string, val string) error {
    this[key]=val
    return nil
}

func (this MemKvStore)Get(key string) string {
    var res, ok=this[key]
    if ok {
        return res
    }
    return ""
}
