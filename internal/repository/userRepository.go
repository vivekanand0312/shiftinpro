package repository

import (
    "gorm.io/gorm"
    "shiftinpro/internal/models"
)

type UserRepository interface {
    CreateUser(user *models.User) error
    GetUserByID(id uint) (*models.User, error)
    GetUserByPhone(phone string) (*models.User, error)
    UpdateUserAddress(userID int, user models.User) error
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) GetUserByID(id uint) (*models.User, error) {
    var user models.User
    result := r.db.First(&user, id)
    return &user, result.Error
}

func (r *userRepository) GetUserByPhone(phone string) (*models.User, error) {
    var user models.User
    result := r.db.Where("phone = ?", phone).First(&user)
    return &user, result.Error
}

func (r *userRepository) UpdateUserAddress(userID int, user models.User) error {
    return r.db.Model(&models.User{}).Where("id = ?", userID).Updates(models.User{
        House:       user.House,
        Area:        user.Area,
        Landmark:    user.Landmark,
        SdAddressID: user.SdAddressID,
    }).Error
}
