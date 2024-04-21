package repos

import (
	"codelabx/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "abhi:abhi1060@tcp(codelabx-auth-db.cbc0yumwi7uk.ap-south-1.rds.amazonaws.com:3306)/codelabx_users?charset=utf8mb4&parseTime=True&loc=Local"
)

var db *gorm.DB

func init() {
	db, _ = connectToDB()
}

func connectToDB() (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error in Connect to DB with : ", err)
	}

	autoMigrateModels(db)
	return db, err
}

func autoMigrateModels(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

func AddUser(user *models.User) int {
	res := db.Save(user)
	if res.Error != nil {
		log.Fatal("err during add user : ", res.Error)
		return 0
	}

	fmt.Println("User saved successfully")
	return 1
}

func UserExists(user *models.User) bool {
	var res int64
	err := db.Model(&models.User{}).Where("user_name = ?", user.UserName).Count(&res).Error
	if err != nil {
		log.Fatal("err during user Exists : ", err)
	}
	return res > 0
}

func IsValidPassword(user *models.User) bool {
	var dbUser = models.User{UserName: user.UserName}
	res := db.First(&dbUser)

	if res.Error != nil {
		log.Fatal("Error during validate pass : ", res.Error)
		return false
	}

	return user.Password == dbUser.Password
}
