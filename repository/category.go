package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	result := []entity.Category{}
	resp := r.db.Raw("SELECT * FROM categories WHERE user_id = ?", id).Scan(&result)
	if resp.Error != nil {
		return []entity.Category{}, resp.Error
	}
	if resp.RowsAffected == 0 {
		return []entity.Category{}, nil
	}

	// if err := r.db.Where("user_id = ?", id).First(&result).Error; err != nil {
	// 	return []entity.Category{}, nil
	// }
	return result, nil // TODO: replace this
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	result := r.db.Create(category)
	if result.Error != nil {
		return 0, result.Error
	}
	return category.ID, nil // TODO: replace this
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	result := r.db.Create(&categories)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	result := entity.Category{}
	resp := r.db.Raw("SELECT * FROM categories WHERE id = ?", id).Scan(&result)
	if resp.Error != nil {
		return entity.Category{}, resp.Error
	}
	if resp.RowsAffected == 0 {
		return entity.Category{}, nil
	}

	// if err := r.db.Where("id = ?", id).First(&result).Error; err != nil {
	// 	return entity.Category{}, nil
	// }
	return result, nil // TODO: replace this
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	result := r.db.Model(&entity.Category{}).Updates(category)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	result := r.db.Delete(&entity.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}
