package repository

import (
	"github.com/EputraP/SMARTHOME_Backend/internal/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	GetAccountById(inputModel *model.Account) (*model.Account, error)
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		db: db,
	}
}

type accountRepository struct {
	db *gorm.DB
}

func (r *accountRepository) GetAccountById(inputModel *model.Account) (*model.Account, error) {
	res := r.db.First(inputModel)
	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil
}


