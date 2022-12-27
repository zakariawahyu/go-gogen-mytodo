package getprofile

import (
	"context"
)

type getprofileInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &getprofileInteractor{
		outport: outputPort,
	}
}

func (r *getprofileInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	existingUser, err := r.outport.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	res.ID = existingUser.ID
	res.Name = existingUser.Name
	res.Email = existingUser.Email
	res.Status = existingUser.Status

	return res, nil
}
