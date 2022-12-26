package entity

import (
	"fmt"
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
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
}

func (req UserRequest) validate() error {
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
	return fmt.Sprintf("%s_%s", user.ID, user.Email)
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

type UserUpdateRequest struct {
	// add field to update here ...
}

func (r *User) Update(req UserUpdateRequest) error {

	// add validation and assignment value here ...

	return nil
}
