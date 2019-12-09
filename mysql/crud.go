package mysql

import (
	"database/sql"
	"github.com/xiongwei9/Gogogo/mysql/model"
	"log"
)

func CreateTable(db *sql.DB) {
	state := `
		CREATE TABLE IF NOT EXISTS users(
			id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
			username VARCHAR(64),
			password VARCHAR(64),
			status INT(4),
			createtime INT(10)
		);`
	if _, err := db.Exec(state); err != nil {
		log.Fatalf("create table failed: %v", err)
	}
	log.Println("create table user success")
}

// 插入数据
func InsertData(db *sql.DB) {
	result, err := db.Exec("insert INTO users(username,password) values(?,?)", "test2", "654321")
	if err != nil {
		log.Fatalf("Insert data failed, err: %v", err)
	}
	lastInsertID, err := result.LastInsertId() // 获取插入数据的自增ID
	if err != nil {
		log.Fatalf("Get insert id failed, err: %v", err)
	}
	log.Println("Insert data id:", lastInsertID)

	rowsAffected, err := result.RowsAffected() // 通过RowsAffected获取受影响的行数
	if err != nil {
		log.Fatalf("Get RowsAffected failed, err: %v", err)
	}
	log.Println("Affected rows:", rowsAffected)
}

// 查询单行
func QueryOne(db *sql.DB) {
	user := new(model.User) // 用new()函数初始化一个结构体对象
	row := db.QueryRow("select id,username,password from users where id=?", 1)
	// row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&user.Id, &user.Username, &user.Password); err != nil {
		log.Fatalf("scan failed, err: %v", err)
	}
	log.Println("Single row data:", *user)
}

// 查询多行
func QueryMulti(db *sql.DB) {
	user := new(model.User)
	rows, err := db.Query("select id,username,password from users where id = ?", 2)

	if err != nil {
		log.Fatalf("Query failed, err: %v", err)
	}
	defer func() {
		err := rows.Close() // 关闭掉未scan的sql连接
		if err != nil {
			log.Fatalf("close rows error: %v", err)
		}
	}()

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password) // 不scan会导致连接不释放
		if err != nil {
			log.Fatalf("Scan failed, err: %v", err)
		}
		log.Println("scan success:", *user)
	}
}

// 更新数据
func UpdateData(db *sql.DB) {
	result, err := db.Exec("UPDATE users set password=? where id=?", "111111", 1)
	if err != nil {
		log.Fatalf("Insert failed, err: %v", err)
	}
	log.Println("update data success:", result)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Get RowsAffected failed, err: %v", err)
		return
	}
	log.Println("Affected rows:", rowsAffected)
}

// 删除数据
func DeleteData(db *sql.DB) {
	result, err := db.Exec("delete from users where id=?", 1)
	if err != nil {
		log.Fatalf("Insert failed, err: %v", err)
	}
	log.Println("delete data success:", result)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("Get RowsAffected failed, err: %v", err)
	}
	log.Println("Affected rows:", rowsAffected)
}
