package usecase

import (
	"context"
	"renderService/layer"
	"renderService/utils"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (u *UserService) SignIn(c *gin.Context) error {
	//decode payload
	var user layer.User

	_, err := utils.DecodeUser(c, &user)
	if err != nil {
		//send error

		return err
	}

	//get grpc client
	grpc := GetGRPCClient(c)

	//sign in
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	userFromDB, err := grpc.User.GetByEmail(ctx, &layer.EmailPayload{Email: *user.Email})
	if err != nil {
		//send error

		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(*userFromDB.Password), []byte(*user.Password))
}

func (u *UserService) SignUp(c *gin.Context) error {
	//decode payload
	var user layer.User

	_, err := utils.DecodeUser(c, &user)
	if err != nil {
		//send error

		return err
	}

	//validate user
	err = utils.ValidateUser(&user, utils.VALIDATE_TYPE_CREATE)
	if err != nil {
		//send error

		return err
	}

	//hashingPassword
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(*user.Password), 5)
	if err != nil {
		//send error

		return err
	}

	hashedString := string(hashedByte)

	//set pass to hashedPassword
	user.Password = &hashedString

	//get grpc client
	grpc := GetGRPCClient(c)

	//sign up
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err = grpc.User.Create(ctx, &user)
	if err != nil {
		//send error

		return err
	}

	return nil
}

func (u *UserService) Edit(c *gin.Context) error {
	//get param id
	id, err := utils.GetParamID(c)
	if err != nil {
		//send error

		return err
	}

	//decode payload
	var user layer.User

	selectQ, err := utils.DecodeUser(c, &user)
	if err != nil {
		//send error

		return err
	}

	//validate user
	err = utils.ValidateUser(&user, utils.VALIDATE_TYPE_UPDATE)
	if err != nil {
		//send error

		return err
	}

	//hash password
	if user.Password != nil {
		hashedByte, err := bcrypt.GenerateFromPassword([]byte(*user.Password), 5)
		if err != nil {
			//send error
			return err
		}

		hashedString := string(hashedByte)
		user.Password = &hashedString
	}

	//get grpc client
	grpc := GetGRPCClient(c)

	//edit user
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = grpc.User.Edit(ctx, &layer.UserEditPayload{Select: selectQ, User: &user, ID: id})
	if err != nil {
		//send error

		return err
	}

	return nil
}

func (u *UserService) GetByID(c *gin.Context) (*layer.User, error) {
	//get param id
	id, err := utils.GetParamID(c)
	if err != nil {
		//send error

		return nil, err
	}

	//get grpc client
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	grpc := GetGRPCClient(c)

	//get user
	return grpc.User.GetByID(ctx, &layer.IDPayload{ID: id})
}
