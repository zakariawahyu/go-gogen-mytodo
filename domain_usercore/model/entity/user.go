package entity

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/errorenum"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/vo"
)

type User struct {
	ID              vo.UserID `bson:"_id" json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	Status          bool      `json:"status"`
	ActivationToken string    `json:"activation_token"`
	Created         time.Time `bson:"created" json:"created"`
}

type UserRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`
	Name         string    `json:"name" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	Password     string    `json:"password" validate:"required,min=6"`
}

func (req UserRequest) validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}

func NewUser(req UserRequest) (*User, error) {

	id, err := vo.NewUserID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	if err = req.validate(); err != nil {
		return nil, err
	}

	var obj User
	obj.ID = id
	obj.Name = req.Name
	obj.Email = req.Email
	obj.Password = req.Password
	obj.Created = req.Now
	obj.Status = false
	obj.ActivationToken = uuid.NewString()

	return &obj, nil
}

func (user *User) IsActive() bool {
	return user.Status == true
}

func (user *User) GetUserData() string {
	return user.Email
}

func (user *User) ActivatedUser() error {
	if user.Status {
		return errorenum.UserAlreadyActivated
	}

	user.Status = true
	return nil
}

func (user *User) CompareActivatedToken(activatedToken string) error {
	if user.ActivationToken != activatedToken {
		return errorenum.UserActivatedTokenNotMatch
	}
	return nil
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (req UserLoginRequest) validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}

func NewUserLogin(req UserLoginRequest) (*User, error) {
	if err := req.validate(); err != nil {
		return nil, err
	}

	var user User
	user.Email = req.Email
	user.Password = req.Password

	return &user, nil
}

type UserUpdateRequest struct {
	// add field to update here ...
}

func (r *User) Update(req UserUpdateRequest) error {

	// add validation and assignment value here ...

	return nil
}
