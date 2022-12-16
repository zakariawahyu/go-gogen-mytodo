package entity

import (
	"time"

	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/model/vo"
)

type Todo struct {
	ID      vo.TodoID `bson:"_id" json:"id"`
	Message string    `json:"message"`
	Checked bool      `json:"checked"`
	Created time.Time `bson:"created" json:"created"`
}

type TodoCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`
	Message      string    `json:"message"`
}

func NewTodo(req TodoCreateRequest) (*Todo, error) {

	id, err := vo.NewTodoID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	// add validation and assignment value here ...

	var obj Todo
	obj.ID = id
	obj.Created = req.Now

	return &obj, nil
}

type TodoUpdateRequest struct {
	// add field to update here ...
}

func (r *Todo) Update(req TodoUpdateRequest) error {

	// add validation and assignment value here ...

	return nil
}
