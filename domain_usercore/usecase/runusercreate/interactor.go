package runusercreate

import (
	"context"

	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type userCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &userCreateInteractor{
		outport: outputPort,
	}
}

func (r *userCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	userObj, err := entity.NewUser(req.UserCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveUser(ctx, userObj)
	if err != nil {
		return nil, err
	}

	res.UserID = userObj.ID

	return res, nil
}
