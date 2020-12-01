package dao

import (
	"VideoAccessServer/internal/config"
	"database/sql"
)


// table_name
// CREATE TABLE `t_user_info` (
// `id` int unsigned NOT NULL AUTO_INCREMENT,
// `name` varchar(64) DEFAULT NULL,
// `age` int DEFAULT NULL,
// PRIMARY KEY (`id`)
// ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

// demo 插入操作
func InsertToDB(conn *sql.DB) {
	config.Loger.Println("insert Sql")
	res, err := conn.Exec("INSERT INTO t_user_info(name, age) values (?, ?)", "hawrk100", 40)
	if err != nil {
		config.Loger.Fatalf("insert into Fail:%v", err)
		return
	}
	config.Loger.Println("res:", res)
}

// demo 查询 操作
func QueryDB(conn *sql.DB) {
	config.Loger.Println(" select sql...")
	data, err := conn.Query("select name, age from t_user_info where 1=1")
	if err != nil {
		config.Loger.Fatalf("select sql fail :%v", err)
		return
	}
	userInfo := make(map[string]int32)
	var name string
	var age int32
	for data.Next() {
		data.Scan(&name, &age)
		config.Loger.Println("get name :", name , ",age:", age)
		userInfo[name] = age
	}
	config.Loger.Printf("get result:%v", userInfo)
	return
}
