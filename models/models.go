package model

type Album struct {
	ID    int     `json:"id" gorm:"primaryKey"`
	Title string  `json:"title" validate:"reqiured,min=2,max=20"`
	Actor string  `json:"actor" validate:"reqiured,min=2,max=20"`
	Price float64 `json:"price" validate:"required,gte=2,lte=120"`
	Email string  `json:"email" validate:"required,email, uniqueEmail"`
	// Password string  `json:"password" validate:"required, min=8, regexp=^[a-zA-Z0-9]+$"`
}

// func ValidateUniqueEmail(fl validator.FieldLevel) bool {
// 	db, ok := fl.Field().Interface().(*gorm.DB)
// 	if !ok {
// 		return false
// 	}
// 	var count int64
// 	db.Model(&Album{}).Where("email = ?", fl.Field().String()).Count(&count)
// 	return count == 0

// }
