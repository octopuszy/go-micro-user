package user

import(
	"github.com/jinzhu/gorm"
	userModel "github.com/octopuszy/go-micro-user/domain/model/user"
)

type UserRepository struct {
	mysqlDb *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository  {
	return &UserRepository{mysqlDb: db}
}

func (r UserRepository) FindUserByName(name string) (*userModel.User, error){
	user := &userModel.User{}
	return user, r.mysqlDb.Table("user2").Where("user_name = ?", name).Find(user).Error
}

func (r UserRepository) FindUserById(id int64) (*userModel.User, error){
	user := &userModel.User{}
	return user, r.mysqlDb.Table("user2").First(user, id).Error
}

func (r UserRepository) CreateUser(user *userModel.User) (int64, error){
	return user.Id, r.mysqlDb.Table("user2").Create(user).Error
}

func (r UserRepository) DelUserById(id int64) error {
	return r.mysqlDb.Table("user2").Where("id = ?", id).Delete(&userModel.User{}).Error
}

func (r UserRepository) UpdateUser(user *userModel.User) error{
	return r.mysqlDb.Table("user2").Model(user).Update(user).Error
}

func (r UserRepository) FindAll(pageIndex int, pageSize int) (all []userModel.User, err error){
	 if pageIndex <= 0 {
		 pageIndex = 1
	 }
	if pageSize <= 0 {
		pageSize = 10
	}
	return all, r.mysqlDb.Table("user2").Limit(pageSize).Offset((pageIndex-1) * pageSize).Find(all).Error
}