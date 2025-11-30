package seeds

import (
	"log"

	"github.com/anddriii/warunggo/user-service/domain/models"
	"gorm.io/gorm"
)

func SeedRole(db *gorm.DB) {
	roles := []models.Role{
		{
			Name: "Super Admin",
		},
		{
			Name: "customer",
		},
	}

	for _, role := range roles {
		if err := db.FirstOrCreate(&role, models.Role{Name: role.Name}).Error; err != nil {
			log.Fatalf("%s: %v", err.Error(), err)
		} else {
			log.Printf("Role %s created", role.Name)
		}
	}
}
