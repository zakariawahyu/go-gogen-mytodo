package runtodocheck

import (
	"context"
	"fmt"
)

type runTodoCheckInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runTodoCheckInteractor{
		outport: outputPort,
	}
}

func (r *runTodoCheckInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// ambil data dengan id tertentu
	todoObj, err := r.outport.FindOneTodoById(ctx, req.TodoID)
	if err != nil {
		return nil, err
	}
	if todoObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	// kemudian ubah checked = true
	err = todoObj.Check()
	if err != nil {
		return nil, err
	}

	// kemudian akan kembali disimppan ke persistenstorage
	err = r.outport.SaveTodo(ctx, todoObj)
	if err != nil {
		return nil, err
	}

	res.Todo = todoObj

	return res, nil
}
