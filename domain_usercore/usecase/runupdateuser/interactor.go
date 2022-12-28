package runupdateuser

import (
	"context"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/entity"
)

type runupdateuserInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runupdateuserInteractor{
		outport: outputPort,
	}
}

func (r *runupdateuserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	userObj, err := entity.NewUserUpdate(req.UserUpdateRequest)
	if err != nil {
		return nil, err
	}

	userExisting, err := r.outport.FindUserByEmail(ctx, req.CurrentEmail)
	if err != nil {
		return nil, err
	}

	userExisting.Name = userObj.Name
	userExisting.Email = userObj.Email
	userExisting.UpdatedAt = userObj.UpdatedAt

	if userObj.Password != "" {
		hashPassword, err := r.outport.HashAndSaltPassword(ctx, []byte(userObj.Password))
		if err != nil {
			return nil, err
		}
		userExisting.Password = hashPassword
	}

	err = r.outport.SaveUser(ctx, userExisting)
	if err != nil {
		return nil, err
	}

	res.ID = userExisting.ID
	res.Name = userExisting.Name
	res.Email = userExisting.Email
	res.Status = userExisting.Status
	res.ActivationToken = userExisting.ActivationToken
	res.CreatedAt = userExisting.CreatedAt
	res.UpdatedAt = userExisting.UpdatedAt

	return res, nil
}
