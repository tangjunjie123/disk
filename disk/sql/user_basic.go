package sql

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Id          int
	Identity    string
	Name        string
	Password    string
	Email       string
	NowVolume   int64
	TotalVolume int64
}

func (t *UserBasic) FindUser() (*gorm.DB, error) {
	tx := Db.Where("Name = ? && Password = ?", t.Name, t.Password).Find(&t)
	fmt.Println(t)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (t *UserBasic) UpdateUser() error {
	tx := Db.Model(&t).Where("name = ?", t.Name).Update("password", t.Password)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func (t *UserBasic) Findemail() (bool, error) {
	basic := &UserBasic{}
	tx := Db.Where("email = ?", t.Email).Find(&basic)
	if tx.Error != nil {
		return false, tx.Error
	}
	if basic.Id == 0 {
		return true, nil
	}
	return false, nil
}
func (t *UserBasic) CreateUser() error {
	user := UserBasic{Name: t.Name}
	fmt.Println(user)
	tx := Db.Where("Name = ?", user.Name).Find(&user)
	fmt.Println(user)
	if tx.Error != nil {
		return tx.Error
	}
	if user.Id == 0 {
		create := Db.Create(t)
		if create.Error != nil {
			return create.Error
		}
		return nil
	}
	return errors.New("账号重复")
}
