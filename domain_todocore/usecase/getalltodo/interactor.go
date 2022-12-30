package getalltodo

import (
	"context"

	"zakariawahyu.com/go-gogen-mytodo/shared/util"
)

type getAllTodoInteractor struct {
	outport  Outport
	outport2 Outport2
}

func NewUsecase(outputPort Outport, outport2 Outport2) Inport {
	return &getAllTodoInteractor{
		outport:  outputPort,
		outport2: outport2,
	}
}

func (r *getAllTodoInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	user, err := r.outport2.FindUserByEmail(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	todoObjs, count, err := r.outport.FindAllTodoByUserID(ctx, req.Page, req.Size, user.ID.String())
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(todoObjs)

	return res, nil
}
