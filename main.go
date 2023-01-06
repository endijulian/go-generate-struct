package main

import (
	"fmt"

	"github.com/fwhezfwhez/model_convert"
)

func main() {
	dsn := "host=hostname user=postgres password=Vpassword dbname=database-name port=0000 sslmode=disable TimeZone=Asia/Shanghai"
	tableName := "tr_communication_registrations"
	fmt.Println(model_convert.TableToStructWithTag(dsn, tableName, "postgres"))
}
