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

	userObj, err := entity.NewUser(req.UserRequest)
	if err != nil {
		return nil, err
	}

	userExisting, _ := r.outport.FindUserByEmail(ctx, userObj.Email)
	if userExisting != nil {
		return nil, errorenum.UserAlreadyExist
	}

	hashPassword, err := r.outport.HashAndSaltPassword(ctx, []byte(userObj.Password))
	if err != nil {
		return nil, err
	}

	userObj.Password = hashPassword
	err = r.outport.SaveUser(ctx, userObj)
	if err != nil {
		return nil, err
	}

	res.ID = userObj.ID
	res.Name = userObj.Name
	res.Email = userObj.Email
	res.Status = userObj.Status
	res.ActivationToken = userObj.ActivationToken
	res.CreatedAt = userObj.CreatedAt
	res.UpdatedAt = userObj.UpdatedAt

	return res, nil
}
