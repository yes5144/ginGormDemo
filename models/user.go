package models

import (
	"log"
	"strings"
)

// User xxx
type User struct {
	BaseModel
	Name      string `json:"name,omitempty"`
	Telephone string `json:"telephone,omitempty" gorm:"unique_index"`
	Password  string `json:"password,omitempty"`
}

// IsTelephoneExist xxx
func (u *User) IsTelephoneExist(tel string) bool {
	Db.Where("telephone=?", tel).First(u)
	log.Println(u.ID)
	if u.ID > 0 {
		return true
	}
	return false
}

// Create xxx
func (u *User) Create(user User) error {
	Db.Create(&user)
	return nil
}

// Delete xxx
func (u *User) Delete(ids string) bool {
	Db.Where("id in (?)", ids).Delete(u)
	return true
}

// Update xxx
func (u *User) Update(user User) bool {
	Db.Model(&user).Update(u)
	return true
}

// SelectIds xxx
func (u *User) SelectIds(ids string) bool {
	Db.Where("id in (?)", []string(strings.Split(ids, ","))).First(u)
	return true
}

// SelectByTel xxx
func (u *User) SelectByTel(tel string) bool {
	Db.Where("telephone = ?", tel).First(u)
	return true
}
