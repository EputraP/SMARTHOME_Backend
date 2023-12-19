package repository

import (
	"github.com/EputraP/SMARTHOME_Backend/internal/dto"
	"github.com/EputraP/SMARTHOME_Backend/internal/model"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Begin() *gorm.DB
	WithTx(tx *gorm.DB) AuthRepository
	GetAccountById(inputModel *model.User) (*model.User, error)
	SearchUser(username string) (*model.User, error)
	CreateUser(input dto.RegisterBody) (*model.User, error)
}

func NewAccountRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}
func (r authRepository) WithTx(tx *gorm.DB) AuthRepository {
	return &authRepository{
		db: tx,
	}
}

func (r authRepository) Begin() *gorm.DB {
	return r.db.Begin()
}
type authRepository struct {
	db *gorm.DB
}

func (r *authRepository) GetAccountById(inputModel *model.User) (*model.User, error) {
	res := r.db.First(inputModel)
	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil
}

func (r authRepository) SearchUser(username string) (*model.User, error) {
	var user *model.User

	res := r.db.First(&user, &model.User{Username: username})

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (r authRepository) CreateUser(input dto.RegisterBody) (*model.User, error) {
	user := &model.User{
		Username: input.UserName,
		Password: input.Password,
		
	}

	if dbc := r.db.Create(&user).Scan(&user); dbc.Error != nil {
		return nil, dbc.Error
	}

	return user, nil
}
