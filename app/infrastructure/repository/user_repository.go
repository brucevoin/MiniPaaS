package repository

import (
	"fmt"
	model "mini-paas/app/domain/model/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserPO struct {
	gorm.Model
	Id    int
	Name  string
	Email string
}
type UserRepository struct {
}

func (s *UserRepository) SaveUser(user *model.User) {
	fmt.Println("save user")

}
func (s *UserRepository) UpdateUser(user *model.User) {

	fmt.Println("update user")
}

func test() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 创建表
	db.AutoMigrate(&UserPO{})

	// 创建记录
	user := UserPO{Name: "Jack", Email: "jack@example.com"}
	db.Create(&user)

	// 查询记录
	var result UserPO
	db.First(&result, "name = ?", "Jack")
	fmt.Println(result)

	// 更新记录
	db.Model(&result).Update("Email", "new_email@example.com")

	// 删除记录
	db.Delete(&result)
}
