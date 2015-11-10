package kvstore

type KvStore interface {
    Set(key string, val map[string]string) error

    // Nonexist returns nil
    Get(key string) map[string]string
}
