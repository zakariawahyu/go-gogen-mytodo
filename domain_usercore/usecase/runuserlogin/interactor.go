package runuserlogin

import (
	"context"
	"fmt"
	"time"
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

	userObj, err := r.outport.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	isValidPass := r.outport.ComparePassword(ctx, userObj.Password, []byte(req.Password))
	if !isValidPass {
		return nil, fmt.Errorf("wrong email and password")
	}

	token, err := r.token.CreateToken([]byte(userObj.Email), time.Hour)
	if err != nil {
		return nil, err
	}

	res.Token = token

	return res, nil
}
