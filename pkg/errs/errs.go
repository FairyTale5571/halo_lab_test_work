package errs

import "errors"

var (
	ErrorNotFound         = errors.New("not found")
	ErrorNotCached        = errors.New("not cached")
	ErrorCantCacheRedis   = errors.New("can't cache redis")
	ErrorCantCacheMemory  = errors.New("can't cache memory")
	ErrorCantDeleteMemory = errors.New("can't delete memory")
	ErrorCantDeleteRedis  = errors.New("can't delete redis")
)
