package sql

import (
	"gorm.io/gorm"
)

type RepositoryPool struct {
	Id       int
	Identity string
	Hash     string
	Name     string
	Ext      string
	Size     int64
	Path     string
	gorm.Model
}

func (t *RepositoryPool) Insert() error {
	tx := Db.Create(t)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (t *RepositoryPool) Find() error {
	tx := Db.Where("identity", t.Identity).Find(&t)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (t *RepositoryPool) FindH() error {
	tx := Db.Where("Hash", t.Hash).Find(&t)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
