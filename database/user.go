package database

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    string    `json:"userID" gorm:"column:user_id"`
	MobileNo  string    `json:"mobileNo" gorm:"column:mobile_no"`
	FirstName string    `json:"firstName" gorm:"column:first_name"`
	LastName  string    `json:"lastName" gorm:"column:last_name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

func GetListByUserID(userID string) (user []User, err error) {
	db := GetDB()
	err = db.Order("created_at desc").Where("user_id = ?", userID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user *User) (User, error) {
	db := GetDB()
	user.UserID = uuid.New().String()
	if err := db.Create(&user).Error; err != nil {
		return *user, err
	}

	return *user, nil
}
