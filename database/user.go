package database

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	userID    string    `json:"userId" gorm:"column:user_id"`
	mobileNo  string    `json:"mobileNo" gorm:"column:mobile_no"`
	firstName string    `json:"titleTH" gorm:"column:first_name"`
	lastName  string    `json:"lastName" gorm:"column:last_name"`
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
	user.userID = uuid.New().String()
	if err := db.Create(&user).Error; err != nil {
		return *user, err
	}

	return *user, nil
}
