package runuserupdate

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type userUpdateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &userUpdateInteractor{
		outport: outputPort,
	}
}

func (r *userUpdateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	userObj, err := r.outport.FindOneUserByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if userObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	// you may need to check the authorization part here
	// is this user allowed to perform this action ?

	err = userObj.Update(req.UserUpdateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveUser(ctx, userObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
