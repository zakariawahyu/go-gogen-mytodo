package getalluser

import (
	"context"

	"zakariawahyu.com/go-gogen-mytodo/shared/util"
)

//go:generate mockery --name Outport -output mocks/

type userGetAllInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &userGetAllInteractor{
		outport: outputPort,
	}
}

func (r *userGetAllInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	userObj, count, err := r.outport.FindAllUser(ctx, req.FindAllUserFilterRequest)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(userObj)

	return res, nil
}
