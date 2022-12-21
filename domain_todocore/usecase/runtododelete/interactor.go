package runtododelete

import (
	"context"
	"fmt"
)

type runtododeleteInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runtododeleteInteractor{
		outport: outputPort,
	}
}

func (r *runtododeleteInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// ambil data dengan id tertentu
	todoObj, err := r.outport.FindOneTodoById(ctx, req.TodoID)
	if err != nil {
		return nil, err
	}
	if todoObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	err = r.outport.DeleteTodoById(ctx, req.TodoID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
