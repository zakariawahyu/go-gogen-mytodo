package runuseractivated

import (
	"context"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/errorenum"
)

type runuseractivatedInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runuseractivatedInteractor{
		outport: outputPort,
	}
}

func (r *runuseractivatedInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	userExisting, err := r.outport.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = userExisting.CompareActivatedToken(req.ActivationToken)
	if err != nil {
		return nil, err
	}

	if userExisting.IsActive() {
		return nil, errorenum.UserAlreadyActivated
	}

	err = userExisting.ActivatedUser()
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveUser(ctx, userExisting)
	if err != nil {
		return nil, err
	}

	return res, nil
}
