package kvstore

type MemKvStore map[string]string
// implementing kvstore

func (MemKvStore this)Set(key string, val string) error {
    this[key]=val
    return nil
}

func (MemKvStore this)Get(key string) string {
    var res, ok=this[key]
    if ok {
        return res
    }
    return ""
}
