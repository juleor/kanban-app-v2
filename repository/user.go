package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	result := entity.User{}
	resp := r.db.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&result)
	if resp.Error != nil {
		return entity.User{}, resp.Error
	}
	if resp.RowsAffected == 0 {
		return entity.User{}, nil
	}
	return result, nil // TODO: replace this
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	result := entity.User{}
	resp := r.db.Raw("SELECT * FROM users WHERE email = ?", email).Scan(&result)
	if resp.Error != nil {
		return entity.User{}, resp.Error
	}
	if resp.RowsAffected == 0 {
		return entity.User{}, nil
	}
	return result, nil // TODO: replace this
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil // TODO: replace this
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	result := r.db.Model(&entity.User{}).Updates(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil // TODO: replace this
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	result := r.db.Delete(&entity.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}
