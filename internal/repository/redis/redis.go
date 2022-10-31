package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"github.com/pkg/errors"
	"strings"
	"time"
)

var _ Repo = (*cacheRepo)(nil)

type Repo interface {
	i()
	Set(key, value string, ttl time.Duration) error
	Get(key string) (string, error)
	TTL(key string) (time.Duration, error)
	Expire(key string, ttl time.Duration) bool
	ExpireAt(key string, ttl time.Time) bool
	Del(key string) bool
	Exists(keys ...string) bool
	Incr(key string) int64
	Close() error
	Version() string
	GetOBJ(key string, dst interface{}) error
	SetObj(key string, value interface{}, ttl time.Duration) error
	GetKeysList(key string, list interface{}) error
	GetSet(key string) ([]string, error)
	AddSet(key string, item ...string) error
}

type cacheRepo struct {
	client *redis.Client
	ctx    context.Context
}

func New(Addr, Pass string, db int) (Repo, error) {
	client, err := redisConnect(Addr, Pass, db)
	if err != nil {
		return nil, err
	}
	return &cacheRepo{
		client: client,
		ctx:    context.Background(),
	}, nil
}

func (c *cacheRepo) i() {}

func redisConnect(Addr, Pass string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Pass,
		DB:       db,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, errors.Wrap(err, "ping redis err")
	}
	return client, nil
}

// Set set some <key,value> into redis
func (c *cacheRepo) Set(key, value string, ttl time.Duration) error {
	if err := c.client.Set(c.ctx, key, value, ttl).Err(); err != nil {
		return errors.Wrapf(err, "redis set key: %s err", key)
	}
	return nil
}

// Get get some key from redis
func (c *cacheRepo) Get(key string) (string, error) {
	value, err := c.client.Get(c.ctx, key).Result()
	if err != nil {
		return "", errors.Wrapf(err, "redis get key: %s err", key)
	}

	return value, nil
}

// TTL get some key from redis
func (c *cacheRepo) TTL(key string) (time.Duration, error) {
	ttl, err := c.client.TTL(c.ctx, key).Result()
	if err != nil {
		return -1, errors.Wrapf(err, "redis get key: %s err", key)
	}

	return ttl, nil
}

// Expire expire some key
func (c *cacheRepo) Expire(key string, ttl time.Duration) bool {
	ok, _ := c.client.Expire(c.ctx, key, ttl).Result()
	return ok
}

// ExpireAt expire some key at some time
func (c *cacheRepo) ExpireAt(key string, ttl time.Time) bool {
	ok, _ := c.client.ExpireAt(c.ctx, key, ttl).Result()
	return ok
}

func (c *cacheRepo) Exists(keys ...string) bool {
	if len(keys) == 0 {
		return true
	}
	value, _ := c.client.Exists(c.ctx, keys...).Result()
	return value > 0
}

func (c *cacheRepo) Del(key string) bool {
	if key == "" {
		return true
	}
	value, _ := c.client.Del(c.ctx, key).Result()
	return value > 0
}

func (c *cacheRepo) Incr(key string) int64 {
	value, _ := c.client.Incr(c.ctx, key).Result()
	return value
}

// Close redis client
func (c *cacheRepo) Close() error {
	return c.client.Close()
}

// Version redis services version
func (c *cacheRepo) Version() string {
	server := c.client.Info(c.ctx, "services").Val()
	spl1 := strings.Split(server, "# Server")
	spl2 := strings.Split(spl1[1], "redis_version:")
	spl3 := strings.Split(spl2[1], "redis_git_sha1:")
	return spl3[0]
}

func (c *cacheRepo) GetOBJ(key string, dst interface{}) error {
	get, err := c.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(get), dst)
}

func (c *cacheRepo) SetObj(key string, value interface{}, ttl time.Duration) error {
	marshal, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(key, string(marshal), ttl)
}
func (c *cacheRepo) GetKeysList(key string, list interface{}) error {
	get := c.client.Keys(c.ctx, key)
	mGet := c.client.MGet(c.ctx, get.Val()...)
	result, err := mGet.Result()
	arr := make([]map[string]interface{}, 0)
	for _, i := range result {
		if v, ok := i.(string); ok {
			m := make(map[string]interface{})
			_ = json.Unmarshal([]byte(v), &m)
			arr = append(arr, m)
		}
	}
	marshal, err := json.Marshal(arr)
	if err != nil {
		return err
	}
	return json.Unmarshal(marshal, list)
}
func (c *cacheRepo) GetSet(key string) ([]string, error) {
	member := c.client.SMembers(c.ctx, key)
	return member.Result()
}
func (c *cacheRepo) AddSet(key string, item ...string) error {
	member := c.client.SAdd(c.ctx, key, item)
	return member.Err()
}
