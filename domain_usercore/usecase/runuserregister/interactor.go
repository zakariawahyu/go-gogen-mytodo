package runuserregister

import (
	"context"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/entity"
)

type runuserregisterInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runuserregisterInteractor{
		outport: outputPort,
	}
}

func (r *runuserregisterInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	userObj, err := entity.NewRegisterUser(req.UserRegisterRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveUser(ctx, userObj)
	if err != nil {
		return nil, err
	}

	res.User = userObj

	return res, nil
}
