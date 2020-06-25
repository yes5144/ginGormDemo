package v1

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yes5144/ginGormDemo/dto"
	"github.com/yes5144/ginGormDemo/models"
	"github.com/yes5144/ginGormDemo/utils"
	"golang.org/x/crypto/bcrypt"
)

// Register xxx
func Register(c *gin.Context) {
	// parse post param
	var reqUser = models.User{}
	e := c.BindJSON(&reqUser)
	if e != nil {
		utils.Fail(c, nil, "parse req data from web err")
		return
	}

	name := reqUser.Name
	telephone := reqUser.Telephone
	password := reqUser.Password
	log.Println("params come from web:", name, telephone, password)

	// check password
	if len(password) < 6 {
		utils.Fail(c, nil, "length of password is less 6")
		return
	}

	// check telephone
	if len(telephone) != 11 {
		utils.Fail(c, nil, "telephone is invalid")
		return
	}

	var user models.User
	if user.IsTelephoneExist(telephone) {
		utils.Fail(c, nil, "telephone has already exist")
		return
	}

	// hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.Fail(c, nil, "password hashed fail msg")
		return
	}

	// if name not exist
	if len(name) == 0 {
		// name = "slkdjfkjkl"
		name = utils.RandomString(10)
	}
	log.Println("params will be inserted into db:", name, telephone, password)
	user = models.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}

	// insert new user into db
	err = user.Create(user)
	if err != nil {
		utils.Fail(c, nil, "insert new user error")
		return
	}

	utils.Success(c, nil, "register success")
}

// Login xxx
func Login(c *gin.Context) {
	// parse post param
	var reqUser = models.User{}
	e := c.BindJSON(&reqUser)
	if e != nil {
		utils.Fail(c, nil, "parse user data from web err")
	}

	telephone := reqUser.Telephone
	password := reqUser.Password
	log.Println("params come from web:", telephone, password)

	// check telephone
	if len(telephone) != 11 {
		utils.Fail(c, nil, "tel is invaid,telephone must be 11")
		return
	}

	// check password
	if len(password) < 6 {
		utils.Fail(c, nil, "length of password can not be less 6")
		return
	}

	// if telephone is in db?
	var user models.User
	ok := user.IsTelephoneExist(telephone)
	if !ok {
		log.Println("telephone not exist")
		return
	}

	// compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		utils.Fail(c, nil, "password is wrong")
		return
	}

	// generateToken
	token, err := utils.GenerateToken(user)
	if err != nil {
		utils.Fail(c, nil, "generate token error")
		return
	}

	log.Println(token)
	utils.Success(c, gin.H{"token": token}, "login success")
}

// UserInfo xxx
func UserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	utils.Success(c, gin.H{"user": dto.ToUserDto(user.(models.User))}, "user info")
}
