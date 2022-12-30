package runtodocreate

import (
	"context"
	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/model/entity"
)

type runTodoCreateInteractor struct {
	outport  Outport
	outport2 Outport2
}

func NewUsecase(outputPort Outport, outport2 Outport2) Inport {
	return &runTodoCreateInteractor{
		outport:  outputPort,
		outport2: outport2,
	}
}

func (r *runTodoCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	todoObj, err := entity.NewTodo(req.TodoCreateRequest)
	if err != nil {
		return nil, err
	}

	user, err := r.outport2.FindUserByEmail(ctx, todoObj.UserID)
	if err != nil {
		return nil, err
	}

	todoObj.UserID = user.ID.String()
	err = r.outport.SaveTodo(ctx, todoObj)
	if err != nil {
		return nil, err
	}

	res.Todo = todoObj

	return res, nil
}
