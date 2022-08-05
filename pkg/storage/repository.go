package storage

type Bucket string

type Storage interface {
	Get(key string, bucket Bucket) (string, error)
	Set(key, value string, bucket Bucket) error
	Delete(id string, bucket Bucket)
}

const (
	Cache Bucket = "cache"
)
