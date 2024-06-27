package storage

import (
	"fmt"
	"go-gin/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() {
	dsn := `host=localhost
            user=postgres
            password=1
            dbname=postgres
            port=5432
            sslmode=disable
            TimeZone=Asia/Shanghai`
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)

	if err != nil {
		return
	}

	db.AutoMigrate(&model.Album{})
}
