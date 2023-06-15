package main

import (
	"app/cmd/driver"
	"app/cmd/infrastructure/dao"
)

func main() {
	conn := dao.ConnectDB()
	dao.InitDB(conn)
	engine := driver.Router(conn)
	engine.Run(":8080")
}
