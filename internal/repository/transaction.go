package repository

import (
	dbstore "github.com/EputraP/SMARTHOME_Backend/internal/store/db"
	"gorm.io/gorm"
)

type TransactionFunc func(tx *gorm.DB) error

type RepositoryTransaction interface {
	Begin() *gorm.DB
}

func AsTransaction(transactionFn TransactionFunc) error {
	db := dbstore.Get()
	tx := db.Begin()

	if err := transactionFn(tx); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
