package redis

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/gomodule/redigo/redis"
// )

// type RedisConf struct {
// 	Enabled  bool   `int:"enabled"`
// 	URL      string `ini:"url"`
// 	MaxIdle  int    `ini:"maxIdle"`  //最大空闲连接数
// 	Password string `ini:"password"` //密码
// }

// type Redis struct {
// 	cli *redis.Pool
// 	cfg *RedisConf
// }

// func NewRedis(cfg *RedisConf) (r *Redis) {
// 	if !cfg.Enabled {
// 		log.Println("redis配置未启用.")
// 	}
// 	r = &Redis{}
// 	r.cfg = cfg
// 	r.cli = &redis.Pool{
// 		MaxIdle:     cfg.MaxIdle,
// 		IdleTimeout: 30 * time.Second,
// 		Dial: func() (redis.Conn, error) {
// 			c, err := redis.DialURL(cfg.URL)
// 			if err != nil {
// 				return nil, fmt.Errorf("redis connection error: %s", err)
// 			}
// 			//验证redis密码
// 			if _, authErr := c.Do("AUTH", cfg.Password); authErr != nil {
// 				return nil, fmt.Errorf("redis auth password error: %s", authErr)
// 			}
// 			return c, err
// 		},
// 		TestOnBorrow: func(c redis.Conn, t time.Time) error {
// 			_, err := c.Do("PING")
// 			if err != nil {
// 				return fmt.Errorf("ping redis error: %s", err)
// 			}
// 			return nil
// 		},
// 	}
// 	return
// }

// func (r *Redis) Set(k string, v string, timelen string) (er error) {
// 	c := r.cli.Get()
// 	defer c.Close()
// 	_, err := c.Do("SET", k, v, "EX", timelen)
// 	return err
// }

// func (r *Redis) Get(k string) (ret interface{}, err error) {
// 	c := r.cli.Get()
// 	defer c.Close()
// 	return c.Do("Get", k)
// }

// func (r *Redis) Del(k string) (er error) {
// 	c := r.cli.Get()
// 	defer c.Close()
// 	_, err := c.Do("del", k)
// 	return err
// }

// func (r *Redis) HSet(p ...interface{}) (er error) {
// 	c := r.cli.Get()
// 	defer c.Close()
// 	_, err := c.Do("HSet", p...)
// 	return err
// }

// func (r *Redis) HGetAll(k string) (ret interface{}, err error) {
// 	c := r.cli.Get()
// 	defer c.Close()
// 	return c.Do("HGetAll", k)
// }

// func (r *Redis) HGet(k string, field string) (string, error) {
// 	c := r.cli.Get()
// 	defer c.Close()
// 	return redis.String(c.Do("HGet", k, field))
// }

// func (r *Redis) HMGET(k string, p ...interface{}) (ret []string, err error) {
// 	c := r.cli.Get()
// 	defer c.Close()
// 	return redis.Strings(c.Do("HMGET", p...))
// }

// func (r *Redis) TTL(k string) (ret interface{}, err error) {
// 	c := r.cli.Get()
// 	defer c.Close()
// 	return c.Do("TTL", k)
// }

// func (r *Redis) HGetJson(k string, field string) (ret []byte, err error) {
// 	return redis.Bytes(r.HGet(k, field))
// }

// func (r *Redis) Publish(channel string, messge interface{}) (ret interface{}, err error) {
// 	c := r.cli.Get()
// 	defer c.Close()
// 	return c.Do("Publish", channel, messge)
// }

// func (r *Redis) LPOP(k string) (string, error) {
// 	c := r.cli.Get()
// 	defer c.Close()
// 	return redis.String(c.Do("LPOP", k))
// }

// func (r *Redis) LLEN(k string) (int, error) {
// 	c := r.cli.Get()
// 	defer c.Close()
// 	return redis.Int(c.Do("LLEN", k))
// }

// func (r *Redis) RPUSH(p ...interface{}) (ret interface{}, err error) {
// 	c := r.cli.Get()
// 	defer c.Close()
// 	return c.Do("RPUSH", p...)
// }

// func (r *Redis) EXPIRE(k string, second int) (er error) {
// 	c := r.cli.Get()
// 	defer c.Close()
// 	_, err := c.Do("EXPIRE", k, second)
// 	return err
// }

// // SetEX Redis Setex 命令为指定的 key 设置值及其过期时间。
// // 如果 key 已经存在， SETEX 命令将会替换旧的值。
// func (r *Redis) SetEX(k string, second int, v string) (err error) {
// 	c := r.cli.Get()
// 	defer c.Close()
// 	_, err = c.Do("SETEX", k, second, v)
// 	return err
// }

// func (r *Redis) EXISTS(k string) (exist bool, err error) {
// 	c := r.cli.Get()
// 	defer c.Close()
// 	result, err := redis.Int(c.Do("EXISTS", k))
// 	if err != nil {
// 		return false, err
// 	}
// 	if result == 1 {
// 		return true, nil
// 	}
// 	return false, nil
// }
