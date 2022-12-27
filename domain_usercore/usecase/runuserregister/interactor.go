package runuserregister

import (
	"context"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/entity"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/errorenum"
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

	userExisting, _ := r.outport.FindUserByEmail(ctx, req.Email)

	if userExisting != nil {
		return nil, errorenum.UserAlreadyExist
	}

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

	res.ID = userObj.ID
	res.Name = userObj.Name
	res.Email = userObj.Email
	res.Status = userObj.Status
	res.ActivationToken = userObj.ActivationToken
	res.Created = userObj.Created

	return res, nil
}
