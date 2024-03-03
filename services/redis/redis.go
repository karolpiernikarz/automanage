package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

var rdb *redis.Client

func Init(host, port, password string, db int) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})
}

func Set(key string, value interface{}, expiration time.Duration) (err error) {
	err = rdb.Set(ctx, key, value, expiration).Err()
	return
}

func IsExist(key string) (exist bool) {
	result, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	if result == 1 {
		return true
	}
	return
}

func HSet(key string, field string, value interface{}) (err error) {
	err = rdb.HSet(ctx, key, field, value).Err()
	return
}

func SAdd(key string, value interface{}) (err error) {
	err = rdb.SAdd(ctx, key, value).Err()
	return
}

func SetNX(key string, value interface{}, expiration time.Duration) (err error) {
	err = rdb.SetNX(ctx, key, value, expiration).Err()
	return
}

func Get(key string) (value string, err error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func Delete(key string) (err error) {
	err = rdb.Del(ctx, key).Err()
	return
}

func GetKeysWithPrefix(prefix string) (keys []string, err error) {
	keys, err = rdb.Keys(ctx, prefix+"*").Result()
	return
}

func GetKeysWithPrefixAndSuffix(prefix string, suffix string) (keys []string, err error) {
	keys, err = rdb.Keys(ctx, prefix+"*"+suffix).Result()
	return
}

func GetKeysWithSuffix(suffix string) (keys []string, err error) {
	keys, err = rdb.Keys(ctx, "*"+suffix).Result()
	return
}

func GetKeys() (keys []string, err error) {
	keys, err = rdb.Keys(ctx, "*").Result()
	return
}

func GetKeysWithPrefixAndSuffixAndLimit(prefix string, suffix string, limit int64) (keys []string, err error) {
	keys, err = rdb.Keys(ctx, prefix+"*"+suffix).Result()
	if err != nil {
		return nil, err
	}
	if limit < int64(len(keys)) {
		keys = keys[:limit]
	}
	return
}

func GetKeysWithPrefixAndLimit(prefix string, limit int64) (keys []string, err error) {
	keys, err = rdb.Keys(ctx, prefix+"*").Result()
	if err != nil {
		return nil, err
	}
	if limit < int64(len(keys)) {
		keys = keys[:limit]
	}
	return
}

func GetKeysWithSuffixAndLimit(suffix string, limit int64) (keys []string, err error) {
	keys, err = rdb.Keys(ctx, "*"+suffix).Result()
	if err != nil {
		return nil, err
	}
	if limit < int64(len(keys)) {
		keys = keys[:limit]
	}
	return
}

func GetKeysWithLimit(limit int64) (keys []string, err error) {
	keys, err = rdb.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}
	if limit < int64(len(keys)) {
		keys = keys[:limit]
	}
	return
}

func GetKeysWithPrefixAndSuffixAndOffset(prefix string, suffix string, offset int64) (keys []string, err error) {
	keys, err = rdb.Keys(ctx, prefix+"*"+suffix).Result()
	if err != nil {
		return nil, err
	}
	if offset < int64(len(keys)) {
		keys = keys[offset:]
	}
	return
}

func GetKeysWithPrefixAndOffset(prefix string, offset int64) (keys []string, err error) {
	keys, err = rdb.Keys(ctx, prefix+"*").Result()
	if err != nil {
		return nil, err
	}
	if offset < int64(len(keys)) {
		keys = keys[offset:]
	}
	return
}

func GetKeysWithSuffixAndOffset(suffix string, offset int64) (keys []string, err error) {
	keys, err = rdb.Keys(ctx, "*"+suffix).Result()
	if err != nil {
		return nil, err
	}
	if offset < int64(len(keys)) {
		keys = keys[offset:]
	}
	return
}

func GetKeysWithOffset(offset int64) (keys []string, err error) {
	keys, err = rdb.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}
	if offset < int64(len(keys)) {
		keys = keys[offset:]
	}
	return
}

func GetKeysWithPrefixAndSuffixAndLimitAndOffset(prefix string, suffix string, limit int64, offset int64) (keys []string, err error) {
	keys, err = rdb.Keys(ctx, prefix+"*"+suffix).Result()
	if err != nil {
		return nil, err
	}
	if limit < int64(len(keys)) {
		keys = keys[:limit]
	}
	if offset < int64(len(keys)) {
		keys = keys[offset:]
	}
	return
}
