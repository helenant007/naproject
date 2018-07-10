package redis

import "github.com/gomodule/redigo/redis"

// init redis
// devel-redis.tkpd:6379

func GET(key string) (interface{}, error) {
	conn, err := redis.Dial("tcp", "devel-redis.tkpd:6379")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	reply, err := conn.Do("GET", key)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func SET(key string, value interface{}) (interface{}, error) {
	conn, err := redis.Dial("tcp", "devel-redis.tkpd:6379")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	reply, err := conn.Do("SET", key, value)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func INCR(key string) (interface{}, error) {
	conn, err := redis.Dial("tcp", "devel-redis.tkpd:6379")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	reply, err := conn.Do("INCR", key)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
