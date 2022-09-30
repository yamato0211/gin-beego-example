package db

import (
	"beego-tutorial/utils"
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func InitDB() (err error) {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		utils.DbUser, utils.DbPass, utils.DbHost, utils.DbPort, utils.DbName,
	)
	orm.RegisterDataBase("default", "postgres", dsn)
	orm.RegisterModel(new(User))
	if err = orm.RunSyncdb("default", true, true); err != nil {
		fmt.Println(err)
	}
	return
}
