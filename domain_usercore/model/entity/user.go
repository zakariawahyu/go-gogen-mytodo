package entity

import (
	"golang.org/x/crypto/bcrypt"
	"time"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/errorenum"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/vo"
)

type User struct {
	ID       vo.UserID `bson:"_id" json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Created  time.Time `bson:"created" json:"created"`
}

type UserRegisterRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req UserRegisterRequest) validate() error {
	if req.Name == "" {
		return errorenum.NameMustNotEmpty
	}

	if req.Email == "" {
		return errorenum.EmailMustNotEmpty
	}

	if req.Password == "" {
		return errorenum.PasswordMustNotEmpty
	}

	return nil
}

func (req UserLoginRequest) validate() error {
	if req.Email == "" {
		return errorenum.EmailMustNotEmpty
	}

	if req.Password == "" {
		return errorenum.PasswordMustNotEmpty
	}

	return nil
}

func NewRegisterUser(req UserRegisterRequest) (*User, error) {

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
	obj.Password = hashAndSalt([]byte(req.Password))
	obj.Created = req.Now

	return &obj, nil
}

func NewLoginUser(req UserLoginRequest) (*User, error) {

	if err := req.validate(); err != nil {
		return nil, err
	}

	var obj User
	obj.Email = req.Email
	obj.Password = req.Password

	return &obj, nil
}

type UserUpdateRequest struct {
	// add field to update here ...
}

func (r *User) Update(req UserUpdateRequest) error {

	// add validation and assignment value here ...

	return nil
}

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
