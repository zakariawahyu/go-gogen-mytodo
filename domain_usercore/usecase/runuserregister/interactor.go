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

	hashPassword, err := r.outport.HashAndSaltPassword(ctx, []byte(req.Password))
	if err != nil {
		return nil, err
	}
	userObj, err := entity.NewUser(entity.UserRequest{
		RandomString: req.RandomString,
		Now:          req.Now,
		Name:         req.Name,
		Password:     hashPassword,
		Email:        req.Email,
	})
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
