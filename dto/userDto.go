package dto

import "github.com/yes5144/ginGormDemo/models"

// UserDto xxx
type UserDto struct {
	Name      string
	Telephone string
}

// ToUserDto xxx
func ToUserDto(user models.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
