package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func init () {
	var err error
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("root:@/digitalent_bank")), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database" + err.Error())
	}
	DB.AutoMigrate(new(Account),new(Transaction))

}
