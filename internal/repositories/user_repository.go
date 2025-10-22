package repositories

import (
	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserById(userID uint) *models.User
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{db: database.DB}
}

func (r *userRepository) GetUserById(userID uint) *models.User {
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil
	}
	return &user
}
