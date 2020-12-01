package config

import (
	"database/sql"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/garyburd/redigo/redis"
	"log"
	"sync"
)

// 日志相关配置
const LogPath string = "log/"
var Loger *log.Logger

// redis 相关配置
var (
	once sync.Once
	conn redis.Conn
)

// 数据库 相关配置
var (
	once2 sync.Once
	MysqlDb *sql.DB
	MysqlDbErr error
)

type RedisCfg struct {
	redisHost string
	redisPort int
	redisSize int
}

type MysqlCfg struct {
	dbUserName string
	dbPassword string
	dbHost string
	dbPort string
	dbDataBase string
	dbCharset string
}

type GlobalCfg struct {
	Rcfg RedisCfg
	DBcfg MysqlCfg
}


//服务全局变量
var (
	 APP_NAME = "VideoAccessServer"
	 MUR_HASH_SEED uint32 = 0x12345678
)
var once3 sync.Once
var Cfg GlobalCfg

func init() {
	GetGlobalConfig()
}

func GetGlobalConfig()  {
	once3.Do(func() {
		c, err := goconfig.LoadConfigFile("./conf.ini")
		if err != nil {
			Loger.Printf("read conf file fail:%v", err)
			return
		}
		Cfg.Rcfg.redisHost, _ = c.GetValue("redis", "redisHost")
		Cfg.Rcfg.redisPort, _ = c.Int("redis", "redisPort")
		Cfg.Rcfg.redisSize, _ = c.Int("redis", "redisPoolSize")

		Cfg.DBcfg.dbUserName, _ = c.GetValue("mysql", "dbUserName")
		Cfg.DBcfg.dbPassword, _ = c.GetValue("mysql", "dbPassWord")
		Cfg.DBcfg.dbHost, _ = c.GetValue("mysql", "dbHost")
		Cfg.DBcfg.dbPort, _ = c.GetValue("mysql", "dbPort")
		Cfg.DBcfg.dbDataBase, _ = c.GetValue("mysql", "dbDataBase")
		Cfg.DBcfg.dbCharset, _ = c.GetValue("mysql", "dbCharset")

		fmt.Printf("get cfg:%v", Cfg)
	})

}

func GetRedisClientInstance() redis.Conn {
	once.Do(func() {
		var err error
		conn, err = redis.Dial("tcp", fmt.Sprintf("%s:%d", Cfg.Rcfg.redisHost,
			Cfg.Rcfg.redisPort))
		if err != nil {
			Loger.Fatalf("get redis instance fail:%v", err)
		}
	})
	return conn
}

func GetMysqlClientInstance() *sql.DB {
	once2.Do(func() {
		conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", Cfg.DBcfg.dbUserName, Cfg.DBcfg.dbPassword,
			Cfg.DBcfg.dbHost,Cfg.DBcfg.dbPort, Cfg.DBcfg.dbDataBase)

		// 打开连接
		MysqlDb, MysqlDbErr = sql.Open("mysql", conn)
		if MysqlDbErr != nil {
			Loger.Fatalf("connect mysql fail:%v", MysqlDbErr)
			panic("mysql conn error")
		}
	})
	return MysqlDb
}