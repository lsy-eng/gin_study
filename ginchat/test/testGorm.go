package main

import (
	"fmt"
	"ginchat/unit3_IM/ginchat/models"
 	"gorm.io/gorm"
 	"gorm.io/driver/mysql"
  
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func main() {
  db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/ginchat@localhost?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{}) //这部分下载了数据库回头看
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&models.UserBasic{})

  // Create
  user := &models.UserBasic{}
  user.Name = "刘诗源"
  db.Create(user)

  

  // Read
  //var product Product
  fmt.Println(db.First(user, 1))
  //db.First(user, 1) // find product with integer primary key
  //db.First(user, "code = ?", "D42") // find product with code D42

  // Update - update product's price to 200
  db.Model(user).Update("PassWord", "1234")
  // Update - update multiple fields
  //db.Model(user).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
  //db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - delete product
  //db.Delete(&product, 1)
}