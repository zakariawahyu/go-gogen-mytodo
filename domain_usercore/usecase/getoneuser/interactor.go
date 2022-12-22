package getoneuser

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type userGetOneInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &userGetOneInteractor{
		outport: outputPort,
	}
}

func (r *userGetOneInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	userObj, err := r.outport.FindOneUserByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	if userObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	res.User = userObj

	return res, nil
}
