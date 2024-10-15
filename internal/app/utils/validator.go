package utils

import (
	"errors"
	model "go-gin/internal/app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ValidateStruct(ctx *gin.Context, s interface{}) error {
	validate := validator.New()

	dbValue, ok := ctx.Get("db")
	if !ok {
		return errors.New("db nout found")
	}
	db, ok := dbValue.(*gorm.DB)
	if !ok {
		return errors.New("cant")
	}

	// Передача db как аргумента в анонимную функцию
	validate.RegisterValidation("uniqueEmail", func(fl validator.FieldLevel) bool {
		// Используем db  внутри анонимной функции
		return validateUniqueEmail(fl, db)
	})

	return validate.Struct(s)
}

func validateUniqueEmail(fl validator.FieldLevel, db *gorm.DB) bool {
	email := fl.Field().String()
	var count int64
	db.Model(&model.Album{}).Where("email = ?", email).Count(&count)
	return count == 0
}
