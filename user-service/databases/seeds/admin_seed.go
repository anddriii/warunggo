package seeds

import (
	"log"

	"github.com/anddriii/warunggo/user-service/common/util"
	"github.com/anddriii/warunggo/user-service/domain/models"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	bytes, err := util.HashPassword("admin123")
	if err != nil {
		log.Fatalf("%s:%v", err.Error(), err)
	}

	modelRole := models.Role{}
	err = db.Where("name = ?", "Super Admin").First(&modelRole).Error
	if err != nil {
		log.Fatalf("%s: %v", err.Error(), err)
	}

	admin := models.User{
		Name:       "super admin",
		Email:      "superadmin@gmail.com",
		Password:   bytes,
		IsVerified: true,
		Role:       []models.Role{modelRole},
	}

	if err := db.FirstOrCreate(&admin, models.User{Email: "superadmin@gmail.com"}).Error; err != nil {
		log.Fatalf("%s: %v", err.Error(), err)
	} else {
		log.Printf("Admin %s created", admin.Name)
	}
}
