package dao

import (
	"VideoAccessServer/internal/config"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/cast"
)

// 设置带有过期时间 的key
func SetRedisStringWithExpired(con redis.Conn, key string, value string , ttl int32) error {
	_, err := con.Do("SET", key, value, "EX", cast.ToString(ttl))
	if err != nil {
		fmt.Errorf("SET KEY Error")
		config.Loger.Fatalf("set key with expired fail:%v", err)
		return err
	}
	config.Loger.Println("set Key :", key, " with expired success")

	return nil
}

// 设置无过期时间的key
func SetRedisString(con redis.Conn, key string,value []byte ) error {
	_, err := con.Do("SET", key, value)
	if err != nil {
		fmt.Errorf("SET KEY Error")
		config.Loger.Fatalf("set key fail: %v",err)
		return err
	}
	config.Loger.Println("set Key :", key, " success")
	return nil
}

// 根据key 取values
func GetRedisString(con redis.Conn, key string) (string ,error) {
	value, err := redis.String(con.Do("GET", key))
	if err != nil {
		fmt.Errorf("GET KEY FAIL")
		config.Loger.Fatalf("get key fail:%v", err)
	}
	config.Loger.Println("get key success:, value:", value)
	return value, nil
}

// 判断 key 是否存在
func GetRedisKeyExist(con redis.Conn, key string) (bool, error) {
	exist , err := redis.Bool(con.Do("EXISTS", key))
	if err != nil {
		fmt.Errorf("GET KEY EXIST FAIL")
		config.Loger.Fatalf(" get key exist fail:%v", err)
		return false, err
	}
	return exist, nil
}

// 删除 Key
func DelRedisKey(con redis.Conn, key string) (error) {
	_, err := con.Do("DEL", key)
	if err != nil {
		fmt.Errorf("DEL KEY FAIL")
		config.Loger.Fatalf("del key fail:%v", err)
		return err
	}
	return nil
}