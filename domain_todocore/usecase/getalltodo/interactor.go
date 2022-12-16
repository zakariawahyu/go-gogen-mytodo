package getalltodo

import (
	"context"

	"zakariawahyu.com/go-gogen-mytodo/shared/util"
)

type getAllTodoInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &getAllTodoInteractor{
		outport: outputPort,
	}
}

func (r *getAllTodoInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	todoObjs, count, err := r.outport.FindAllTodo(ctx, req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(todoObjs)

	return res, nil
}
