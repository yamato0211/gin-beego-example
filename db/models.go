package db

import (
	"github.com/astaxie/beego/orm"
	"github.com/google/uuid"
)

func (user *User) BeforeCreate(o orm.Ormer) (err error) {
	user.Id = uuid.New().String()
	return
}

type User struct {
	Id       string `json:"id" orm:"pk"`
	Name     string `json:"name" orm:"size(40)"`
	Email    string `json:"email" orm:"size(64);unique"`
	Password string `json:"password"`
}
