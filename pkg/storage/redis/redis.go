package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/fairytale5571/privat_test/pkg/logger"
	"github.com/fairytale5571/privat_test/pkg/storage"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	db     *redis.Client
	logger *logger.Wrapper
}

var ctx = context.Background()

func New(uri string) (*Redis, error) {
	res := &Redis{
		logger: logger.New("redis"),
	}

	opt, err := redis.ParseURL(uri)
	if err != nil {
		res.logger.Errorf("cant parse redis url %s", err.Error())
		return nil, err
	}
	res.db = redis.NewClient(opt)
	return res, nil
}

func (r *Redis) Get(key string, bucket storage.Bucket) (string, error) {
	return r.db.Get(ctx, fmt.Sprintf("%s::%s", bucket, key)).Result()
}

func (r *Redis) Set(key, value string, bucket storage.Bucket) error {
	return r.db.Set(ctx, fmt.Sprintf("%s::%s", bucket, key), value, 0).Err()
}

func (r *Redis) SetTTL(key, value string, bucket storage.Bucket, ttl time.Duration) error {
	return r.db.Set(ctx, fmt.Sprintf("%s::%s", bucket, key), value, ttl).Err()
}

func (r *Redis) Delete(id string, payments storage.Bucket) {
	if err := r.db.Del(ctx, fmt.Sprintf("%s::%s", payments, id)).Err(); err != nil {
		r.logger.Errorf("delete redis key %s", err.Error())
	}
}
