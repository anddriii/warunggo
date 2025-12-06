package repository

import (
	"context"
	"errors"

	"github.com/anddriii/warunggo/user-service/domain/models"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type RoleRepoImpl struct {
	db *gorm.DB
}

// Create implements IRoleRepo.
func (r *RoleRepoImpl) Create(ctx context.Context, req *models.Role) error {
	role := models.Role{
		ID:   req.ID,
		Name: req.Name,
	}

	err := r.db.WithContext(ctx).Create(&role).Error
	if err != nil {
		log.Errorf("[RoleRepository -1] Create, %v", err)
		return nil
	}

	return nil
}

// Delete implements IRoleRepo.
func (r *RoleRepoImpl) Delete(ctx context.Context, id int64) error {
	modelRole := models.Role{}

	err := r.db.WithContext(ctx).Where("id = ?", id).Preload("Users").First(&modelRole).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("404")
			log.Infof("[ROleRepository -1] Delete: Role not found")
			return err
		}
		log.Errorf("[ROleRepository -2] Delete: %v", err)
		return err
	}

	if len(modelRole.Users) > 0 {
		err := errors.New("400")
		log.Infof("[RoleRepository -3] Delete: Role is associated with user")
		return err
	}

	if err := r.db.WithContext(ctx).Delete(&modelRole).Error; err != nil {
		log.Errorf("[RoleRepository -3] Delete: %v", err)
		return nil
	}

	return nil
}

// GetAll implements IRoleRepo.
func (r *RoleRepoImpl) GetAll(ctx context.Context, search string) ([]models.Role, error) {
	roles := []models.Role{}

	err := r.db.WithContext(ctx).Where("name ILIKE ?", "%"+search+"%").Find(&roles).Error
	if err != nil {
		log.Errorf("[RoleRepository-1] GetALl: %v", err)
		return nil, err
	}

	if len(roles) == 0 {
		err := errors.New("404")
		log.Infof("[RoleRepository-2]: No Role Found")
		return nil, err
	}

	return roles, nil
}

// GetByID implements IRoleRepo.
func (r *RoleRepoImpl) GetByID(ctx context.Context, id int64) (*models.Role, error) {
	roles := models.Role{}
	err := r.db.WithContext(ctx).Where("id == ?", id).First(&roles).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := errors.New("404")
			log.Errorf("[RoleRepository-1]: GetByID: Role Not Found")
			return nil, err
		}
		log.Errorf("[RoleRepository-2]: GetByID: %v", err)
		return nil, err
	}

	return &roles, nil
}

// Update implements IRoleRepo.
func (r *RoleRepoImpl) Update(ctx context.Context, req *models.Role, id int64) error {
	role := models.Role{
		Name: req.Name,
	}

	err := r.db.WithContext(ctx).Model(&models.Role{}).Where("id = ?", id).Updates(&role).Error
	if err != nil {
		log.Errorf("[RoleRepository -1] Update: %v", err)
		return nil
	}
	return nil
}

func NewROleRepo(db *gorm.DB) IRoleRepo {
	return &RoleRepoImpl{db: db}
}
