package kvstore

type KvStore interface {
    Set(key string, val string) error

    // Nonexist returns empty
    Get(key string) string
}
