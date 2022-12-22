package entity

import (
	"time"

	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/vo"
)

type User struct {
	ID      vo.UserID `bson:"_id" json:"id"`
	Created time.Time `bson:"created" json:"created"`
}

type UserCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`
}

func NewUser(req UserCreateRequest) (*User, error) {

	id, err := vo.NewUserID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	// add validation and assignment value here ...

	var obj User
	obj.ID = id
	obj.Created = req.Now

	return &obj, nil
}

type UserUpdateRequest struct {
	// add field to update here ...
}

func (r *User) Update(req UserUpdateRequest) error {

	// add validation and assignment value here ...

	return nil
}
