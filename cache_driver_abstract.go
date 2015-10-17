package yotsuba

type CacheDriver interface {
    Load(key string) []byte
    Save(key string, content []byte)
}
