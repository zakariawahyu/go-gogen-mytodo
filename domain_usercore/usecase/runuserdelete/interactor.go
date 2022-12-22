package runuserdelete

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type userDeleteInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &userDeleteInteractor{
		outport: outputPort,
	}
}

func (r *userDeleteInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

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

	err = r.outport.DeleteUser(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
