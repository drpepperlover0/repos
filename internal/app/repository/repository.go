package repository

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/drpepperlover0/internal/app/types"
	"github.com/drpepperlover0/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) types.UserRepository {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) Create(ctx context.Context, user *models.User) error {
	if !emailVerify(user.Email) {
		return errors.New("email verification error")
	}

	fmt.Println(user.ID)
	user.ID = 0

	return u.db.WithContext(ctx).Create(&user).Error
}

func (u *UserRepo) Get(ctx context.Context, id int) (*models.User, error) {
	var user *models.User

	if err := u.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		} else {
			return nil, fmt.Errorf("users finding error: %w", err)
		}
	}

	return user, nil
}

func (u *UserRepo) GetAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User

	if err := u.db.WithContext(ctx).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no records in DB")
		} else {
			return nil, fmt.Errorf("users finding error: %w", err)
		}
	}

	return users, nil
}

func (u *UserRepo) Delete(ctx context.Context, id int) error {
	user, _ := u.Get(ctx, id)

	if err := u.db.WithContext(ctx).Where("id = ?", id).Delete(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("not a single record was found to delete")
		} else {
			return fmt.Errorf("users delete error: %w", err)
		}
	}

	if err := u.db.WithContext(ctx).Exec("UPDATE users SET id=id-1 WHERE id > $1", id).Error; err != nil {
		return fmt.Errorf("update identity error: %w", err)
	}

	return nil
}

func emailVerify(email string) bool {
	reg := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return reg.MatchString(email)
}
