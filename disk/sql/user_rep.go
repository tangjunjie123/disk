package sql

type UserRepository struct {
	Id                 int
	Identity           string
	UserIdentity       string
	ParentId           int64
	RepositoryIdentity string
	Ext                string
	Name               string
}

func (t *UserRepository) Insert() error {
	tx := Db.Create(&t)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (t *UserRepository) Find() []UserRepository {
	repositories := []UserRepository{}
	Db.Where("Ext = ? or Name = ?", t.Ext, t.Name).Find(&repositories)
	//添加索引 ext_Name_UserIdentity
	return repositories
	//7b83ba9a-084f-11ee-bca8-8cc84be8ee13
}
