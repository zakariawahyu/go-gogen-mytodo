package runuserlogin

import (
	"context"
	"time"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/entity"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/errorenum"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/token"
)

type runuserloginInteractor struct {
	outport Outport
	token   token.JWTToken
}

func NewUsecase(outputPort Outport, jwtToken token.JWTToken) Inport {
	return &runuserloginInteractor{
		outport: outputPort,
		token:   jwtToken,
	}
}

func (r *runuserloginInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	userObj, err := entity.NewUserLogin(req.UserLoginRequest)
	if err != nil {
		return nil, err
	}

	userExisting, err := r.outport.FindUserByEmail(ctx, userObj.Email)
	if err != nil {
		return nil, err
	}

	isValidPass := r.outport.ComparePassword(ctx, userExisting.Password, []byte(userObj.Password))
	if !isValidPass {
		return nil, errorenum.WrongEmailOrPassword
	}

	if !userExisting.IsActive() {
		return nil, errorenum.UserIsNotActive
	}

	userData := userExisting.GetUserData()
	token, err := r.token.CreateToken([]byte(userData), time.Hour)
	if err != nil {
		return nil, err
	}

	res.Token = token

	return res, nil
}
