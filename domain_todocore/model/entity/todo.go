package entity

import (
	"time"
	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/model/errorenum"
	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/model/vo"
)

type Todo struct {
	ID        vo.TodoID `bson:"_id" json:"id"`
	UserID    string    `bson:"user_id" json:"user_id"`
	Message   string    `json:"message"`
	Checked   bool      `json:"checked"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type TodoCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`
	UserID       string    `json:"_"`
	Message      string    `json:"message"`
}

func NewTodo(req TodoCreateRequest) (*Todo, error) {

	id, err := vo.NewTodoID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	// add validation and assignment value here ...
	if req.Message == "" {
		return nil, errorenum.MessageMustNotEmpty
	}

	var obj Todo
	obj.ID = id
	obj.UserID = req.UserID
	obj.Message = req.Message
	obj.Checked = false
	obj.CreatedAt = req.Now
	obj.UpdatedAt = req.Now

	return &obj, nil
}

type TodoUpdateRequest struct {
	// add field to update here ...
}

func (r *Todo) Update(req TodoUpdateRequest) error {

	// add validation and assignment value here ...

	return nil
}

func (r *Todo) Check() error {
	if r.Checked {
		return errorenum.TodoHasBeenChecked
	}
	r.Checked = true

	return nil
}
