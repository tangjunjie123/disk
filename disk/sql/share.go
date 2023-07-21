package sql

import "gorm.io/gorm"

type ShareBasic struct {
	gorm.Model
	Id                     int
	Identity               string
	UserIdentity           string
	UserRepositoryIdentity string
	RepositoryIdentity     string
	ExpiredTime            int
	ClickNum               int
}

func (t *ShareBasic) Insert() error {
	tx := Db.Create(t)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
