package main

import (
	"beego-tutorial/db"
	"beego-tutorial/utils"
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	utils.LoadEnv()
	db.InitDB()
	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	api.GET("/post_db", func(c *gin.Context) {
		var user db.User
		o := orm.NewOrm()
		user.BeforeCreate(o)
		user.Name = "Yamato"
		user.Email = "test@test.com"
		user.Password = "password"
		_, err := o.Insert(&user)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(200, gin.H{
			"message": "post success",
		})
	})

	api.GET("/get_db", func(c *gin.Context) {
		var users []db.User
		o := orm.NewOrm()
		o.QueryTable(new(db.User)).All(&users)
		println(users)
		c.JSON(200, gin.H{
			"message": users,
		})
	})

	api.Run(fmt.Sprintf(":%s", utils.ApiPort))
}
